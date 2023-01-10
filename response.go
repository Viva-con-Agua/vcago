package vcago

import (
	"encoding/json"
	"net/http"
)

// Response represents the default api response struct
// Status defines the response status code
// Type defines the response type. Can be success or error
// Message shows action information
// Model shows the collection that would be attached
// Payload contains the response model
type Response struct {
	Status  int         `json:"-"`
	Type    string      `json:"type" bson:"type"`
	Message string      `json:"message" bson:"message"`
	Model   string      `json:"model,omitempty" bson:"model,omitempty"`
	Payload interface{} `json:"payload,omitempty" bson:"payload,omitempty"`
}

// IDParam data struct for handling '/:id'.
// ID needs to be a uuid.
type IDParam struct {
	ID string `param:"id" validate:"uuid"`
}

// DeletedResponse used for handling response in deleted case.
type DeletedResponse struct {
	ID string `json:"id"`
}

// Response returns an tuple that can be used with echo.Context.JSON.
func (i *Response) Response() (int, *Response) {
	return i.Status, i
}

// NewResp creates new Response model.
func NewResp(status int, typ string, message string, model string, payload interface{}) *Response {
	return &Response{
		Status:  status,
		Type:    typ,
		Message: message,
		Model:   model,
		Payload: payload,
	}

}

// Error implements an error interface for handling Responses as Error.
// The function returns an json.Marshal of the error as string.
func (i *Response) Error() string {
	res, _ := json.Marshal(i)
	return string(res)
}

// NewCreated returns a Response model intended for a POST request that creates a model.
//
// Status: 201 Created
//
// JSON:
//
//	{
//		"type": "success",
//		"message": "successfully_created",
//		"model": model,
//		"payload": payload
//	}
func NewCreated(model string, payload interface{}) *Response {
	return NewResp(http.StatusCreated, "success", "successfully_created", model, payload)
}

// NewUpdated returns a Response model intended for a PUT request that updates a model.
//
// Status: 200 OK
//
// JSON:
//
//	{
//		"type": "success",
//		"message": "successfully_updated",
//		"model": model,
//		"payload": payload
//	}
func NewUpdated(model string, payload interface{}) *Response {
	return NewResp(http.StatusOK, "success", "successfully_updated", model, payload)
}

// NewDeleted returns a Response model intended for a DELETE request that deletes a model.
//
// Status: 200 OK
//
// JSON:
//
//	{
//		"type": "success",
//		"message": "successfully_deleted",
//		"model": model,
//		"payload": payload
//	}
func NewDeleted(model string, payload interface{}) *Response {
	return NewResp(http.StatusOK, "success", "successfully_deleted", model, payload)
}

// NewSelected returns a Response model intended for a GET request that selects a model or list.
//
// Status: 200 OK
//
// JSON:
//
//	{
//		"type": "success",
//		"message": "successfully_selected",
//		"model": model,
//		"payload": payload
//	}
func NewSelected(model string, payload interface{}) *Response {
	return NewResp(http.StatusOK, "success", "successfully_selected", model, payload)
}

// NewExecuted returns an Response model intended for a request that execute an process.
//
// Status: 200 OK
//
// JSON:
//
//	{
//		"type": "success",
//		"message": "successfully_executed",
//		"model": model,
//		"payload": payload
//	}
func NewExecuted(model string, payload interface{}) *Response {
	return NewResp(http.StatusOK, "success", "successfully_executed", model, payload)
}

// NewBadRequest returns an Response model intended for an bad request response.
//
// Status: 400 Bad Request
//
// JSON with payload:
//
//	{
//		"type": "error",
//		"message": message,
//		"model": model,
//		"payload": payload
//	}
//
// JSON without payload:
//
//	{
//		"type": "error",
//		"message": message,
//		"model": model,
//	}
func NewBadRequest(model string, message string, payload ...interface{}) *Response {
	return NewResp(http.StatusBadRequest, "error", message, model, payload)
}

// NewInternalServerError returns an Response model intended for an internal server error response.
// The payload param is optional.
//
// Status: 500 Internal Server Error
//
// JSON with payload:
//
//	{
//		"type": "error",
//		"message": "internal_server_error",
//		"model": model,
//		"payload": payload
//	}
//
// JSON without payload:
//
//	{
//		"type": "error",
//		"message": "internal_server_error",
//		"model": model
//	}
func NewInternalServerError(model string, payload ...interface{}) *Response {
	return NewResp(http.StatusInternalServerError, "error", "internal_server_error", model, payload)
}

// NewConflict returns an Response model intended for an conflict error response.
//
// Status: 409 Conflict
//
// JSON with payload:
//
//	{
//		"type": "error",
//		"message": "conflict",
//		"model": model,
//		"payload": payload
//	}
//
// JSON without payload:
//
//	{
//		"type": "error",
//		"message": "conflict",
//		"model": model
//	}
func NewConflict(model string, payload ...interface{}) *Response {
	return NewResp(http.StatusConflict, "error", "conflict", model, payload)
}

// NewNotFound returns an Response model intended for an not found error response.
//
// Status: 404 Not Found
//
// JSON with payload:
//
//	{
//		"type": "error",
//		"message": "not_found",
//		"model": model,
//		"payload": payload
//	}
//
// JSON without payload:
//
//	{
//		"type": "error",
//		"message": "not_found",
//		"model": model
//	}
func NewNotFound(model string, payload ...interface{}) *Response {
	return NewResp(http.StatusNotFound, "error", "not_found", model, payload)
}

// NewPermissionDenied returns an Response model intended for an permission denied error response.
//
// Status: 400 Bad Request
//
// JSON with payload:
//
//	{
//		"type": "error",
//		"message": "permission_denied",
//		"model": model,
//		"payload": payload
//	}
//
// JSON without payload:
//
//	{
//		"type": "error",
//		"message": "permission_denied",
//		"model": model
//	}
func NewPermissionDenied(model string, payload ...interface{}) *Response {
	return NewResp(http.StatusBadRequest, "error", "permission_denied", model, payload)
}
