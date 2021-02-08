package vmod

import (
	"context"
	"time"

	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/Viva-con-Agua/vcago/vutils"
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
func NewLinkToken(ctx context.Context, tCase string, expired time.Duration, modelID string, scope string) (*LinkToken, error) {
	code, err := vutils.RandomBase64(32)
	if err != nil {
		return nil, verr.InternalServerError(ctx, err)
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
func (l *LinkToken) NewCode(ctx context.Context, expired time.Duration) (*LinkToken, error) {
	code, err := vutils.RandomBase64(32)
	if err != nil {
		return nil, verr.InternalServerError(ctx, err)
	}
	l.Code = code
	l.Expired = time.Now().Add(expired).Unix()
	l.Created = time.Now().Unix()
	return l, nil
}
