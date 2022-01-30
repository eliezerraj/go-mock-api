package controllers

import (
	"net/http"

	"github.com/go-chi/chi"

	//"github.com/go-mock-api/internal/viper"
	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/handlers/http/handlers"
	//"github.com/go-mock-api/internal/handlers/http/model"
//	model_core "github.com/go-mock-api/internal/core/model"
	"github.com/go-mock-api/internal/services"

)

type Balance struct {
	requestHandlers  handlers.RequestHandlers
	responseHandlers handlers.ResponseHandler
	service          services.BalanceService
}

func NewBalanceController(	requestHandlers handlers.RequestHandlers,
							responseHandlers handlers.ResponseHandler,
							service services.BalanceService ) Balance {
		return Balance{	requestHandlers:  requestHandlers,
						responseHandlers: responseHandlers,
						service:          service,
						}		
}

func (b Balance) GetPath() string {
	return constants.BalancePath
}

func (b Balance) Route(r chi.Router) {
	r.Get("/list", b.listBalance)
}

func (b Balance) listBalance(w http.ResponseWriter, r *http.Request) {

//	result := []model_core.Balance{}
//	m1 := model_core.Balance{ Id: "001"}
//	result = append(result, m1)
	
	result2, err := b.service.List(r.Context())
	if err != nil {
		b.responseHandlers.Exception(w, r, err)
		return
	}

	b.responseHandlers.Ok(w, result2)
}
