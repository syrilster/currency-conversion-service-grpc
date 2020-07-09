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

func NewClient() *client {
	return &client{}
}

type client struct{}

func (client *client) ExchangeRate(ctx context.Context, request Request) (*Response, error) {
	response := &Response{}
	cc, err := grpc.Dial("localhost"+":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := gen.NewCurrencyExchangeServiceClient(cc)

	req := &gen.Request{
		FromCurrency: request.FromCurrency,
		ToCurrency:   request.ToCurrency,
	}
	res, err := c.GetExchangeRate(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling gRPC: %v", err)
	}
	log.Printf("Response from Service: %v", res.ConversionMultiple)
	response.FromCurrency = request.FromCurrency
	response.ToCurrency = request.ToCurrency
	response.ConversionMultiple = fmt.Sprintf("%f", res.ConversionMultiple)

	return response, nil
}
