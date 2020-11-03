package verr

import "net/http"

type (
	ResponseError struct {
		Code     int
		Response interface{}
	}

	ResponseMessage struct {
		Message string                 `json:"message"`
		Data    map[string]interface{} `json:"data,omitempty"`
	}
)

var (
	RespErrorInternalServer = ResponseError{
		Code:     http.StatusInternalServerError,
		Response: RespInternalServerError,
	}
	RespInternalServerError = ResponseMessage{Message: "internal_server_error"}
)
