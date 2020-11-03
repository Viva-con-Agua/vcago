package vmod

type (
	ResponseMessage struct {
		Message string                 `json:"message"`
		Data    map[string]interface{} `json:"data,omitempty"`
	}
)

var (
	RespInternalServerError = ResponseMessage{Message: "internal_server_error"}
)
