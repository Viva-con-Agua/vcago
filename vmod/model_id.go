package vmod

type ModelID struct {
	MID  string `json:"id" bson:"id"`
	Type string `json:"type" bson:"type"`
}

func NewModelID(mID string, mType string) *ModelID {
	return &ModelID{
		MID:  mID,
		Type: mType,
	}
}
