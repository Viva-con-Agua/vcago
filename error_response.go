package vcago

import (
	"net/http"
)

type (
	//Response represents an response json in case of an error.
	ErrorResponse struct {
		Status    int         `json:"-"`
		ErrorType string      `json:"error_type"`
		Message   string      `json:"message"`
		Payload   interface{} `json:"body,omitempty"`
	}
)

//InternalServerError creates an echo.HTTPError with the status http.StatusInternalServerError
func InternalServerError() (int, ErrorResponse) {
	return http.StatusInternalServerError, ErrorResponse{Message: "internal_server_error"}
}

//NewErrorResponse creates an new error response.
func NewErrorResponse(status int, errorType string, message string, payload ...interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status:    status,
		ErrorType: errorType,
		Message:   message,
		Payload:   payload,
	}
}

func (i *ErrorResponse) Error() string {
	return "error_type: " + i.ErrorType + "; message: " + i.Message + ";"
}

func (i *ErrorResponse) Response() (int, *ErrorResponse) {
	return i.Status, i
}

func NewBadRequest(message string, payload ...interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status:    http.StatusBadRequest,
		ErrorType: "bad request",
		Message:   message,
		Payload:   payload,
	}
}

func NewConflict(message string, payload ...interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status:    http.StatusConflict,
		ErrorType: "conflict",
		Message:   message,
		Payload:   payload,
	}
}

func NewNotFound(message string, payload ...interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status:    http.StatusNotFound,
		ErrorType: "not found",
		Message:   message,
		Payload:   payload,
	}
}

//BadRequest creates an echo.HTTPError with the status http.StatusBadRequest
func BadRequest(message string, body ...interface{}) (int, *ErrorResponse) {
	return http.StatusBadRequest, &ErrorResponse{Message: message, Payload: body}
}

//Conflict creates an echo.HTTPError with the status http.StatusConflict
func Conflict(message string, payload ...interface{}) (int, *ErrorResponse) {
	return http.StatusConflict, &ErrorResponse{Message: message, Payload: payload}
}

//NotFound creates an echo.HTTPError with the status http.StatusNotFound
func NotFound(message string, payload ...interface{}) (int, *ErrorResponse) {
	return http.StatusNotFound, &ErrorResponse{Message: message, Payload: payload}
}
