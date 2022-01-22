package vlog

type ErrorResponse struct {
	Status  int         `json:"-"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
	Coll    string      `json:"collection,omitempty"`
	Model   string      `json:"model,omitempty"`
}
