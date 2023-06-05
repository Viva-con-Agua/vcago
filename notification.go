package vcago

import (
	"github.com/Viva-con-Agua/vcago/vmod"
)

type (
	NotificationData struct {
		To           string       `json:"to" bson:"to"`
		Service      string       `json:"service" bson:"service"`
		Scope        string       `json:"scope" bson:"scope"`
		Lang         string       `json:"lang" bson:"lang"`
		Content      vmod.Content `json:"content" bson:"content"`
		User         vmod.User    `json:"user" bson:"user"`
		OriginUserID string       `json:"origin_user_id" bson:"origin_user_id"`
	}

	NotificationResponse struct {
		To           string       `json:"to" bson:"to"`
		Service      string       `json:"service" bson:"service"`
		Scope        string       `json:"scope" bson:"scope"`
		Lang         string       `json:"lang" bson:"lang"`
		Content      vmod.Content `json:"content" bson:"content"`
		User         vmod.User    `json:"user" bson:"user"`
		OriginUserID string       `json:"origin_user_id" bson:"origin_user_id"`
		From         string       `json:"from" bson:"from"`
		Subject      string       `json:"subject" bson:"subject"`
		Body         string       `json:"body" bson:"body"`
	}
)

func NewMNotificationData(to string, service string, scope string, lang string, user_id string) *NotificationData {
	return &NotificationData{
		To:           to,
		Service:      service,
		Scope:        scope,
		Lang:         lang,
		OriginUserID: user_id,
	}
}

func (i *NotificationData) Response() *NotificationResponse {
	return &NotificationResponse{
		To:           i.To,
		Service:      i.Service,
		Scope:        i.Scope,
		Lang:         i.Lang,
		Content:      i.Content,
		User:         i.User,
		OriginUserID: i.OriginUserID,
	}
}

func (i *NotificationData) AddUser(user *vmod.User) {
	i.User = *user
}

func (i *NotificationData) AddContent(content *vmod.Content) {
	i.Content = *content
}
