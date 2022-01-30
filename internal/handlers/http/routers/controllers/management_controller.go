package controllers

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/go-mock-api/internal/viper"
	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/handlers/http/handlers"
	"github.com/go-mock-api/internal/handlers/http/model"
	"github.com/go-mock-api/internal/services"

)

type Management struct {
	requestHandlers  handlers.RequestHandlers
	responseHandlers handlers.ResponseHandler
}

func NewManagementController(requestHandlers handlers.RequestHandlers,
							responseHandlers handlers.ResponseHandler) Management {
	return Management{	requestHandlers:  requestHandlers,
						responseHandlers: responseHandlers,
	}
}

func (m Management) GetPath() string {
	return constants.ManagementPath
}

func (m Management) Route(r chi.Router) {
	r.Get("/health", m.checkHealth)
	r.Get("/info", m.getInfo)
}

func (m Management) checkHealth(w http.ResponseWriter, _ *http.Request) {
	resp := services.CheckHealth()
	if resp.Status != constants.UP {
		m.responseHandlers.InternalServerError(w, resp)
		return
	}
	m.responseHandlers.Ok(w, resp)
}

func (m Management) getInfo(w http.ResponseWriter, _ *http.Request) {
	git := model.ManagerInfoResponse{
		App: &model.ManagerInfoResponseApp{
			Name:        viper.Application.App.Name,
			Description: viper.Application.App.Description,
			Version:     viper.Application.App.Version,
		},
	}
	m.responseHandlers.Ok(w, git)
}
