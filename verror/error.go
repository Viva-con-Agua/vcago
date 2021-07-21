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
		File          string        `json:"file"`
		Line          int           `json:"line"`
		Error         error         `json:"-"`
		ErrorMsg      string        `json:"error_msg"`
		Time          int64         `json:"time"`
		TimeString    string        `json:"time_string"`
		ErrorResponse ErrorResponse `json:"response"`
	}
	//ErrorResponse represents an response json in case of an error.
	ErrorResponse struct {
		Status  int         `json:"-"`
		Message string      `json:"message"`
		Body    interface{} `json:"body,omitempty"`
		Coll    string      `json:"collection,omitempty"`
		Model   string      `json:"model,omitempty"`
	}
)

var InternalServerError = echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{Message: "internal_server_error"})

//New creates new error
func New(err error, coll ...string) *Error {
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

func BadRequestResponse(message string) error {
	return echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{Message: message})
}

func (verr *Error) Body(body interface{}) *Error {
	verr.ErrorResponse.Body = body
	return verr
}

func (verr *Error) Status(status int) *Error {
	verr.ErrorResponse.Status = status
	return verr
}

func (verr *Error) Response() (int, interface{}) {
	return verr.ErrorResponse.Status, verr.ErrorResponse
}

func (verr *Error) Mongo(coll string) *Error {
	response := new(ErrorResponse)
	response.Coll = coll
	if strings.Contains(verr.ErrorMsg, "duplicate key error") {
		response.Message = "duplicate key error"
		response.Status = http.StatusConflict
	}
	if verr.Error == mongo.ErrNoDocuments {
		response.Message = "model not found"
		response.Status = http.StatusNotFound
	}
	if verr.Error == ErrMongoUpdate {
		response.Message = "no updated document"
		response.Status = http.StatusNotFound
	}
	if verr.Error == ErrMongoDelete {
		response.Message = "no deleted document"
		response.Status = http.StatusNotFound
	}
	verr.ErrorResponse = *response
	return verr
}

/*
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

}*/
