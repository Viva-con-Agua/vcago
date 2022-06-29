package vcago

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var Settings = SettingTypeLoad()

type SettingType struct {
	Error []bool
}

func SettingTypeLoad() *SettingType {
	godotenv.Load(".env")
	return new(SettingType)
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

func (i *SettingType) StringEnv(key string, lvl string, dVal string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		i.Error = append(i.Error, envLogError(key, notSet, lvl, dVal))
		return dVal
	}
	i.Error = append(i.Error, true)
	return val
}

func (i *SettingType) IntEnv(key string, lvl string, dVal int) int {
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

//GetEnvStringList as
func (i *SettingType) StringListEnv(key string, lvl string, dVal []string) []string {
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

//GetEnvBool load a key from environment variables as bool.
func (i *SettingType) BoolEnv(key string, lvl string, dVal bool) bool {
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

func (i *SettingType) String(key string, lvl string, dVal string) string {
	val := flag.String(key, i.StringEnv(key, lvl, dVal), "")
	flag.Parse()
	return *val
}

func (i *SettingType) Bool(key string, lvl string, dVal bool) bool {
	val := flag.Bool(key, i.BoolEnv(key, lvl, dVal), "")
	flag.Parse()
	return *val
}

func (i *SettingType) Int(key string, lvl string, dVal int) int {
	val := flag.Int(key, i.IntEnv(key, lvl, dVal), "")
	flag.Parse()
	return *val
}

func (i *SettingType) StringList(key string, lvl string, dVal []string) []string {
	ddVal := ""
	for n, _ := range dVal {
		ddVal = ddVal + dVal[n] + ","
	}
	val := flag.String(key, i.StringEnv(key, lvl, ddVal), "")
	flag.Parse()
	valList := strings.Split(*val, ",")
	if valList == nil {
		i.Error = append(i.Error, envLogError(key, notSet, lvl, dVal))
		return dVal

	}
	return valList
}
