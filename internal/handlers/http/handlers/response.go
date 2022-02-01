package handlers

import (
	"net/http"
	"encoding/json"
	"sync"
//	"strconv"
//"fmt"

	"github.com/go-mock-api/internal/exceptions"
//	"github.com/go-mock-api/internal/utils"
//	"github.com/go-mock-api/internal/utils/loggers"
//	"github.com/go-mock-api/internal/utils/constants"
//	"github.com/go-mock-api/internal/core/model"

)

var (
	responseHandlers ResponseHandler
	onceResponse sync.Once
)

type ResponseHandlerImpl struct{}

func NewResponseHandlers() ResponseHandler {
	return ResponseHandlerImpl{}
}

type ResponseHandler interface {
	Ok(w http.ResponseWriter, data interface{})
	InternalServerError(w http.ResponseWriter, data interface{})
	Exception(w http.ResponseWriter, r *http.Request, err error)
	BadRequest(w http.ResponseWriter, r *http.Request, err error)

}

func GetResponseHandlersInstance() ResponseHandler {
	onceResponse.Do(func() {
		responseHandlers = NewResponseHandlers()
	})
	return responseHandlers
}


func (h ResponseHandlerImpl) Ok(w http.ResponseWriter, data interface{}) {
	response(w, data, http.StatusOK)
}

func (h ResponseHandlerImpl) InternalServerError(w http.ResponseWriter, data interface{}) {
	response(w, data, http.StatusInternalServerError)
}

func (h ResponseHandlerImpl) BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	data := h.GetResponseError(r, err)
	response(w, data, http.StatusBadRequest)
}

func (h ResponseHandlerImpl) Exception(w http.ResponseWriter, r *http.Request, err error) {
	httpError := exceptions.GetHttpError(err)
//	fmt.Println("=========>",httpError)
/*	logContext, valid := utils.FindToContext(r.Context(), constants.LogContext).(*loggers.LogContext)
	if valid {
		logContext.StackTrace = httpError.StackTracer()
		utils.AddToContext(r.Context(), constants.LogContext, &logContext)
	} else {
		loggers.GetLogger().Warn("Invalid context founded 1")
	}

	var requestId string
	ctxRequest, ctxValid := utils.FindToContext(r.Context(), constants.ContextRequest).(model.ContextRequest)
	if !ctxValid {
		loggers.GetLogger().Warn("Invalid context founded 2")
	} else {
		issuerId := ctxRequest.IssuerID
		requestId = ctxRequest.RequestId
		w.Header().Add("issuer_id", strconv.Itoa(int(issuerId)))
		w.Header().Add("request_id", requestId)
	}
*/
	resp := exceptions.NewErrorResponse("", httpError.Exception.Error(), httpError.Code)
	response(w, resp, httpError.HttpStatusCode)
}

func headers(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}

func response(w http.ResponseWriter, data interface{}, httpStatus int) {
	headers(w)
	w.WriteHeader(httpStatus)
	if data != nil {
		if bytes, e := json.Marshal(data); e != nil {
			w.WriteHeader(http.StatusInternalServerError)
			handler := exceptions.NewErrorResponse("", e.Error(), exceptions.SystemErrorCode)
			bytes, _ := json.Marshal(handler)
			_, _ = w.Write(bytes)
		} else {
			_, _ = w.Write(bytes)
		}
	}
}

func (h ResponseHandlerImpl) GetResponseError(r *http.Request, err error) exceptions.ErrorResponse {
	httpError := exceptions.GetHttpError(err)
	return exceptions.NewErrorResponse("", httpError.Exception.Error(), httpError.Code)
}