package vcago

import (
	"time"

	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/golang-jwt/jwt"
)

//AccessToken represents the default access_token contains the basic user informations.
type AccessToken struct {
	ID            string              `json:"id,omitempty" bson:"_id"`
	Email         string              `json:"email" bson:"email" validate:"required,email"`
	FirstName     string              `json:"first_name" validate:"required"`
	LastName      string              `json:"last_name" validate:"required"`
	FullName      string              `json:"full_name"`
	DisplayName   string              `json:"display_name"`
	Roles         vmod.RoleListCookie `json:"system_roles"`
	Country       string              `json:"country"`
	PrivacyPolicy bool                `json:"privacy_policy"`
	Confirmed     bool                `json:"confirmed"`
	LastUpdate    string              `json:"last_update"`
	jwt.StandardClaims
}

//NewAccessToken creates an new access_token from vmod.User model.
func NewAccessToken(user *vmod.User) *AccessToken {
	return &AccessToken{
		user.ID,
		user.Email,
		user.FirstName,
		user.LastName,
		user.FullName,
		user.DisplayName,
		*user.Roles.Cookie(),
		user.Country,
		user.PrivacyPolicy,
		user.Confirmd,
		user.LastUpdate,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
}

//SignedString is used for Sign an accesstoken based on the secret param.
func (i *AccessToken) SignedString(secret string) (string, error) {
	temp := jwt.NewWithClaims(jwt.SigningMethodHS256, i)
	return temp.SignedString([]byte(secret))
}
