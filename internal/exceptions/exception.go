package exceptions

import(
	"fmt"
	"errors"
	"strings"
	"github.com/go-playground/validator/v10"
)

const (
	Prefix                  = "ERRO-001"
	InternalServerErrorCode  = Prefix + "-500"
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
	ErrInternalServerError = Exception(errors.New("an internal error has occurred, contact the SRE team"))
)

var httpErrorList = [...]HttpError{

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