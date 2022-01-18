package vutils

import (
	"time"

	"github.com/google/uuid"
)

type (
	LinkToken struct {
		ID        string `bson:"_id" json:"id"`
		Code      string `bson:"code" json:"code"`
		ModelID   string `bson:"model_id" json:"model_id"`
		Scope     string `bson:"scope" json:"scope"`
		ExpiresAt int64  `bson:"expires_at" json:"expires_at"`
		Created   int64  `bson:"created" json:"created"`
	}
)

func NewLinkToken(scope string, expired time.Duration, modelID string) (r *LinkToken, err error) {
	var code string
	code, err = RandomBase64(32)
	if err != nil {
		return
	}
	r = &LinkToken{
		ID:        uuid.New().String(),
		Code:      code,
		Scope:     scope,
		ExpiresAt: time.Now().Add(expired).Unix(),
		Created:   time.Now().Unix(),
		ModelID:   modelID,
	}
	return
}

func (r *LinkToken) NewCode(expired time.Duration) (err error) {
	var code string
	if code, err = RandomBase64(32); err != nil {
		return
	}
	now := time.Now()
	r.Code = code
	r.ExpiresAt = now.Add(expired).Unix()
	r.Created = now.Unix()
	return
}
