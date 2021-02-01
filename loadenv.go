package vcago

import (
	"log"
	"os"
	"strconv"
	"strings"
)
const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
	notSet = "is not set in the .env file."
)



func envLogError(key string, e string, lvl string, dVal interface{}) bool{
	if lvl == "n" {
		return true
	} 
	if lvl == "w" {
		log.Print(string(colorYellow), "Warning: ", string(colorWhite), key, " ", e, " Default value: ", dVal)
		return true
	}
	if lvl == "e" {
		log.Print(string(colorRed), "Error: ", string(colorWhite), key, " ", e , ". Required for run service.")
		return false
	}
	log.Print(string(colorRed), "Error: ", string(colorWhite), "wrong lvl type. Please set n,w,e.")
	return false
}
//GetEnvString loads a key from enviroment variables as string. 
//The lvl param defines the log level. For warnings set "w" and for error set "e". 
//If the variable is not used or can be ignored use n for do nothing.
//The default value can be set by the dVal param.
func GetEnvString(key string, lvl string, dVal string)(string, bool) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return dVal, envLogError(key, notSet, lvl, dVal)
	}
	return val, true
}

//GetEnvInt loads a key from enviroment variables as int.
//The lvl param defines the log level.
//For warnings set "w" and for error set "e". 
//If the variable is not used or can be ignored use n for do nothing. 
//The default value can be set by the dVal param.
func GetEnvInt(key string, lvl string, dVal int)(int, bool) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return dVal, envLogError(key, notSet, lvl, dVal)
	}
    valInt, err := strconv.Atoi(val)
	if err != nil {
        return dVal, envLogError(key, "is not an integer.", lvl, dVal)
	}
    return valInt, true

}
//GetEnvStringList as
func GetEnvStringList(key string, lvl string, dVal []string)([]string, bool) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return dVal, envLogError(key, notSet, lvl, dVal)
	}
    valList := strings.Split(val, ",")
	if valList != nil {
        return dVal, envLogError(key, "is an empty list.", lvl, dVal)
	}
    return valList, true

}
