package vmod

type (

	//Modified represents update and create timestamp for all models.
	Modified struct {
		Updated int64 `json:"updated" bson:"updated,omitempty"`
		Created int64 `json:"created" bson:"created,omitempty"`
	}
)

// InitModified initial Modified model, cTime is the current time Unix format.
func InitModified(cTime int64) *Modified {
	m := Modified{Updated: cTime, Created: cTime}
	return &m
}
