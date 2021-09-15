package vcago

import (
	"time"

	"github.com/google/uuid"
)

type (
	//LinkToken is used for handling link with token
	LinkToken struct {
		ID      string `bson:"_id" json:"token_id"`
		Code    string `bson:"code" json:"code"`
		Tcase   string `json:"t_case"`
		Scope   string `json:"scope"`
		Expired int64  `json:"expired"`
		Created int64  `json:"created"`
		ModelID string `json:"model_id" bson:"model_id"`
	}
)

//NewLinkToken initial a Token with a 32bit random string Base64 encoded for Web handling. Set expired time max 1 month.
func NewLinkToken(tCase string, expired time.Duration, modelID string, scope string) (*LinkToken, error) {
	code, err := RandomBase64(32)
	if err != nil {
		return nil, err
	}
	return &LinkToken{
		ID:      uuid.New().String(),
		Code:    code,
		Tcase:   tCase,
		Scope:   scope,
		Expired: time.Now().Add(expired).Unix(),
		Created: time.Now().Unix(),
		ModelID: modelID,
	}, nil
}

//NewCode generate a new code for LinkTokens
func (l *LinkToken) NewCode(expired time.Duration) (*LinkToken, error) {
	code, err := RandomBase64(32)
	if err != nil {
		return nil, err
	}
	l.Code = code
	l.Expired = time.Now().Add(expired).Unix()
	l.Created = time.Now().Unix()
	return l, nil
}
