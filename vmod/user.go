package vmod

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

type (
	//User represents the user default user information they are shared with all viva con agua services.
	User struct {
		ID            string   `json:"id,omitempty" bson:"_id"`
		Email         string   `json:"email" bson:"email"`
		FirstName     string   `bson:"first_name" json:"first_name"`
		LastName      string   `bson:"last_name" json:"last_name"`
		FullName      string   `bson:"full_name" json:"full_name"`
		RealName      string   `bson:"real_name" json:"real_name"`
		DisplayName   string   `bson:"display_name" json:"display_name"`
		Roles         RoleList `json:"system_roles" bson:"system_roles"`
		Country       string   `bson:"country" json:"country"`
		PrivacyPolicy bool     `bson:"privacy_policy" json:"privacy_policy"`
		Confirmd      bool     `bson:"confirmed" json:"confirmed"`
		LastUpdate    string   `bson:"last_update" json:"last_update"`
	}
	DeletionRequest struct {
		ID       string   `json:"id" bson:"_id"`
		Email    string   `json:"email" bson:"email"`
		UserID   string   `json:"user_id" bson:"user_id"`
		Modified Modified `json:"modified" bson:"modified"`
	}
	DeletionResponse struct {
		ID            string            `bson:"_id" json:"id"`
		UserID        string            `bson:"user_id" json:"user_id"`
		Service       string            `bson:"service" json:"service"`
		StatusMessage string            `bson:"status_message" json:"status_message"`
		StatusType    string            `bson:"status_type" json:"status_type"`
		Data          map[string]string `bson:"data" json:"data"`
	}
)

// CheckUpdate checks if the lastUpdate time string is older as the users LastUpdate param.
// If the function return true, the user needs to be updated in this service.
func (i *User) CheckUpdate(lastUpdate string) bool {
	current, _ := time.Parse(time.RFC3339, i.LastUpdate)
	last, _ := time.Parse(time.RFC3339, lastUpdate)
	return current.Unix() > last.Unix()
}

// Load loads an interface in an vcago.User model
func (i *User) Load(user interface{}) (err error) {
	var ok bool
	if i, ok = user.(*User); !ok {
		return errors.New("not an vcago.User")
	}
	return
}

func WalkStruct(data interface{}, prefix string, result *map[string]string) map[string]string {
	walk(reflect.ValueOf(data), prefix, *result)
	return *result
}

func walk(v reflect.Value, prefix string, result map[string]string) {
	// Dereference pointer if needed
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			fieldValue := v.Field(i)

			key := field.Name
			if prefix != "" {
				key = prefix + "." + key
			}

			walk(fieldValue, key, result)
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i)
			key := fmt.Sprintf("%s[%d]", prefix, i)
			walk(item, key, result)
		}

	default:
		// Convert primitive values into strings
		result[prefix] = fmt.Sprintf("%v", v.Interface())
	}
}
