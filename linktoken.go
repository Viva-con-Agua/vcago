package vcago

import (
	"time"

	"github.com/google/uuid"
)

type (
	//LinkToken is used for handling link with token
	LinkToken struct {
		ID        string   `json:"id" bson:"_id"`
		Code      string   `json:"code" bson:"code"`
		ExpiresAt int64    `json:"expires_at" bson:"expires_at"`
		Scope     string   `json:"scope" bson:"scope"`
		UserID    string   `json:"user_id" bson:"user_id"`
		Modified  Modified `json:"modified" bson:"modified"`
	}
)

//NewLinkToken initial a Token with a 32bit random string Base64 encoded for Web handling. Set expired time max 1 month.
func NewLinkToken(expired time.Duration, userID string, scope string) (*LinkToken, error) {
	code, err := RandomBase64(32)
	if err != nil {
		return nil, err
	}
	return &LinkToken{
		ID:        uuid.New().String(),
		Code:      code,
		ExpiresAt: time.Now().Add(expired).Unix(),
		Scope:     scope,
		UserID:    userID,
		Modified:  NewModified(),
	}, nil
}

//NewCode generate a new code for LinkTokens
func (l *LinkToken) NewCode(expired time.Duration) (*LinkToken, error) {
	code, err := RandomBase64(32)
	if err != nil {
		return nil, err
	}
	l.Code = code
	l.ExpiresAt = time.Now().Add(expired).Unix()
	return l, nil
}
