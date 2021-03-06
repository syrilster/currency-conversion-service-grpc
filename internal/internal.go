package internal

import (
	"fmt"
	"github.com/syrilster/currency-conversion-service-grpc/internal/config"
	"github.com/syrilster/currency-conversion-service-grpc/internal/currencyconversion"
	"github.com/syrilster/currency-conversion-service-grpc/internal/currencyexchange"
	"github.com/syrilster/currency-conversion-service-grpc/internal/middlewares"
	"net/http"
)

//StatusRoute health check route
func StatusRoute() (route config.Route) {
	route = config.Route{
		Path:    "/status",
		Method:  http.MethodGet,
		Handler: middlewares.RuntimeHealthCheck(),
	}
	return route
}

type ServerConfig interface {
	Version() string
	BaseURL() string
	CurrencyExchangeClient() currencyexchange.ClientInterface
}

func SetupServer(cfg ServerConfig) *config.Server {
	basePath := fmt.Sprintf("/%v", cfg.Version())
	currencyExchangeService := currencyconversion.NewService(cfg.CurrencyExchangeClient())
	server := config.NewServer().
		WithRoutes(
			"", StatusRoute(),
		).
		WithRoutes(
			basePath,
			currencyconversion.Route(currencyExchangeService),
		)
	return server
}
