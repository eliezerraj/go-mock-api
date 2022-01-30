package routers

import (
	"net/http"

	"github.com/go-chi/chi"

)

type Health struct {
	service          services.HealthService
	requestHandlers  handlers.RequestHandlers
	responseHandlers handlers.ResponseHandler
	//uuidValidator    middlewares.UUIDValidatorMiddleware
}

func NewHealth(service services.HealthService, requestHandlers handlers.RequestHandlers, responseHandlers handlers.ResponseHandler, uuidValidator middlewares.UUIDValidatorMiddleware) Travels {
	return Health{
		service:          service,
		requestHandlers:  requestHandlers,
		responseHandlers: responseHandlers,
	//	uuidValidator:    uuidValidator,
	}
}

func (h Health) Route(r chi.Router) {
	r.Get("/", h.getHealthStatus)
}

func (h Health) getHealthStatus(w http.ResponseWriter, r *http.Request) {

	t.responseHandlers.Ok(w, response.ToTravelPageResponse("oi....."))
}