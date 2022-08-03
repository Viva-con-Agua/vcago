package vcago

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

//Settings represents the global SettingType variable and can be used for load config parameters.
var Settings = SettingHandlerLoad()

//SettingHandler represents and handler for load Settings via flag or environment variable.
type SettingHandler struct {
	Error  []bool
	Config map[string]interface{}
}

//SettingHandlerLoad loads all variables form an .env file and return an SettingHandler.
func SettingHandlerLoad() *SettingHandler {
	godotenv.Load(".env")
	settings := new(SettingHandler)
	settings.Config = make(map[string]interface{})
	return settings
}

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

//envLogError print all warnings and errors.
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

//settingsLogError print all warnings and errors.
func settingsLogError(key string, e string, lvl string, dVal interface{}) bool {
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

//stringEnv used to load an string variable from environment variables.
func (i *SettingHandler) stringEnv(key string, lvl string, dVal string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		i.Error = append(i.Error, envLogError(key, notSet, lvl, dVal))
		return dVal
	}
	i.Error = append(i.Error, true)
	return val
}

//intEnv used to load an int variable from environment variables.
func (i *SettingHandler) intEnv(key string, lvl string, dVal int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		i.Error = append(i.Error, envLogError(key, notSet, lvl, dVal))
		return dVal
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		i.Error = append(i.Error, envLogError(key, notSet, lvl, dVal))
		return dVal
	}
	i.Error = append(i.Error, true)
	return valInt

}

//stringListEnv used to load an int variable from environment variables.
func (i *SettingHandler) stringListEnv(key string, lvl string, dVal []string) []string {
	val, ok := os.LookupEnv(key)
	if !ok {
		i.Error = append(i.Error, envLogError(key, notSet, lvl, dVal))
		return dVal
	}
	valList := strings.Split(val, ",")
	if valList == nil {
		i.Error = append(i.Error, envLogError(key, notSet, lvl, dVal))
		return dVal

	}
	i.Error = append(i.Error, true)

	return valList
}

//boolEnv load a key from environment variables as bool.
func (i *SettingHandler) boolEnv(key string, lvl string, dVal bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		i.Error = append(i.Error, envLogError(key, notSet, lvl, dVal))
		return dVal
	}
	if val == "true" {
		i.Error = append(i.Error, true)
		return true
	}
	if val == "false" {
		i.Error = append(i.Error, true)
		return false
	}
	i.Error = append(i.Error, envLogError(key, notSet, lvl, dVal))
	return dVal
}

//String loads an string config variable.
//The function will first looking for an flag, than the environment variables and as default dVal.
func (i *SettingHandler) String(key string, lvl string, dVal string) string {
	val := flag.String(key, i.stringEnv(key, lvl, dVal), "")
	return *val
}

func (i *SettingHandler) Bool(key string, lvl string, dVal bool) bool {
	val := flag.Bool(key, i.boolEnv(key, lvl, dVal), "")
	return *val
}

func (i *SettingHandler) Int(key string, lvl string, dVal int) int {
	val := flag.Int(key, i.intEnv(key, lvl, dVal), "")
	return *val
}

func (i *SettingHandler) StringList(key string, lvl string, dVal []string) []string {
	ddVal := ""
	for n := range dVal {
		ddVal = ddVal + dVal[n] + ","
	}
	val := flag.String(key, i.stringEnv(key, lvl, ddVal), "")
	valList := strings.Split(*val, ",")
	if valList == nil {
		i.Error = append(i.Error, envLogError(key, notSet, lvl, dVal))
		return dVal

	}
	return valList
}

func (i *SettingHandler) Load() {
	flag.Parse()
}
