package vmod

import "errors"

type (
	//User represents the user default user information they are shared with all viva con agua services.
	Organisation struct {
		ID           string `json:"id" bson:"_id"`
		Email        string `json:"email" bson:"email"`
		Name         string `json:"name" bson:"name"`
		Country      string `json:"country" bson:"country"`
		Abbreviation string `json:"abbreviation" bson:"abbreviation"`
		DefaultAspID string `json:"default_asp_id" bson:"default_asp_id"`
	}
)

// Load loads an interface in an vcago.User model
func (i *Organisation) Load(organisation interface{}) (err error) {
	var ok bool
	if i, ok = organisation.(*Organisation); !ok {
		return errors.New("not an vcago.Organisation")
	}
	return
}
