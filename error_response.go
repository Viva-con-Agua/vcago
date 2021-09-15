package vcago

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	//Response represents an response json in case of an error.
	ErrorResponse struct {
		Status  int         `json:"-"`
		Message string      `json:"message"`
		Body    interface{} `json:"body,omitempty"`
		Coll    string      `json:"collection,omitempty"`
		Model   string      `json:"model,omitempty"`
	}
)

//HTTPErrorHandler handles echo.HTTPError and return the correct response.
func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	response := new(interface{})
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		response = &he.Message
		if he.Code == http.StatusInternalServerError {
			c.JSON(code, ErrorResponse{Message: "internal_server_error"})
		} else {
			res := new(interface{})
			json.Unmarshal([]byte((*response).(string)), res)
			c.JSON(code, res)
		}
	} else if resp, ok := err.(*MongoError); ok {
		c.JSON(resp.Response())
	} else {
		c.JSON(code, ErrorResponse{Message: "internal_server_error"})
	}
}

//InternalServerError creates an echo.HTTPError with the status http.StatusInternalServerError
func InternalServerError() (int, ErrorResponse) {
	return http.StatusInternalServerError, ErrorResponse{Message: "internal_server_error"}
}

//BadRequest creates an echo.HTTPError with the status http.StatusBadRequest
func BadRequest(message string, body ...interface{}) (int, *ErrorResponse) {
	return http.StatusBadRequest, &ErrorResponse{Message: message, Body: body}
}

//Conflict creates an echo.HTTPError with the status http.StatusConflict
func Conflict(message string, body ...interface{}) (int, *ErrorResponse) {
	return http.StatusConflict, &ErrorResponse{Message: message, Body: body}
}

//NotFound creates an echo.HTTPError with the status http.StatusNotFound
func NotFound(message string, body ...interface{}) (int, *ErrorResponse) {
	return http.StatusNotFound, &ErrorResponse{Message: message, Body: body}
}
