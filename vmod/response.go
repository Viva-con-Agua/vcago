package vmod

import "net/http"

type (
	//APIResponse type that will response in success cases
	APIResponse struct {
		Message string      `json:"message"`
		Model   string      `json:"model"`
		Payload interface{} `json:"payload,omitempty"`
	}
)

//RespCreated creates the created response
func RespCreated(i interface{}, model string) (int, interface{}) {
	return http.StatusCreated, APIResponse{Message: "successful_created", Model: model, Payload: i}
}

//RespUpdated creates the updated response
func RespUpdated(model string) (int, interface{}) {
	return http.StatusOK, APIResponse{Message: "successful_updated", Model: model}
}

//RespDeleted creates the updated response
func RespDeleted(model string) (int, interface{}) {
	return http.StatusOK, APIResponse{Message: "successful_deleted", Model: model}
}

//RespSelected creates the selected response
func RespSelected(i interface{}, model string) (int, interface{}) {
	return http.StatusOK, APIResponse{Message: "successful_selected", Model: model, Payload: i}
}

//RespCreateBase creates the created response without payload
func RespCreateBase(model string) (int, interface{}) {
	return http.StatusOK, APIResponse{Message: "successful_created", Model: model}
}

//RespExecuted creates the execution response for handling for example sending emails.
func RespExecuted(model string) (int, interface{}) {
	return http.StatusOK, APIResponse{Message: "successful_executed", Model: model}
}
