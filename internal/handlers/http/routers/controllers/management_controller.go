package controllers

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/go-mock-api/internal/viper"
	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/handlers/http/handlers"
	"github.com/go-mock-api/internal/core/model"
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
	result := services.CheckHealth()
	if result.Status != constants.UP {
		m.responseHandlers.InternalServerError(w, result)
		return
	}
	m.responseHandlers.Ok(w, ToManagerHealthResponse(result))
}

func (m Management) getInfo(w http.ResponseWriter, _ *http.Request) {
	result := model.ManagerInfo{
		App: &model.ManagerInfoApp{
			Name:        viper.Application.App.Name,
			Description: viper.Application.App.Description,
			Version:     viper.Application.App.Version,
		},
	}
	m.responseHandlers.Ok(w, result)
}
//-----------------------------
func ToManagerHealthDBResponse(m model.ManagerHealthDB) model.ManagerHealthDB {
	return model.ManagerHealthDB{
		Status:        	m.Status,
	}
}

func ToManagerHealthDiskSpaceResponse(m model.ManagerHealthDiskSpace) model.ManagerHealthDiskSpace {
	return model.ManagerHealthDiskSpace{
		Status:        	m.Status,
		Total:			m.Total,
		Free:			m.Free,
		Threshold:		m.Threshold,
	}
}

func ToManagerHealthResponse(m model.ManagerHealth) model.ManagerHealth {
	return model.ManagerHealth{
		Status:        	m.Status,
		DB:				ToManagerHealthDBResponse(m.DB),
		DiskSpace:		ToManagerHealthDiskSpaceResponse(m.DiskSpace),
	}
}
