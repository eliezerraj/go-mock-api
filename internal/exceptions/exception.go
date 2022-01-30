package exceptions

import(
	
)

const (
	PrefixCards                  = "CARDS-001"
	CardInternalServerErrorCode  = PrefixCards + "-500"
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