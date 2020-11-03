package vmod

import (
	"time"

	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/Viva-con-Agua/vcago/vutils"
	"github.com/google/uuid"
)

type (
	//Token represents token for handlings signup, password reset and ...
	Token struct {
		ID      string `bson:"_id" json:"token_id"`
		Code    string `bson:"code" json:"code"`
		Tcase   string `json:"t_case"`
		Expired int64  `json:"expired"`
		Created int64  `json:"created"`
		ModelID string `json:"model_id" bson:"model_id"`
	}
)

//InitToken initial a Token with a 32bit random string Base64 encoded for Web handling. Use cTime for current time and set expired time max 1 month.
func InitToken(tCase string, cTime int64, expired time.Duration, modelID string) (token *Token, apiErr *verr.ApiError) {
	code, err := vutils.RandomBase64(32)
	if err != nil {
		return nil, verr.GetApiError(err, &verr.RespErrorInternalServer)
	}
	token = new(Token)
	token.ID = uuid.New().String()
	token.Code = code
	token.Tcase = tCase
	token.Expired = time.Unix(cTime, 0).Add(expired).Unix()
	token.Created = cTime
	token.ModelID = modelID
	return token, nil
}
