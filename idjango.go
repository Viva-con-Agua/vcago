package vcago

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"runtime"
	"time"
)

// Mongo represents the initial struct for an Mongo connection.
type IDjangoHandler struct {
	URL    string
	Key    string
	Delay  int
	Export bool
}

// LoadEnv loads the Host and Port From .env file.
// Host can be set via NATS_HOST
// Port can be set via NATS_PORT
func NewIDjangoHandler() *IDjangoHandler {
	return &IDjangoHandler{
		URL:    Settings.String("IDJANGO_URL", "w", "https://idjangostage.vivaconagua.org"),
		Key:    Settings.String("IDJANGO_KEY", "w", ""),
		Delay:  Settings.Int("IDJANGO_DELAY", "w", 15),
		Export: Settings.Bool("IDJANGO_EXPORT", "w", false),
	}
}

func (i *IDjangoHandler) Post(data interface{}, path string, sleep ...bool) (err error) {
	if sleep != nil {
		if sleep[0] {
			time.Sleep(time.Duration(i.Delay) * time.Second)
		}
	}
	if i.Export {
		var jsonData []byte
		if jsonData, err = json.Marshal(data); err != nil {
			return
		}
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
			if bodyBytes, err = io.ReadAll(response.Body); err != nil {
				return NewIDjangoError(err, response.StatusCode, nil)
			}
			var body interface{}
			if err = json.Unmarshal(bodyBytes, &body); err != nil {
				return NewIDjangoError(err, 500, string(bodyBytes))
			}
			return NewIDjangoError(nil, response.StatusCode, body)
		}

	}
	return
}

func (i *IDjangoHandler) Put(data interface{}, path string, sleep ...bool) (err error) {
	if sleep != nil {
		if sleep[0] {
			time.Sleep(time.Duration(i.Delay) * time.Second)
		}
	}
	if i.Export {
		var jsonData []byte
		if jsonData, err = json.Marshal(data); err != nil {
			return
		}
		request := new(http.Request)
		request, err = http.NewRequest("PUT", i.URL+path, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		request.Header.Set("Authorization", "Api-Key "+i.Key)
		client := &http.Client{}
		response := new(http.Response)
		response, err = client.Do(request)
		if err != nil {
			return NewIDjangoError(err, 500, nil)
		}
		defer response.Body.Close()
		if response.StatusCode != 200 {
			var bodyBytes []byte
			if bodyBytes, err = io.ReadAll(response.Body); err != nil {
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
	if message == "" && body != nil {
		if bodyMap, ok := body.(map[string]interface{}); ok {
			if errMsg, exists := bodyMap["error_message"].(string); exists && errMsg != "" {
				message = errMsg
			}
		}
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
	return i.Message
}

func (i *IDjangoError) Log() string {
	res, _ := json.Marshal(i)
	return string(res)
}
