package vmod

import "time"

type (
	//Modified represents update and create timestamp for all models.
	Modified struct {
		Updated int64 `json:"updated" bson:"updated,omitempty"`
		Created int64 `json:"created" bson:"created,omitempty"`
	}
)

//NewModified initial Modified model, cTime is the current time Unix format.
func NewModified() *Modified {
	cTime := time.Now().Unix()
	m := Modified{Updated: cTime, Created: cTime}
	return &m
}

//Update set new Updated
func (modified *Modified) Update() *Modified {
	modified.Updated = time.Now().Unix()
	return modified
}
