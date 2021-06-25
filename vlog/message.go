package vlog

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Viva-con-Agua/vcago/verror"
)

type (
	//Request represents a http.Request as json.
	Request struct {
		Method string              `json:"method"`
		URL    string              `json:"url"`
		Proto  string              `json:"proto"`
		Header map[string][]string `json:"header"`
	}
	//Message represents a massage that can be logged
	Message struct {
		Request Request      `json:"request"`
		Error   verror.Error `json:"error"`
	}
)

//NewMessage creates a new Message from http.Request and verror.Error.
func NewMessage(request *http.Request, err *verror.Error) *Message {
	header := make(map[string][]string)
	for name, headers := range request.Header {
		header[name] = headers
	}
	r := &Request{
		Method: request.Method,
		URL:    request.URL.String(),
		Proto:  request.Proto,
		Header: header,
	}
	return &Message{
		Request: *r,
		Error:   *err,
	}
}

//Print prints the message to the io.
func (m *Message) Print() {
	r, _ := json.Marshal(m.Request)
	e, _ := json.Marshal(m.Error)
	fmt.Println("---###---")
	fmt.Println("Request: " + string(r))
	fmt.Println("Error: " + string(e))
}

//PrintPretty pretty print the Message json
func (m *Message) PrintPretty() {
	p, _ := json.MarshalIndent(m, "", "\t")
	fmt.Println(string(p))
}
