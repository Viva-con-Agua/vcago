package vmod

type (
	//MailBase base model for request mail-backend via nats
	MailBase struct {
		To           string      `json:"to" bson:"to"`
		Case string `json:"case"`
		Scope string `json:"scope"`
	}
	//MailCode used for request mail-backend via nats
	MailCode struct {
		Name string `json:"name"`
		Code string `json:"code"`
		*MailBase
	}
)

