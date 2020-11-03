package vcago

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/Viva-con-Agua/vcago/vmod"
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

func LogApiError(e *verr.ApiError, c echo.Context, i interface{}) {
	u := c.Get("user")
	var u_string string
	if user, ok := u.(*vmod.User); ok {
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
		formatRequestPrint(c.Request()), "\n",
		string(colorYellow), "Request_Body: ", string(colorWhite), "\n",
		string(body), "\n",
		string(colorBlue), "### END ERROR", string(colorWhite), "\n\n",
	)

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
