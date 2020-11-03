package vmod

import (
	"time"

	"github.com/Viva-con-Agua/vcago/vutils"
	"github.com/google/uuid"
)

type (
	Token struct {
		ID      string `bson:"_id" json:"token_id"`
		Code    string `bson:"code" json:"code"`
		Tcase   string `json:"t_case"`
		Expired int64  `json:"expired"`
		Created int64  `json:"created"`
		ModelId string `json:"model_id" bson:"model_id"`
	}
)

func InitToken(t_case string, c_time int64, expired time.Duration, model_id string) (token *Token, err error) {
	code, err := vutils.RandomBase64(32)
	if err != nil {
		return nil, err
	}
	token = new(Token)
	token.ID = uuid.New().String()
	token.Code = code
	token.Tcase = t_case
	token.Expired = time.Unix(c_time, 0).Add(expired).Unix()
	token.Created = c_time
	token.ModelId = model_id
	return token, err
}
