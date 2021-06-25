package verror

import (
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	//Error represents an error for vcago
	Error struct {
		File       string `json:"file"`
		Line       int    `json:"line"`
		Error      error  `json:"-"`
		ErrorMsg   string `json:"error_msg"`
		Time       int64  `json:"time"`
		TimeString string `json:"time_string"`
	}
	//ErrorResponse represents an response json in case of an error.
	ErrorResponse struct {
		Message string `json:"message"`
		Coll    string `json:"collection,omitempty"`
	}
)

var InternalServerErrorMsg = echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "internal_server_error"})

//New creates new error
func New(err error) *Error {
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	function := runtime.FuncForPC(pc[0])
	_, line := function.FileLine(pc[0])
	file := runtime.FuncForPC(pc[0]).Name()
	var now = time.Now().Unix()
	return &Error{
		Line:       line,
		File:       file,
		Error:      err,
		ErrorMsg:   err.Error(),
		Time:       now,
		TimeString: time.Unix(now, 0).String(),
	}
}

//ErrorResponse handle http response in error case
func (verr *Error) ErrorResponse(collection ...string) (int, interface{}) {
	response := new(ErrorResponse)
	if collection != nil {
		response.Coll = collection[0]
	}
	if strings.Contains(verr.Error.Error(), "duplicate key error") {
		response.Message = "duplicate key error"
		return http.StatusConflict, response
	}
	if verr.Error == mongo.ErrNoDocuments {
		response.Message = "model not found"
		return http.StatusNotFound, response
	}
	if verr.Error == ErrMongoUpdate {
		response.Message = "no updated document"
		return http.StatusNotFound, response
	}
	if verr.Error == ErrMongoDelete {
		response.Message = "no deleted document"
		return http.StatusNotFound, response
	}
	return http.StatusInternalServerError, ErrorResponse{Message: "internal_server_error"}

}
