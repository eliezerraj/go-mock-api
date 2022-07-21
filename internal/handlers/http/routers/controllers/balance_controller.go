package controllers

import (
	"net/http"
	"github.com/go-chi/chi"

	"github.com/go-mock-api/internal/utils/constants"
	"github.com/go-mock-api/internal/handlers/http/handlers"
	"github.com/go-mock-api/internal/handlers/http/middleware"
	"github.com/go-mock-api/internal/core/model"
	"github.com/go-mock-api/internal/services"
	"github.com/go-mock-api/internal/utils/loggers"
)

type Balance struct {
	requestHandlers  handlers.RequestHandlers
	responseHandlers handlers.ResponseHandler
	service          services.BalanceService
	validator    	 middleware.ValidatorMiddleware
}

func NewBalanceController(	requestHandlers handlers.RequestHandlers,
							responseHandlers handlers.ResponseHandler,
							service services.BalanceService,
							validator middleware.ValidatorMiddleware,
							 ) Balance {
		return Balance{	requestHandlers:  requestHandlers,
						responseHandlers: responseHandlers,
						service:          service,
						validator:    	  validator,
						}		
}

func (b Balance) GetPath() string {
	return constants.BalancePath
}

func (b Balance) Route(r chi.Router) {
	r.Use(b.validator.Validate())
	r.Get("/list", b.listBalance)
	r.Route("/list_by_id/id={id}&sk={sk}", func(rRouter chi.Router) {
		rRouter.Get("/", b.listById)
	})
	r.Post("/save", b.saveBalance)
	r.Route("/id={id}", func(rRouter chi.Router) {
		rRouter.Get("/", b.findById)
	})
}

func (b Balance) listBalance(w http.ResponseWriter, r *http.Request) {
	loggers.GetLogger().Named(constants.Controller).Info("listBalance") 
	result, err := b.service.List(r.Context())
	if err != nil {
		b.responseHandlers.Exception(w, r, err)
		return
	}
	b.responseHandlers.Ok(w, ToBalanceListResponse(result))
}

func (b Balance) saveBalance(w http.ResponseWriter, r *http.Request) {
	loggers.GetLogger().Named(constants.Controller).Info("saveBalance") 
	var balance model.Balance
	err := b.requestHandlers.BindJson(r, &balance)
	if err != nil {
		b.responseHandlers.Exception(w, r, err)
		return
	}

	result, err := b.service.Save(r.Context(), balance)
	if err != nil {
			b.responseHandlers.Exception(w, r, err)
			return
	}
		b.responseHandlers.Ok(w, ToBalanceResponse(result))
}

func (b Balance) findById(w http.ResponseWriter, r *http.Request) {
	loggers.GetLogger().Named(constants.Controller).Info("findById") 
	id := b.requestHandlers.GetURLParam(r, constants.PathParamDefault)

	result, err := b.service.FindById(r.Context(), id)
	
	if err != nil {
		b.responseHandlers.Exception(w, r, err)
		return
	}
	b.responseHandlers.Ok(w, ToBalanceResponse(result))
}

func (b Balance) listById(w http.ResponseWriter, r *http.Request) {
	loggers.GetLogger().Named(constants.Controller).Info("listById") 

	id := b.requestHandlers.GetURLParam(r, constants.PathParamDefault)
	sk := b.requestHandlers.GetURLParam(r, constants.PathParamSk)

	balance := model.Balance{}
	balance.Id = id
	balance.Account = sk

	result, err := b.service.ListById(r.Context(), balance)
	
	if err != nil {
		b.responseHandlers.Exception(w, r, err)
		return
	}
	b.responseHandlers.Ok(w, ToBalanceListResponse(result))
}

//-----------------------
func ToBalanceListResponse(t []model.Balance) []model.Balance {
	list := make([]model.Balance, 0)
	for _, v := range t {
		list = append(list, ToBalanceResponse(v))
	}
	return list
}

func ToBalanceResponse(b model.Balance) model.Balance {
	return model.Balance{
		Id:        	b.Id,
		Account: 	b.Account,
		Amount: 	b.Amount,
		DateBalance: b.DateBalance,
		Description: b.Description ,
	}
}