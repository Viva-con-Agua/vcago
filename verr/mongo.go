package verr

import (
	"errors"
	"net/http"
	"runtime"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
    //MongoConflict represents the reponse type for mongo database Conficts.
    MongoConflict struct {
        Message string `json:"message"`
        Coll string `json:"collection"`
        Value string `json:"value"`
    }
    //MongoBase represents response type basic mongo error with collection involved.
    MongoBase struct {
        Message string `json:"message"`
        Coll string `json:"collection,omitempty"`
    }
)

//MongoInsertOneError handles error response and logging for collection.insertOne() function.
func MongoInsertOneError(c echo.Context, err error, coll string) error {
    if strings.Contains(err.Error(), "duplicate key error"){
        LogError(c, err, "debug")
        return echo.NewHTTPError(http.StatusConflict, handleDuplicateKey(err))
    }
    LogError(c, err, "prod")
    return InternalServerErrorMsg
} 
//MongoFindOneError handles error response and logging for collection.findOne() function.
func MongoFindOneError(c echo.Context, err error, coll string) error {
    if err == mongo.ErrNoDocuments {
        LogError(c, err, "debug")
        return echo.NewHTTPError(http.StatusNotFound, MongoBase{Message: "not_found", Coll: coll})
    }
    LogError(c, err, "prod")
    return InternalServerErrorMsg
}


// Kann echt weg jo   | 
//                   \/

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

func logError(c echo.Context, errString string, coll string) {
	//get infos about function
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	function := runtime.FuncForPC(pc[0])

	//fill ApiError
	_, line := function.FileLine(pc[0])
	file := runtime.FuncForPC(pc[0]).Name()   //get infos about function
	log.Print(
			"\n",
			string(colorRed), "Error Message: \n",
			"\t", string(colorWhite), errString, "\n",
			"\tFile: [", file, "]\n",
			"\tLine: [", line, "]\n",
			//string(colorYellow), "Session_User: "+string(colorWhite)+"\n", u, string(user), "\n",
			//string(colorYellow), "Request_Header: ", string(colorWhite), "\n\t",
			//formatRequestPrint(c.Request()), "\n",
		//	string(colorYellow), "Request_Body: ", string(colorWhite), "\n",

			string(colorBlue), "### END ERROR", string(colorWhite), "\n\n",
		)

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


