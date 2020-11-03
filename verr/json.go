package verr

import (
	"strings"
)

type (
	JsonError struct {
		Key   string
		Error string
	}
)

func JsonErrorResponse(e error) (j_list []JsonError) {
	json_error := new(JsonError)
	error_list := strings.Split(e.Error(), "\n")
	for _, val := range error_list {
		e_list := strings.Split(val, "Key: ")
		e_list = strings.Split(e_list[1], " Error:")
		json_error.Key = e_list[0]
		json_error.Error = e_list[1]
		j_list = append(j_list, *json_error)
	}
	return j_list
}
