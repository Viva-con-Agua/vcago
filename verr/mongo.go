package verr

import (
	"errors"
	"net/http"
	"runtime"
	"strings"


	"go.mongodb.org/mongo-driver/mongo"
)

type (
    //Response type for mongo database Conficts
    MongoConflict struct {
        Message string `json:"message"`
        Coll string `json:"collection"`
        Value string `json:"value"`
    }
    MongoBase struct {
        Message string `json:"message"`
        Coll string `json:"collection"`
    }
)

func NewMongoCollError(err error, coll string) error {
    return errors.New(err.Error() + " collection: " + coll)
}

func handleDuplicateKey(err error) *MongoConflict {
		cut1 := strings.Split(err.Error(), "collection: ")
		cut2 := strings.Split(cut1[1], " index: ") 
		cut3 := strings.Split(cut2[1], "dup")
		value := strings.Split(cut3[0], "_")

		coll := strings.Split(cut2[0], ".")
        return &MongoConflict{
            Message: "duplicate key error",
            Coll: coll[1],
            Value: value[0],
        }
}

func handleNoDocument(err error) *MongoBase {
    cut := strings.Split(err.Error(), "collection:")
    return &MongoBase{
        Message: cut[0],
        Coll: cut[1],
    }
}

func MongoHandleError(err error) *APIError{
	//get infos about function
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
    _, line := f.FileLine(pc[0])
    apiErr := &APIError{
        Error: err,
        Line: line,
        File: runtime.FuncForPC(pc[0]).Name(),
    }
    if strings.Contains(err.Error(), "duplicate key error"){
        apiErr.Code = http.StatusConflict
        apiErr.Body = handleDuplicateKey(err)
        apiErr.Level = false
        return apiErr
    } else if strings.Contains(err.Error(), mongo.ErrNoDocuments.Error()) {
        apiErr.Code = http.StatusNotFound
        apiErr.Body = handleNoDocument(err)
        apiErr.Level = false
        return apiErr
    }
    return apiErr.InternalServerError()
}


