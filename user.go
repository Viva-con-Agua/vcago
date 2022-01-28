package vcago

import (
	"errors"
)

type (
	//User represents the user model
	User struct {
		ID       string   `json:"id,omitempty" bson:"_id"`
		Email    string   `json:"email" bson:"email" validate:"required,email"`
		Profile  Profile  `json:"profile" bson:"profile"`
		Roles    string   `json:"roles" bson:"roles"`
		Modified Modified `json:"modified" bson:"modified"`
	}
	//Profile contains the profile information of an user model
	Profile struct {
		FirstName   string `bson:"first_name" json:"first_name" validate:"required"`
		LastName    string `bson:"last_name" json:"last_name" validate:"required"`
		FullName    string `bson:"full_name" json:"full_name"`
		DisplayName string `bson:"display_name" json:"display_name"`
		Gender      string `bson:"gender" json:"gender"`
		Country     string `bson:"country" json:"country"`
		Avatar      Avatar `bson:"avatar" json:"avatar"`
	}
	//Avatar contains the link and picturetype of the user avatar
	Avatar struct {
		URL  string `bson:"url" json:"url"`
		Type string `bson:"type" json:"type"`
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
