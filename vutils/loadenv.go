package vutils

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

//HTTPBaseCookie defines the base cookie setup for vcago
var HTTPBaseCookie http.Cookie

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
	notSet      = "is not set in the .env file."
)

//LoadEnv used for loading environment variables.
type LoadEnv []bool

var Config = LoadConfig()

//NatsHost is the ip of the nats service.
var NatsHost string

//NatsPort is the port ot the nats service.
var NatsPort string

//Env load the environment variables for vcago

func LoadConfig() *LoadEnv {
	godotenv.Load(".env")
	return new(LoadEnv)
}

func envLogError(key string, e string, lvl string, dVal interface{}) bool {
	if lvl == "n" {
		return true
	}
	if lvl == "w" {
		log.Print(string(colorYellow), "Warning: ", string(colorWhite), key, " ", e, " Default value: ", dVal)
		return true
	}
	if lvl == "e" {
		log.Print(string(colorRed), "Error: ", string(colorWhite), key, " ", e, ". Required for run service.")
		return false
	}
	log.Print(string(colorRed), "Error: ", string(colorWhite), "wrong lvl type. Please set n,w,e.")
	return false
}

//GetEnvString loads a key from enviroment variables as string.
//The lvl param defines the log level. For warnings set "w" and for error set "e".
//If the variable is not used or can be ignored use n for do nothing.
//The default value can be set by the dVal param.
func (l LoadEnv) GetEnvString(key string, lvl string, dVal string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		l = append(l, envLogError(key, notSet, lvl, dVal))
		return dVal
	}
	l = append(l, true)
	return val
}

//GetEnvInt loads a key from enviroment variables as int.
//The lvl param defines the log level.
//For warnings set "w" and for error set "e".
//If the variable is not used or can be ignored use n for do nothing.
//The default value can be set by the dVal param.
func (l LoadEnv) GetEnvInt(key string, lvl string, dVal int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		l = append(l, envLogError(key, notSet, lvl, dVal))
		return dVal
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		l = append(l, envLogError(key, notSet, lvl, dVal))
		return dVal
	}
	l = append(l, true)
	return valInt

}

//GetEnvStringList as
func (l LoadEnv) GetEnvStringList(key string, lvl string, dVal []string) []string {
	val, ok := os.LookupEnv(key)
	if !ok {
		l = append(l, envLogError(key, notSet, lvl, dVal))
		return dVal
	}
	valList := strings.Split(val, ",")
	if valList == nil {
		l = append(l, envLogError(key, notSet, lvl, dVal))
		return dVal

	}
	l = append(l, true)

	return valList
}

//GetEnvBool load a key from environment variables as bool.
func (l LoadEnv) GetEnvBool(key string, lvl string, dVal bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		l = append(l, envLogError(key, notSet, lvl, dVal))
		return dVal
	}
	if val == "true" {
		l = append(l, true)
		return true
	}
	if val == "false" {
		l = append(l, true)
		return false
	}
	l = append(l, envLogError(key, notSet, lvl, dVal))
	return dVal
}

//Validate check if LoadEnv is valid and log.Fatal if on entry is false.
func (l LoadEnv) Validate() {
	for i := range l {
		if !l[i] {
			log.Fatal("Please set enviroment variables in the .env file. Read logs above.")
		}
	}
}
