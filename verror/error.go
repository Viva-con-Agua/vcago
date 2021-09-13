package verror

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	//Response represents an response json in case of an error.
	Response struct {
		Status  int         `json:"-"`
		Message string      `json:"message"`
		Body    interface{} `json:"body,omitempty"`
		Coll    string      `json:"collection,omitempty"`
		Model   string      `json:"model,omitempty"`
	}
)

//type ErrorHandler map[error]error

//func (i *ErrorHandler) Init() *ErrorHandler {
//	i := make[DefaultError]
//}

//func (i *ErrorHandler) Set(key error, value error) *ErrorHandler {
//	i[key] = value
//}

//HTTPErrorHandler handles echo.HTTPError and return the correct response.
func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	response := new(interface{})
	he, ok := err.(*echo.HTTPError)
	if ok {
		code = he.Code
		response = &he.Message
	}
	if code == http.StatusInternalServerError {
		c.JSON(code, Response{Message: "internal_server_error"})
	} else {
		c.JSON(code, response)
	}
}

//InternalServerError creates an echo.HTTPError with the status http.StatusInternalServerError
func InternalServerError(err error) error {
	return echo.NewHTTPError(http.StatusInternalServerError, err)
}

//BadRequest creates an echo.HTTPError with the status http.StatusBadRequest
func BadRequest(message string, body ...interface{}) error {
	return echo.NewHTTPError(http.StatusBadRequest, Response{Message: message, Body: body})
}

//Conflict creates an echo.HTTPError with the status http.StatusConflict
func Conflict(message string, body ...interface{}) error {
	return echo.NewHTTPError(http.StatusConflict, &Response{Message: message, Body: body})
}

//NotFound creates an echo.HTTPError with the status http.StatusNotFound
func NotFound(message string, body ...interface{}) error {
	return echo.NewHTTPError(http.StatusNotFound, &Response{Message: message, Body: body})
}
