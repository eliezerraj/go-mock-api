package handlers

import (
	"sync"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/go-chi/render"

	"github.com/go-mock-api/internal/exceptions"

)

var (
	requestHandlers RequestHandlers
	onceRequest     sync.Once
)

type RequestHandlersImpl struct{}

type RequestHandlers interface {
	BindJson(r *http.Request, destination interface{}) error
}

func NewRequestHandlers() RequestHandlers {
	return RequestHandlersImpl{}
}

func GetRequestHandlersInstance() RequestHandlers {
	onceRequest.Do(func() {
		requestHandlers = NewRequestHandlers()
	})
	return requestHandlers
}

func (h RequestHandlersImpl) BindJson(r *http.Request, destination interface{}) error {
	e := render.DecodeJSON(r.Body, destination)
	if e != nil {
		return exceptions.Throw(e, exceptions.ErrJsonDecode)
	}
	return h.validateJsonFields(destination)
}

func (h RequestHandlersImpl) validateJsonFields(input interface{}) error {
	v := validator.New()
	if err := v.Struct(input); err != nil {
		return err
	}
	return nil
}