package vmdb

import (
	"encoding/json"
	"runtime"

	"go.mongodb.org/mongo-driver/bson"
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
