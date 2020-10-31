package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	ResponseMessage struct {
		Message string                 `json:"message"`
		Data    map[string]interface{} `json:"data,omitempty"`
	}
)

func RespMessage(m string, k *string, i interface{}) interface{} {
	r := new(ResponseMessage)
	r.Message = m
	if k != nil && i != nil {
		data := make(map[string]interface{})
		data[*k] = i
	}
	return r
}
func RespNoContent(k string, i interface{}) interface{} {
	data := make(map[string]interface{})
	data[k] = i
	response := new(ResponseMessage)
	response.Message = "Not found"
	response.Data = data
	return response
}

func RespConflict(k string, i interface{}) interface{} {
	data := make(map[string]interface{})
	data[k] = i
	response := new(ResponseMessage)
	response.Message = "Model already exists"
	response.Data = data
	return response
}

func RespInternelServerError() interface{} {
	response := new(ResponseMessage)
	response.Message = "Internel server error, please check logs"
	return response
}

func RespCreated() interface{} {
	response := new(ResponseMessage)
	response.Message = "Successful created"
	return response
}

func RespHandlingError(c echo.Context, api_err *ApiError, i interface{}) (err error) {
	if api_err.Error != nil {
		if api_err.Error == ErrorUnauthorized {
			return c.JSON(http.StatusUnauthorized, api_err.Message)
		}
		if api_err.Error == ErrorForbidden {
			return c.JSON(http.StatusForbidden, api_err.Message)
		}
		if api_err.Error == ErrorConflict {
			return c.JSON(http.StatusConflict, api_err.Message)
		}
		api_err.LogError(c, i)
		return c.JSON(http.StatusInternalServerError, api_err.Message)
	}
	return nil
}
