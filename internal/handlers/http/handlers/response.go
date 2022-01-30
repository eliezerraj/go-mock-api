package handlers

import (
	"net/http"
	"encoding/json"
	"sync"

	"github.com/go-mock-api/internal/exceptions"
	"github.com/go-mock-api/internal/utils"
	"github.com/go-mock-api/internal/utils/loggers"
	"github.com/go-mock-api/internal/utils/constants"

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

func (h ResponseHandlerImpl) Exception(w http.ResponseWriter, r *http.Request, err error) {
	httpError := exceptions.GetHttpError(err)

	logContext, valid := utils.FindToContext(r.Context(), constants.LogContext).(*loggers.LogContext)
	if valid {
		logContext.StackTrace = httpError.StackTracer()
		utils.AddToContext(r.Context(), constants.LogContext, &logContext)
	} else {
		loggers.GetLogger().Warn("Invalid context founded")
	}

//	var requestId string
//	ctxRequest, ctxValid := utils.FindToContext(r.Context(), constants.ContextRequest).(base.ContextRequest)
//	if !ctxValid {
//		loggers.GetLogger().Warn("Invalid context founded")
//	} 

	resp := exceptions.NewErrorResponse("requestId", httpError.Exception.Error(), httpError.Code)
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
			handler := exceptions.NewErrorResponse("", e.Error(), exceptions.InternalServerErrorCode)
			bytes, _ := json.Marshal(handler)
			_, _ = w.Write(bytes)
		} else {
			_, _ = w.Write(bytes)
		}
	}
}