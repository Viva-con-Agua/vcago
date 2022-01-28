package vcago

import "net/http"

//Response represents the default api response struct
// Message shows action information
// Model shows the collection that would be attached
// Payload contains the response model
type Response struct {
	Message string      `json:"message" bson:"message"`
	Model   string      `json:"model,omitempty" bson:"model,omitempty"`
	Payload interface{} `json:"payload,omitempty" bson:"payload,omitempty"`
}

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
}
