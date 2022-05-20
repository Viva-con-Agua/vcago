package vmdb

import (
	"net/http"
	"strings"

	"github.com/Viva-con-Agua/vcago"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
type Error struct {
	Err        error  `json:"-" bson:"-"`
	Collection string `json:"collection" bson:"collection"`
	Line       int    `json:"line" bson:"line"`
	File       string `json:"file" bson:"file"`
}

//NewMongoError creates an mongo Error for an given parameter set
func NewError(err error, value interface{}, filter interface{}, database string, collection string) *Error {
	fOutput, _ := json.Marshal(filter)
	temp := new(interface{})
	json.Unmarshal(fOutput, temp)
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	_, line := f.FileLine(pc[0])
	file := runtime.FuncForPC(pc[0]).Name()
	return &Error{
		Type:       "mongo",
		Err:        err,
		Message:    err.Error(),
		Filter:     temp,
		Value:      value,
		Database:   database,
		Collection: collection,
		Line:       line,
		File:       file,
	}
}*/

func Log(c vcago.Context, err error) func(i ...interface{}) {
	if err == mongo.ErrNoDocuments {
		return c.Logger().Debug
	}
	return c.Logger().Error
}

func getCollectionFromDupKey(err error) string {
	temp := strings.Split(err.Error(), "collection: ")
	temp = strings.Split(temp[1], " key:")
	return temp[0]
}

func getKeyFromDupKey(err error) string {
	temp := strings.Split(err.Error(), "key: {")
	temp = strings.Split(temp[1], "}")
	return temp[0]
}

func Response(err error) (int, interface{}) {
	if mongo.IsDuplicateKeyError(err) {
		return vcago.NewResp(
			http.StatusConflict,
			"error",
			"duplicate key error",
			getCollectionFromDupKey(err),
			getKeyFromDupKey(err),
		).Response()
	}
	return vcago.NewInternalServerError("").Response()

	/*switch err.Error() {
	case mongo.ErrNoDocuments:
		return vcago.NewResp(http.StatusNotFound, "error", "document not found", nil, nil).Response()
	case ErrMongoUpdate:
		return NewResp(http.StatusNotFound, "error", "document not updated", nil, nil).Response()
	case ErrMongoDelete:
		return NewResp(http.StatusNotFound, "error", "document not deleted", nil, nil).Response()
	default:
		return NewInternalServerError(i.Collection).Response()

	}*/
}
