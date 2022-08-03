package vmod

import "time"

//Modified contains update and create time for an model.
type Modified struct {
	Updated int64 `json:"updated" bson:"updated"`
	Created int64 `json:"created" bson:"created"`
}

//NewModified initial Modified model.
func NewModified() Modified {
	t := time.Now().Unix()
	m := Modified{Updated: t, Created: t}
	return m
}

//Update updates the Updated key to the current time.
func (i *Modified) Update() {
	i.Updated = time.Now().Unix()
}
