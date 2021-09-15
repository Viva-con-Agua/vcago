package vcago

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
)

//Mongo represents the initial struct for an Mongo connection.
type IDjango struct {
	URL    string
	Key    string
	Export bool
}

//LoadEnv loads the Host and Port From .env file.
//Host can be set via NATS_HOST
//Port can be set via NATS_PORT
func (i *IDjango) LoadEnv() {
	i.URL = Config.GetEnvString("IDJANGO_URL", "w", "https://idjangostage.vivaconagua.org")
	i.Key = Config.GetEnvString("IDJANGO_KEY", "w", "")
	i.Export = Config.GetEnvBool("IDJANGO_EXPORT", "w", false)
}

func (i *IDjango) Post(data interface{}, path string) (err error) {
	if i.Export {
		var jsonData []byte
		if jsonData, err = json.Marshal(data); err != nil {
			return
		}
		log.Print(i.URL + path)
		request := new(http.Request)
		request, err = http.NewRequest("POST", i.URL+path, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		request.Header.Set("Authorization", "Api-Key "+i.Key)
		client := &http.Client{}
		response := new(http.Response)
		response, err = client.Do(request)
		if err != nil {
			return NewIDjangoError(err, 500, nil)
		}
		defer response.Body.Close()
		if response.StatusCode != 201 {
			var bodyBytes []byte
			if bodyBytes, err = ioutil.ReadAll(response.Body); err != nil {
				return NewIDjangoError(err, response.StatusCode, nil)
			}
			body := new(interface{})
			if err = json.Unmarshal(bodyBytes, body); err != nil {
				return NewIDjangoError(err, 500, string(bodyBytes))
			}
			return NewIDjangoError(nil, response.StatusCode, body)
		}

	}
	return
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

type IDjangoError struct {
	Err     error       `json:"error" bson:"error"`
	Message string      `json:"message" bson:"message"`
	Code    int         `json:"code" bson:"code"`
	Body    interface{} `json:"body" bson:"body"`
	Line    int         `json:"line" bson:"line"`
	File    string      `json:"file" bson:"file"`
}

func (i *IDjangoError) Error() string {
	res, _ := json.Marshal(i)
	return string(res)
}
