package vcago

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type MailSend struct {
	URL  string
	Key  string
	Mode string
}

func NewMailSend() *MailSend {
	return &MailSend{
		URL:  Settings.String("MAIL_URL", "w", "http://factory-api.localhost"),
		Key:  Settings.String("MAIL_API_KEY", "w", "secret"),
		Mode: Settings.String("MAIL_MODE", "w", "local"),
	}
}

func (i *MailSend) Send(mail *MailData) {
	if i.Mode == "local" {
		i.Print(mail)
	} else if i.Mode == "nats" {
		i.Nats(mail)
	} else if i.Mode == "post" {
		go i.Post(mail)
	}
}

func (i *MailSend) Nats(mail *MailData) {
	Nats.Publish("mail.send", mail)
}

func (i *MailSend) Print(mail *MailData) {
	output, _ := json.MarshalIndent(mail, "", "\t")
	log.Print(string(output))
}

func (i *MailSend) Subscribe() {
	Nats.Subscribe("mail.send", func(m *MailData) { i.Post(m) })
}

func (i *MailSend) Post(mailData *MailData) {
	var jsonData []byte
	var err error
	if jsonData, err = json.Marshal(mailData); err != nil {
		log.Print(err)
		return
	}
	request := new(http.Request)
	if request, err = http.NewRequest("POST", i.URL+"/mails/send", bytes.NewBuffer(jsonData)); err != nil {
		log.Print(err)
		return
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", "Bearer "+i.Key)
	client := &http.Client{}
	response := new(http.Response)
	response, err = client.Do(request)
	if err != nil {
		log.Print(NewIDjangoError(err, response.StatusCode, nil))
		return
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		var bodyBytes []byte
		if bodyBytes, err = ioutil.ReadAll(response.Body); err != nil {
			log.Print(NewIDjangoError(err, response.StatusCode, nil))
			return
		}
		body := new(interface{})
		if err = json.Unmarshal(bodyBytes, body); err != nil {
			log.Print(NewIDjangoError(err, 500, string(bodyBytes)))
			return
		}
		log.Print(NewIDjangoError(nil, response.StatusCode, body))
		return
	}
}

func (i *MailSend) PostCycularMail(data *CycularMail) {
	var jsonData []byte
	var err error
	if jsonData, err = json.Marshal(data); err != nil {
		log.Print(err)
		return
	}
	request := new(http.Request)
	if request, err = http.NewRequest("POST", i.URL+"/mails/send/cycle", bytes.NewBuffer(jsonData)); err != nil {
		log.Print(err)
		return
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", "Bearer "+i.Key)
	client := &http.Client{}
	response := new(http.Response)
	response, err = client.Do(request)
	if err != nil {
		log.Print(NewIDjangoError(err, response.StatusCode, nil))
		return
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		var bodyBytes []byte
		if bodyBytes, err = ioutil.ReadAll(response.Body); err != nil {
			log.Print(NewIDjangoError(err, response.StatusCode, nil))
			return
		}
		body := new(interface{})
		if err = json.Unmarshal(bodyBytes, body); err != nil {
			log.Print(NewIDjangoError(err, 500, string(bodyBytes)))
			return
		}
		log.Print(NewIDjangoError(nil, response.StatusCode, body))
		return
	}
}
