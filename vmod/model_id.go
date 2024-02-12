package vmod

type ModelID struct {
	MID  string `json:"m_id" bson:"m_id"`
	Type string `json:"type" bson:"type"`
}

func NewModelID(mID string, mType string) *ModelID {
	return &ModelID{
		MID:  mID,
		Type: mType,
	}
}
