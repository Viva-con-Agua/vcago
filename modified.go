package vcago

import "time"

type Modified struct {
	Updated int64 `json:"updated" bson:"updated"`
	Created int64 `json:"created" bson:"created"`
}

//NewModified initial Modified model, cTime is the current time Unix format.
func NewModified() Modified {
	t := time.Now().Unix()
	m := Modified{Updated: t, Created: t}
	return m
}

func (i *Modified) Update() {
	i.Updated = time.Now().Unix()
}
