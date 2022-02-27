package vcago

import "encoding/json"

const StatusDONE = "done"
const StatusPROCESS = "process"
const StatusFAILED = "failed"
const StatusCREATED = "created"
const StatusOPEN = "open"
const StatusINTERNAL = "internal_error"
const StatusBADREQUEST = "bad_request"

type Status struct {
	StatusType    string `bson:"status_type" json:"status_type"`
	StatusMessage string `bson:"status_message" json:"status_message"`
}

func NewStatus() *Status {
	return &Status{}
}

func NewStatusInternal(message error) *Status {
	return &Status{
		StatusType:    StatusINTERNAL,
		StatusMessage: message.Error(),
	}
}

func NewStatusBadRequest(message error) *Status {
	return &Status{
		StatusType:    StatusBADREQUEST,
		StatusMessage: message.Error(),
	}
}

func (i *Status) Set(t string, message string) *Status {
	i.StatusType = t
	i.StatusMessage = message
	return i
}

func (i *Status) Done(message string) {
	i.StatusType = StatusDONE
	i.StatusMessage = message
}

func (i *Status) Process(message string) {
	i.StatusType = StatusPROCESS
	i.StatusMessage = message
}
func (i *Status) Failed(message string) {
	i.StatusType = StatusFAILED
	i.StatusMessage = message
}
func (i *Status) Created(message string) {
	i.StatusType = StatusCREATED
	i.StatusMessage = message
}
func (i *Status) Open(message string) {
	i.StatusType = StatusOPEN
	i.StatusMessage = message
}

func (i *Status) Internal(message string) {
	i.StatusType = StatusINTERNAL
	i.StatusMessage = message
}

func (i *Status) ValidateDone() error {
	if i.StatusType == StatusDONE {
		return i
	}
	return nil
}

func (i *Status) Error() string {
	res, _ := json.Marshal(i)
	return string(res)
}

func (i *Status) Response() (int, interface{}) {
	switch i.StatusType {
	case StatusINTERNAL:
		return InternalServerError()
	case StatusBADREQUEST:
		return BadRequest("status", i)
	default:
		return BadRequest("status", i)
	}
}
