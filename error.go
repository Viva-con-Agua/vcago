package vcago

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Error struct {
	ID      string `json:"id"`
	Time    string `json:"time"`
	Level   string `json:"level"`
	File    string `json:"file"`
	Line    int    `json:"line"`
	Message string `json:"message"`
	Err     error  `json:"-"`
	Model   string `json:"model,omitempty"`
	Type    string `json:"type"`
}

var LogLevel = Settings.String("LOG_LEVEL", "w", "DEBUG")

func (i *Error) Log() string {
	err, _ := json.Marshal(i)
	return string(err)
}

func (i *Error) Error() string {
	return i.Err.Error()
}

func NewError(err error, lvl string, t string) *Error {
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	_, line := f.FileLine(pc[0])
	file := runtime.FuncForPC(pc[0]).Name()
	return &Error{
		Time:    time.Now().String(),
		Level:   lvl,
		File:    file,
		Line:    line,
		Message: err.Error(),
		Err:     err,
		Type:    t,
	}
}

func (i *Error) AddModel(model string) *Error {
	i.Model = model
	return i
}

func (i *Error) Print(id string) {
	i.ID = id
	if LogLevel == "DEBUG" {
		fmt.Println(i.Log())
	} else if LogLevel == "ERROR" {
		if i.Level == "ERROR" {
			fmt.Println(i.Log())
		}
	}
}

// MongoErrorResponseHandler handles the response for the MongoError type.
func (i *Error) Response() (int, interface{}) {
	switch i.Type {
	case "mongo":
		return i.MongoResponse()
	case "bind":
		return i.BindResponse()
	case "validation":
		return i.ValidationResponse()
	default:
		return NewInternalServerError(i.Model).Response()
	}
}

type MongoDuplicatedErrorResponse struct {
	Response
	Type    string `example:"error"`
	Message string `example:"duplicate key error: <key>"`
}

type MongoNoDocumentErrorResponse struct {
	Response
	Type    string `example:"error"`
	Message string `example:"document not found"`
}

func (i *Error) MongoResponse() (int, interface{}) {
	if strings.Contains(i.Message, "duplicate key error") {
		temp := strings.Split(i.Message, "key: {")
		temp = strings.Split(temp[1], "}")
		response := &Response{
			Status:  http.StatusConflict,
			Type:    "error",
			Message: "duplicate key error: " + temp[0],
			Model:   i.Model,
		}
		return response.Response()
	}
	switch i.Err {
	case mongo.ErrNoDocuments:
		response := &Response{
			Status:  http.StatusNotFound,
			Type:    "error",
			Message: "document not found",
			Model:   i.Model,
		}
		return response.Response()
	default:
		return NewInternalServerError(i.Model).Response()
	}
}

type BindErrorResponse struct {
	Response
	Type    string   `example:"error"`
	Message string   `example:"bind error"`
	Payload []string `example:"cant bind string to int"`
}

func (i *Error) BindResponse() (int, interface{}) {
	response := new(ValidationError)
	response.Bind(i.Err)
	return NewBadRequest(i.Model, "bind error", response).Response()
}

type ValidationErrorResponse struct {
	Response
	Type    string   `example:"error"`
	Message string   `example:"validate error"`
	Payload []string `example:"currency is required"`
}

func (i *Error) ValidationResponse() (int, interface{}) {
	response := new(ValidationError)
	response.Valid(i.Err)
	return NewBadRequest(i.Model, "validation error", response).Response()
}
