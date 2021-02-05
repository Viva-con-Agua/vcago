package verr

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

//LogLevel can be set by SetLogLevel. Default value is debug
var LogLevel = "debug"

//SetLogLevel can be used for set LogLevel variable
func SetLogLevel(l string) {
    LogLevel = l
}


//LogError prints error in service logs
func LogError(ctx context.Context, err error, lvl string) {
	if LogLevel == "debug" {
		printError(ctx, err)
	} else if LogLevel == "prod"{
		if lvl == "prod" {
			printError(ctx, err)
		}
	}
}

//printError log error to console.
func printError(ctx context.Context, err error) {
	//get infos about function
	pc := make([]uintptr, 10)
	runtime.Callers(4, pc)
	function := runtime.FuncForPC(pc[0])

	//fill ApiError
	_, line := function.FileLine(pc[0])
	file := runtime.FuncForPC(pc[0]).Name()   //get infos about function
	log.Print(
			"\n",
			string(ColorRed), "Error Message: \n",
			"\t", string(ColorWhite), err.Error(), "\n",
			"\tFile: [", file, "]\n",
			"\tLine: [", line, "]\n",
			//string(colorYellow), "Session_User: "+string(colorWhite)+"\n", u, string(user), "\n",
			string(ColorYellow), "Request_Header: ", string(ColorWhite), "\n\t",
			//formatRequest(c.Request()), "\n",
			//string(colorYellow), "Request_Body: ", string(colorWhite), "\n",

			string(ColorBlue), "### END ERROR", string(ColorWhite), "\n\n",
		)

}


//formatRequest format string output from formatRequest
func formatRequest(r *http.Request) string {
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
