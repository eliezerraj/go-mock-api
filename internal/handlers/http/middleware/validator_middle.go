package middleware

import (
	"net/http"
	//"fmt"

	"github.com/go-mock-api/internal/exceptions"
	//"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/handlers/http/handlers"

)

type ValidatorMiddlewareImpl struct {
	responseHandlers handlers.ResponseHandler
}

type ValidatorMiddleware interface {
	Validate() func(next http.Handler) http.Handler
}

func NewValidatorMiddleware(responseHandlers handlers.ResponseHandler) ValidatorMiddleware {
	return ValidatorMiddlewareImpl{
		responseHandlers: responseHandlers,
	}
}

func (v ValidatorMiddlewareImpl) Validate() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//bearToken := r.Header.Get(constants.Authorization)
			//fmt.Println("bearToken :", bearToken)
			if 1==2 {
				v.responseHandlers.BadRequest(w, r, exceptions.ErrTokenUnreachable)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}