package vcago

import (
	"encoding/json"
	"net/http"
)

//Response represents the default api response struct
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

//Response returns an tuple for handling in echo.Context.JSON
func (i *Response) Response() (int, *Response) {
	return i.Status, i
}

//NewResp creates new Response model.
func NewResp(status int, typ string, message string, model string, payload interface{}) *Response {
	return &Response{
		Status:  status,
		Type:    typ,
		Message: message,
		Model:   model,
		Payload: payload,
	}

}

func (i *Response) Error() string {
	res, _ := json.Marshal(i)
	return string(res)
}

//NewCreated returns a Response model intended for a POST request that creates a model.
func NewCreated(model string, payload interface{}) *Response {
	return NewResp(http.StatusCreated, "success", "successfully_created", model, payload)
}

//NewUpdated returns a Response model intended for a PUT request that updates a model.
func NewUpdated(model string, payload interface{}) *Response {
	return NewResp(http.StatusOK, "success", "successfully_updated", model, payload)
}

//NewDeleted returns a Response model intended for a DELETE request that deletes a model.
func NewDeleted(model string, payload interface{}) *Response {
	return NewResp(http.StatusOK, "success", "successfully_deleted", model, payload)
}

//NewSelected returns a Response model intended for a GET request that selects a model or list.
func NewSelected(model string, payload interface{}) *Response {
	return NewResp(http.StatusOK, "success", "successfully_selected", model, payload)
}

func NewExecuted(model string, payload interface{}) *Response {
	return NewResp(http.StatusOK, "success", "successfully_executed", model, payload)
}
func NewBadRequest(model string, message string, payload ...interface{}) *Response {
	return NewResp(http.StatusBadRequest, "error", message, model, payload)
}
func NewInternalServerError(model string, payload ...interface{}) *Response {
	return NewResp(http.StatusInternalServerError, "error", "internal_server_error", model, payload)
}
func NewConflict(model string, payload ...interface{}) *Response {
	return NewResp(http.StatusConflict, "error", "conflict", model, payload)
}

func NewNotFound(model string, payload ...interface{}) *Response {
	return NewResp(http.StatusNotFound, "error", "not_found", model, payload)
}

func NewPermissionDenied(model string, payload ...interface{}) *Response {
	return NewResp(http.StatusBadRequest, "error", "permission_denied", model, payload)
}

/*
//NewResponse can be used for create Response struct
func NewResponse(model string, payload interface{}) *Response {
	return &Response{
		Model:   model,
		Payload: payload,
	}
}

func (i *Response) Created() (int, *Response) {
	i.Message = "successfully created"
	return http.StatusCreated, i
}

func (i *Response) Updated() (int, *Response) {
	i.Message = "successfully updated"
	return http.StatusOK, i
}
func (i *Response) Executed() (int, *Response) {
	i.Message = "successfully executed"
	return http.StatusOK, i
}
func (i *Response) Selected() (int, *Response) {
	i.Message = "successfully selected"
	return http.StatusOK, i
}
func (i *Response) Deleted() (int, *Response) {
	i.Message = "successfully deleted"
	return http.StatusOK, i
}

func (i *Response) InternalServerError() (int, *Response) {
	i.Message = "internal server error"
	return http.StatusInternalServerError, i
}
func (i *Response) BadRequest() (int, *Response) {
	i.Message = "bad request"
	return http.StatusBadRequest, i
}

func (i *Response) Conflict() (int, *Response) {
	i.Message = "conflict"
	return http.StatusConflict, i
}

func (i *Response) NotFound() (int, *Response) {
	i.Message = "not found"
	return http.StatusNotFound, i
}*/
