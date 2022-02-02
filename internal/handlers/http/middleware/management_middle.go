package middleware

import (
	"net/http"
	"fmt"

	"github.com/go-mock-api/internal/exceptions"
	"github.com/go-mock-api/internal/handlers/http/handlers"
	"github.com/go-mock-api/internal/viper"

)

type ManagementMiddleware interface {
	Management(next http.Handler) http.Handler
}

type ManagementMiddlewareImpl struct {
	responseHandlers handlers.ResponseHandler
}

func NewManagementMiddleware(responseHandlers handlers.ResponseHandler) ManagementMiddleware {
	return ManagementMiddlewareImpl{
		responseHandlers: responseHandlers,
	}
}

func (m ManagementMiddlewareImpl) Management(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if viper.Application.Setup.ResponseStatusCode != 200 {
			m.responseHandlers.BadRequest(w, r, exceptions.ErrEnforced)
			return
		}
		fmt.Println("OLAAAAA", viper.Application.Setup)
		next.ServeHTTP(w, r)
	})
}