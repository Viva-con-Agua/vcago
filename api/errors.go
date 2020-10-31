package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	ApiError struct {
		Error     error
		Model     string
		Line      int
		FileName  string
		UserEmail string
		UserUuid  string
	}
	JsonError struct {
		Key   string
		Error string
	}
)

var (
	ErrorNotFound     = errors.New("NotFound")
	ErrorConflict     = errors.New("Conflict")
	ErrorPassword     = errors.New("Password")
	ErrorUserNotFound = errors.New("user not found")
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

func FormatRequestPrint(r *http.Request) string {
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
func GetError(err error) *ApiError {
	//new ApiError
	api_error := new(ApiError)
	//get infos about function
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	//fill ApiError
	api_error.Error = err
	_, api_error.Line = f.FileLine(pc[0])
	api_error.FileName = runtime.FuncForPC(pc[0]).Name()
	return api_error
}

func (e *ApiError) PrintErrorNoSession(c echo.Context, i interface{}) {
	error_list := strings.Split(e.Error.Error(), ": ")
	body, _ := json.MarshalIndent(i, "", "\t")
	log.Print(
		"\n",
		string(colorRed), error_list[0], ": \n",
		"\t", string(colorWhite), error_list[1], "\n",
		"\tFile: [", e.FileName, "]\n",
		"\tLine: [", e.Line, "]\n",
		string(colorYellow), "Request_Header: ", string(colorWhite), "\n\t",
		FormatRequestPrint(c.Request()), "\n",
		string(colorYellow), "Request_Body: ", string(colorWhite), "\n",
		string(body), "\n",
		string(colorBlue), "### END ERROR", string(colorWhite), "\n\n",
	)

}

func (e *ApiError) LogError(c echo.Context, i interface{}) {
	u := c.Get("user")
	var u_string string
	if user, ok := u.(*User); ok {
		u_string = string(colorYellow) + "Session_Email: " + string(colorWhite) + user.Email + "\n" +
			string(colorYellow) + "Session_UUID: " + string(colorWhite) + user.ID + "\n"
	} else {
		u_string = ""
	}
	body, _ := json.MarshalIndent(i, "", "\t")
	log.Print(
		"\n",
		string(colorRed), "Error Message: \n",
		"\t", string(colorWhite), e.Error.Error(), "\n",
		"\tFile: [", e.FileName, "]\n",
		"\tLine: [", e.Line, "]\n",
		u_string,
		string(colorYellow), "Request_Header: ", string(colorWhite), "\n\t",
		FormatRequestPrint(c.Request()), "\n",
		string(colorYellow), "Request_Body: ", string(colorWhite), "\n",
		string(body), "\n",
		string(colorBlue), "### END ERROR", string(colorWhite), "\n\n",
	)

}

func JsonErrorResponse(e error) (j_list []JsonError) {
	json_error := new(JsonError)
	error_list := strings.Split(e.Error(), "\n")
	for _, val := range error_list {
		e_list := strings.Split(val, "Key: ")
		e_list = strings.Split(e_list[1], " Error:")
		json_error.Key = e_list[0]
		json_error.Error = e_list[1]
		j_list = append(j_list, *json_error)
	}
	return j_list

}
