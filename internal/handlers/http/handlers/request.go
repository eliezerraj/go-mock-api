package handlers

import (
	"sync"

)


var (
	requestHandlers RequestHandlers
	onceRequest     sync.Once
)

type RequestHandlers interface {
}

type RequestHandlersImpl struct{}

func NewRequestHandlers() RequestHandlers {
	return RequestHandlersImpl{}
}

func GetRequestHandlersInstance() RequestHandlers {
	onceRequest.Do(func() {
		requestHandlers = NewRequestHandlers()
	})
	return requestHandlers
}