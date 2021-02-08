package verr

import (
	"log"
	"os"
    "strconv"
)

//GetEnvWithLog Validate .env Variables
func GetEnvWithLog(key string) (string, bool){
    val, ok := os.LookupEnv(key)
    if !ok {
        log.Printf("Warning: %s not set\n", key)
        return "", ok
    }
    return val, ok
}
//GetIntEnvWithLog Validate .env Variables
func GetIntEnvWithLog(key string) (int, bool){
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
