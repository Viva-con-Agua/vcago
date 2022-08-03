package vmod

import (
	"errors"
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
		DisplayName   string   `bson:"display_name" json:"display_name"`
		Roles         RoleList `json:"system_roles" bson:"system_roles"`
		Country       string   `bson:"country" json:"country"`
		PrivacyPolicy bool     `bson:"privacy_policy" json:"privacy_policy"`
		Confirmd      bool     `bson:"confirmed" json:"confirmed"`
		LastUpdate    string   `bson:"last_update" json:"last_update"`
	}
)

//CheckUpdate checks if the lastUpdate time string is older as the users LastUpdate param.
//If the function return true, the user needs to be updated in this service.
func (i *User) CheckUpdate(lastUpdate string) bool {
	current, _ := time.Parse(time.RFC3339, i.LastUpdate)
	last, _ := time.Parse(time.RFC3339, lastUpdate)
	if current.Unix() > last.Unix() {
		return true
	}
	return false
}

//Load loads an interface in an vcago.User model
func (i *User) Load(user interface{}) (err error) {
	var ok bool
	if i, ok = user.(*User); !ok {
		return errors.New("not an vcago.User")
	}
	return
}
