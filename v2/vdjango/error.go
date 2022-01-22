package vdjango

import (
	"encoding/json"
	"runtime"
)

type IDjangoError struct {
	Err     error       `json:"error" bson:"error"`
	Message string      `json:"message" bson:"message"`
	Code    int         `json:"code" bson:"code"`
	Body    interface{} `json:"body" bson:"body"`
	Line    int         `json:"line" bson:"line"`
	File    string      `json:"file" bson:"file"`
}

func NewIDjangoError(err error, code int, body interface{}) *IDjangoError {
	var message = ""
	if err != nil {
		message = err.Error()
	}
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	_, line := f.FileLine(pc[0])
	file := runtime.FuncForPC(pc[0]).Name()
	return &IDjangoError{
		Err:     err,
		Message: message,
		Code:    code,
		Body:    body,
		Line:    line,
		File:    file,
	}
}

func (i *IDjangoError) Error() string {
	res, _ := json.Marshal(i)
	return string(res)
}
