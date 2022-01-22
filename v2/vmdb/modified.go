package vmdb

import "time"

type (
	//Modified represents update and create timestamp for all models.
	Modified struct {
		count   int64 `json:"count" bson:"count"`
		Updated int64 `json:"updated" bson:"updated`
		Created int64 `json:"created" bson:"created"`
	}
)

//NewModified initial Modified model, cTime is the current time Unix format.
func NewModified() *Modified {

	current := time.Now().Unix()
	m := Modified{count: 0, Updated: current, Created: current}
	return &m
}

//Update set new Updated
func (i *Modified) Update() {
	i.count = i.count + 1
	i.Updated = time.Now().Unix()
}
