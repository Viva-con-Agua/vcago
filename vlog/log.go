package vlog

import (
	"net/http"

	"github.com/Viva-con-Agua/vcago/verror"
)

type (
	//Logger represents a logging handler for handling vcago errors.
	Logger struct {
		natsIP   string
		natsPORT string
		mode     string
	}
)

//New creates a new Logger
func New(natsIP string, natsPORT string, mode string) *Logger {
	return &Logger{
		natsIP:   natsIP,
		natsPORT: natsPORT,
		mode:     mode,
	}
}

var global = New("-", "-", "local")

//Print handle logging process.
//Provide Message.Print(), Message.PrintPretty()
func Print(request *http.Request, err *verror.Error) {
	switch global.mode {
	case "local":
		NewMessage(request, err).Print()
	case "local-pretty":
		NewMessage(request, err).PrintPretty()
	}
}

//Config loads a configuration for the global Logger
func Config(natsIP string, natsPORT string, mode string) {
	global = New(natsIP, natsPORT, mode)
}
