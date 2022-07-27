package vcago

import (
	"encoding/json"
	"log"

	"github.com/Viva-con-Agua/vcago/vmod"
)

type MailData struct {
	TO          string    `json:"to" bson:"to"`
	Service     string    `json:"service" bson:"service"`
	Scope       string    `json:"scope" bson:"scope"`
	Lang        string    `json:"lang" bson:"lang"`
	User        vmod.User `json:"user" bson:"user"`
	LinkToken   LinkToken `json:"link_token" bson:"link_token"`
	CurrentUser MailUser  `json:"current_user" bson:"current_user"`
	ContactUser MailUser  `json:"contact_user" bson:"contact_user"`
}

type MailUser struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewMailData(to string, service string, scope string, lang string) *MailData {
	return &MailData{
		TO:      to,
		Service: service,
		Scope:   scope,
	}
}

func (i *MailData) Send() (err error) {
	mode := Settings.String("MAIL_MODE", "w", "local")
	if mode == "local" {
		output, _ := json.MarshalIndent(i, "", "\t")
		log.Print(string(output))
	}
	return
}

func (i *MailData) AddUser(user *vmod.User) {
	i.User = *user
}

func (i *MailData) AddLinkToken(token *LinkToken) {
	i.LinkToken = *token
}

func (i *MailData) AddCurrentUser(id string, email string, firstName string, lastName string) {
	i.CurrentUser = MailUser{
		ID:        id,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
}

func (i *MailData) AddContactUser(id string, email string, firstName string, lastName string) {
	i.ContactUser = MailUser{
		ID:        id,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
}
