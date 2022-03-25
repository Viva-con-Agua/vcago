package vcago

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//AdminRequest represents model for admin requests.
type AdminRequest struct {
	URL string
}

func NewAdminRequest() *AdminRequest {
	return &AdminRequest{
		URL: Config.GetEnvString("ADMIN_URL", "n", "http://172.4.5.3"),
	}
}

func (i *AdminRequest) Get(path string) (r *Response, err error) {
	url := i.URL + path
	request := new(http.Request)
	request, err = http.NewRequest("GET", url, nil)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response := new(http.Response)
	if response, err = client.Do(request); err != nil {
		return
	}
	defer response.Body.Close()
	var bodyBytes []byte
	if response.StatusCode != 200 {
		if bodyBytes, err = ioutil.ReadAll(response.Body); err != nil {
			return
		}
		body := new(Response)
		if err = json.Unmarshal(bodyBytes, body); err != nil {
			return
		}
		return
	}
	r = new(Response)
	if bodyBytes, err = ioutil.ReadAll(response.Body); err != nil {
		return
	}
	if err = json.Unmarshal(bodyBytes, r); err != nil {
		log.Print(err)
		return
	}
	return
}
