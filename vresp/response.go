package vcago

import "net/http"

//Response represents the default api response struct
// Message shows action information
// Payload contains the response model
type Response struct {
	Message string      `json:"message"`
	Model   string      `json:"model,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

//NewResponse can be used for create Response struct
func NewResponse(model string, payload interface{}) *Response {
	return &Response{
		Model:   model,
		Payload: payload,
	}
}

func (i *Response) Created() (int, interface{}) {
	i.Message = "successfully created"
	return http.StatusCreated, i
}

func (i *Response) Updated() (int, interface{}) {
	i.Message = "successfully updated"
	return http.StatusOK, i
}
func (i *Response) Executed() (int, interface{}) {
	i.Message = "successfully executed"
	return http.StatusOK, i
}
func (i *Response) Selected() (int, interface{}) {
	i.Message = "successfully selected"
	return http.StatusOK, i
}
func (i *Response) Deleted() (int, interface{}) {
	i.Message = "successfully deleted"
	return http.StatusOK, i
}
