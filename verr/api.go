package verr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/labstack/echo/v4"
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
	//APIError represents the return value of all database controllers in case of an error.
	APIError struct {
		Code  int
		Body  interface{}
		Error error
		Line  int
		File  string
		Level bool
	}
	//ResponseError represents the json body of an error response.
	ResponseError struct {
		Message string                 `json:"message"`
		Data    map[string]interface{} `json:"data,omitempty"`
	}
)

//NewAPIError catches there context of calling and creates an APIError contains this informations.
//File is set to the name of the file and Line to the line in file NewAPIError was called from.
func NewAPIError(err error) *APIError {
	//new ApiError
	a := new(APIError)

	//get infos about function
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])

	//fill ApiError
	a.Error = err
	_, a.Line = f.FileLine(pc[0])
	a.File = runtime.FuncForPC(pc[0]).Name()
	return a
}

//Conflict extend APIError for http status 409 Conflict and adds a message m that can response to client.
func (a *APIError) Conflict(m string) *APIError {
	a.Code = http.StatusConflict
	a.Body = ResponseError{Message: m}
	a.Level = false
	return a
}

//Forbidden extend APIError for http status  Forbidden and adds a message m that can response to client.
func (a *APIError) Forbidden(m string) *APIError {
	a.Code = http.StatusForbidden
	a.Body = ResponseError{Message: m}
	a.Level = false
	return a
}

//Unauthorized extend APIError for http status 401 Unauthorized and adds a message m that can response to client.
func (a *APIError) Unauthorized() *APIError {
	a.Code = http.StatusUnauthorized
	a.Body = ResponseError{Message: "unauthorized"}
	a.Level = false
	return a
}

//NotFound extend APIError for http status 404 Not Found and adds a message m that can response to client.
func (a *APIError) NotFound(m string) *APIError {
	a.Code = http.StatusNotFound
	a.Body = ResponseError{Message: m}
	a.Level = false
	return a
}

//InternalServerError extend APIError for http status 500 Internal Server Error
func (a *APIError) InternalServerError() *APIError {
	a.Code = http.StatusInternalServerError
	a.Body = ResponseError{Message: "internal_server_error"}
	a.Level = true
	return a
}

//Log print APIError to server logs
func (a *APIError) Log(c echo.Context) {
	if a.Level == true || os.Getenv("LOG_LEVEL") == "debug" {
		u := c.Get("user")
		user, _ := json.MarshalIndent(u, "", "\t")
		log.Print(
			"\n",
			string(colorRed), "Error Message: \n",
			"\t", string(colorWhite), a.Error.Error(), "\n",
			"\tFile: [", a.File, "]\n",
			"\tLine: [", a.Line, "]\n",
			string(colorYellow), "Session_User: "+string(colorWhite)+"\n", u, string(user), "\n",
			string(colorYellow), "Request_Header: ", string(colorWhite), "\n\t",
			formatRequestPrint(c.Request()), "\n",
			string(colorYellow), "Request_Body: ", string(colorWhite), "\n",

			string(colorBlue), "### END ERROR", string(colorWhite), "\n\n",
		)
	}
}

func formatRequestPrint(r *http.Request) string {
	// Create return string
	var request []string // Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)                             // Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host)) // Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}
	return strings.Join(request, "\n\t")

}
