package verr

import (
	"net/http"

	"github.com/Viva-con-Agua/vcago/vmod"
)

type (
	ResponseError struct {
		Code     int
		Response interface{}
	}
)

var (
	RespErrorInternalServer = ResponseError{
		Code:     http.StatusInternalServerError,
		Response: vmod.RespInternalServerError,
	}
)
