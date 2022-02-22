package vcago

import (
	"encoding/json"
	"log"
)

type MailData struct {
	TO        string    `json:"to" bson:"to"`
	Service   string    `json:"service" bson:"service"`
	Scope     string    `json:"scope" bson:"scope"`
	Lang      string    `json:"lang" bson:"lang"`
	User      User      `json:"user" bson:"user"`
	LinkToken LinkToken `json:"link_token" bson:"link_token"`
}

func NewMailData(to string, service string, scope string, lang string) *MailData {
	return &MailData{
		TO:      to,
		Service: service,
		Scope:   scope,
	}
}

func (i *MailData) Send() (err error) {
	mode := Config.GetEnvString("MAIL_MODE", "w", "local")
	if mode == "local" {
		output, _ := json.MarshalIndent(i, "", "\t")
		log.Print(string(output))
	}
	return
}

func (i *MailData) AddUser(user *User) {
	i.User = *user
}

func (i *MailData) AddLinkToken(token *LinkToken) {
	i.LinkToken = *token
}
