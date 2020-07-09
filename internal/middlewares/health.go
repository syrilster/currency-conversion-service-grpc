package middlewares

import (
	"github.com/syrilster/currency-conversion-service-grpc/internal/util"
	"net/http"
)

func RuntimeHealthCheck() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		util.WithBodyAndStatus("All OK", http.StatusOK, w)
	}
}
