package vcago

import (
	"errors"
)

type (
	//User represents the user model
	User struct {
		ID            string   `json:"id,omitempty" bson:"_id"`
		Email         string   `json:"email" bson:"email" validate:"required,email"`
		FirstName     string   `bson:"first_name" json:"first_name" validate:"required"`
		LastName      string   `bson:"last_name" json:"last_name" validate:"required"`
		FullName      string   `bson:"full_name" json:"full_name"`
		DisplayName   string   `bson:"display_name" json:"display_name"`
		Roles         RoleList `json:"system_roles" bson:"system_roles"`
		Country       string   `bson:"country" json:"country"`
		PrivacyPolicy bool     `bson:"privacy_policy" json:"privacy_policy"`
		Confirmd      bool     `bson:"confirmed" json:"confirmed"`
		LastUpdate    string   `bson:"last_update" json:"last_update"`
	}
)

//Load loads an interface in an vcago.User model
func (i *User) Load(user interface{}) (err error) {
	var ok bool
	if i, ok = user.(*User); !ok {
		return NewStatusInternal(errors.New("not an vcago.User"))
	}
	return
}
