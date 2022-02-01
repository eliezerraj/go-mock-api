package exceptions

import(
	"fmt"
	"errors"
	"strings"
	"github.com/go-playground/validator/v10"
)

const (
	Prefix            = "SYSTEM-XPTO"
	SystemErrorCode  = Prefix + "-999"
)

type ErrorResponse struct {
	ErrorResponseDetail `json:"error"`
}

type ErrorResponseDetail struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Code        string `json:"code"`
}

func NewErrorResponse(id string, msg string, code string) ErrorResponse {
	return ErrorResponse{
		ErrorResponseDetail{
			Id:          id,
			Description: msg,
			Code:        code,
		},
	}
}

//----------------------------------
type Exception error

type HttpError struct {
	Exception      Exception
	HttpStatusCode int
	Code           string
}

const (
	oneOf    = "oneof"
	required = "required"
	min      = "min"
	max      = "max"
)

var (
	ErrInternalServerError = Exception(errors.New("ERROR => an internal error has occurred, contact the SRE team"))
	ErrList                = Exception(errors.New("ERROR => could not get any data from database"))
	ErrSave                = Exception(errors.New("ERROR => could not insert data from database"))
	ErrJsonDecode          = Exception(errors.New("ERROR => could not convert data to json"))
	ErrJsonCode            = Exception(errors.New("ERROR => could not convert json to data"))
	ErrSaveDatabase 	   = Exception(errors.New("ERROR => could not save data to database"))
	ErrNoDataFound 		   = Exception(errors.New("ERROR => no data found with parameters informed"))
	ErrTokenUnreachable    = Exception(errors.New("ERROR => unable to access the token"))
)

var httpErrorList = [...]HttpError{
	{Exception: ErrInternalServerError, HttpStatusCode: 500, Code: SystemErrorCode},
	{Exception: ErrList, HttpStatusCode: 400, Code: SystemErrorCode},
	{Exception: ErrSave, HttpStatusCode: 400, Code: SystemErrorCode},
	{Exception: ErrJsonDecode, HttpStatusCode: 500, Code: SystemErrorCode},
	{Exception: ErrJsonCode, HttpStatusCode: 500, Code: SystemErrorCode},
	{Exception: ErrSaveDatabase, HttpStatusCode: 500, Code: SystemErrorCode},
	{Exception: ErrNoDataFound, HttpStatusCode: 404, Code: SystemErrorCode},
	{Exception: ErrTokenUnreachable, HttpStatusCode: 401, Code: SystemErrorCode},
}

func Throw(old error, new error) error {
	new = fmt.Errorf("%w", old)
	return new
}

func IsValidationError(err error) (bool, *HttpError) {
	v := &validator.ValidationErrors{}
	var validateErr HttpError
	if errors.As(err, v) {
		for _, err := range err.(validator.ValidationErrors) {
			validateErr = generateValidateHttpError(err)
			break
		}
		return true, &validateErr
	}
	return false, nil
}

func GetHttpError(err error) HttpError {
	if ok, httpError := IsValidationError(err); ok {
		return *httpError
	}
	aux := strings.Split(err.Error(), ":")[0]
	for _, e := range httpErrorList {
		if e.Exception.Error() == aux {
			return e
		}
	}
	return HttpError{
		Exception:      ErrInternalServerError,
		HttpStatusCode: 500,
		Code:           "500",
	}
}

func generateValidateHttpError(err validator.FieldError) HttpError {
	switch err.Tag() {
	case oneOf:
		m := fmt.Sprintf("%s  informed is not valid", err.Field())
		return buildValidationError(m, "500")
	case required:
		m := fmt.Sprintf("Field %s is required", err.Field())
		return buildValidationError(m, "500")
	case min:
		m := fmt.Sprintf("The size of field %s is required", err.Field())
		return buildValidationError(m, "500")
	case max:
		m := fmt.Sprintf("Field %s is required", err.Field())
		return buildValidationError(m, "500")
	default:
		m := fmt.Sprintf("Validation error for '%s'", err.Field())
		return buildValidationError(m, "500")
	}
	return HttpError{}
}

func buildValidationError(msg, code string) HttpError {
	ErrValidateException := Exception(errors.New(msg))
	return HttpError{Exception: ErrValidateException, HttpStatusCode: 400, Code: code}
}

func (e HttpError) StackTracer() string {
	return fmt.Sprintf("%+v", e.Exception)
}