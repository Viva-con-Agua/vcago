package vdjango

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Viva-con-Agua/vcago/vutils"
)

type IDjango struct {
	URL    string
	Key    string
	Export bool
}

func NewIDjango() (r *IDjango) {
	r = &IDjango{
		URL:    vutils.Config.GetEnvString("IDJANGO_URL", "w", "https://idjangostage.vivaconagua.org"),
		Key:    vutils.Config.GetEnvString("IDJANGO_KEY", "w", ""),
		Export: vutils.Config.GetEnvBool("IDJANGO_EXPORT", "w", false),
	}
	return
}

func (i *IDjango) Post(data interface{}, path string) (err error) {
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
			if bodyBytes, err = ioutil.ReadAll(response.Body); err != nil {
				return NewIDjangoError(err, response.StatusCode, nil)
			}
			body := new(interface{})
			if err = json.Unmarshal(bodyBytes, body); err != nil {
				return NewIDjangoError(err, 500, string(bodyBytes))
			}
			return NewIDjangoError(nil, response.StatusCode, body)
		}

	} else {
		var val []byte
		if val, err = json.MarshalIndent(data, "", "    "); err != nil {
			return
		}
		log.Print(string(val))
	}
	return
}
