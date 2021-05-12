package verr

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

//GetEnvWithLog Validate .env Variables
func GetEnvWithLog(key string) (string, bool) {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("Warning: %s not set\n", key)
		return "", ok
	}
	return val, ok
}

//GetIntEnvWithLog Validate .env Variables
func GetIntEnvWithLog(key string) (int, bool) {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("Warning: %s not set\n", key)
		return 0, false
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("Error: %s is not an Integer \n", key)
		return 0, false
	}
	return valInt, true
}

func InsertErrorResponse(c echo.Context, err error) error {
	if strings.Contains(err.Error(), "duplicate key error") {
		LogError(c.Request().Context(), err, "debug")
		return echo.NewHTTPError(http.StatusConflict, handleDuplicateKey(err))
	}
	LogError(c.Request().Context(), err, "prod")
	return InternalServerErrorMsg
}

func ListErrorResponse(c echo.Context, err error) error {
	if err == mongo.ErrNoDocuments {
		LogError(c.Request().Context(), err, "debug")
		return echo.NewHTTPError(http.StatusNotFound, MongoBase{Message: "not_found"})
	}
	LogError(c.Request().Context(), err, "prod")
	return InternalServerErrorMsg
}
