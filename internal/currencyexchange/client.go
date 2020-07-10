package currencyexchange

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/syrilster/currency-conversion-service-grpc/gRPC/gen"

	"google.golang.org/grpc"
)

type ClientInterface interface {
	ExchangeRate(ctx context.Context, request Request) (*Response, error)
}

func NewClient(host string, port string) *client {
	return &client{ClientHost: host, ClientPort: port}
}

type client struct {
	ClientHost string
	ClientPort string
}

func (c *client) ExchangeRate(ctx context.Context, request Request) (*Response, error) {
	contextLogger := log.WithContext(ctx)
	response := &Response{}

	connection, err := grpc.Dial(c.ClientHost+":"+c.ClientPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("gRPC target error: could not connect: %v", err)
	}
	defer connection.Close()

	rpcClient := gen.NewCurrencyExchangeServiceClient(connection)

	req := &gen.Request{
		FromCurrency: request.FromCurrency,
		ToCurrency:   request.ToCurrency,
	}
	res, err := rpcClient.GetExchangeRate(context.Background(), req)
	if err != nil {
		contextLogger.WithError(err).Errorf("Error calling the currency exchange RPC")
		return nil, err
	}
	contextLogger.Infof("Response from currency-exchange-service: %v", res.ConversionMultiple)
	response.FromCurrency = request.FromCurrency
	response.ToCurrency = request.ToCurrency
	response.ConversionMultiple = fmt.Sprintf("%f", res.ConversionMultiple)

	return response, nil
}
