package verr

import (
	"runtime"
)

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

type (
	ApiError struct {
		Error         error
		ResponseError *ResponseError
		Line          int
		FileName      string
	}
)

func GetApiError(err error, r_err *ResponseError) *ApiError {
	//new ApiError
	api_error := new(ApiError)

	//get infos about function
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])

	//fill ApiError
	api_error.Error = err
	api_error.ResponseError = r_err
	_, api_error.Line = f.FileLine(pc[0])
	api_error.FileName = runtime.FuncForPC(pc[0]).Name()
	return api_error
}
