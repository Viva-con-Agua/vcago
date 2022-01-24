package vcago

import (
	"encoding/json"
	"runtime"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//MongoError represents the struct of an mongo Error type.
type MongoError struct {
	Err        error       `json:"-" bson:"-"`
	Message    string      `json:"message" bson:"message"`
	Filter     interface{} `json:"filter" bson:"filter"`
	Value      interface{} `json:"value" bson:"value"`
	Database   string      `json:"database" bson:"database"`
	Collection string      `json:"collection" bson:"collection"`
	Line       int         `json:"line" bson:"line"`
	File       string      `json:"file" bson:"file"`
}

//NewMongoError creates an mongo Error for an given parameter set
func NewMongoError(err error, value interface{}, filter bson.M, database string, collection string) *MongoError {
	fOutput, _ := json.Marshal(filter)
	temp := new(interface{})
	json.Unmarshal(fOutput, temp)
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	_, line := f.FileLine(pc[0])
	file := runtime.FuncForPC(pc[0]).Name()
	return &MongoError{
		Err:        err,
		Message:    err.Error(),
		Filter:     temp,
		Value:      value,
		Database:   database,
		Collection: collection,
		Line:       line,
		File:       file,
	}
}

//Error return string of the error
func (i *MongoError) Error() string {
	res, _ := json.Marshal(i)
	return string(res)
}

//Response return the ErrorResponse for handling in httpErrorHandler
func (i *MongoError) Response() (int, interface{}) {
	if strings.Contains(i.Message, "duplicate key error") {
		temp := strings.Split(i.Message, "key: {")
		temp = strings.Split(temp[1], "}")
		return Conflict("duplicate key error", "key: {"+temp[0]+"}")
	}
	switch i.Err {
	case mongo.ErrNoDocuments:
		return NotFound("document not found", i.Filter)
	case ErrMongoUpdate:
		return NotFound("document not updated", i.Filter)
	case ErrMongoDelete:
		return NotFound("document not deleted", i.Filter)
	default:
		return InternalServerError()
	}
}
