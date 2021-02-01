package vmod

type (
	//NatsMailCode represents the email backend model
	NatsMailCode struct {
		To    string `json:"to"`
		UserID string `json:"user_id"`
		Name  string `json:"name"`
		Code  string `json:"code"`
		Scope string `json:"scope"`
	}
)
