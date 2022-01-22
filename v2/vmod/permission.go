package vmod

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type (
	//Permission represents role to an access
	Permission map[string]AccessList

	//Access represents access to an model.
	Access struct {
		Role    string `json:"role" bson:"role"`
		Created int64  `json:"created" bson:"created"`
	}
	//AccessList list of access model
	AccessList []Access
	//PValidate used for controller permission handling
	PValidate map[string][]string
)

//NewPermission map role to []Access(modelID, created) and initial Permission.
func NewPermission(app string, role string) *Permission {
	access := new(Access)
	permission := make(Permission)
	access.Role = role
	access.Created = time.Now().Unix()
	permission[app] = append(permission[app], *access)
	return &permission
}

//Add Access(modelID created) to role.
func (p *Permission) Add(app string, role string) *Permission {
	s := (*p)[role]
	for _, v := range s {
		if v.Role == role {
			return p
		}
	}
	access := new(Access)
	access.Role = role
	access.Created = time.Now().Unix()
	(*p)[role] = append((*p)[role], *access)
	return p
}

//Delete Access from role. If []Access is nil the role will remove form Permission.
func (p *Permission) Delete(role string, modelID string) *Permission {
	s := (*p)[role]
	for i, v := range s {
		if v.Role == modelID {
			s = append(s[:i], s[i+1:]...)
			break
		}
	}
	(*p)[role] = s
	if s == nil {
		delete((*p), role)
	}
	return p
}

//Map converts an access list into map
func (a AccessList) Map() map[string]string {
	confMap := map[string]string{}
	for _, v := range a {
		confMap[v.Role] = v.Role
	}
	return confMap
}

//Validate validates the permission for middleware functions
func (p *Permission) Validate(pValidate *PValidate) bool {
	for k, v := range *pValidate {
		if val, ok := (*p)[k]; ok {
			aMap := val.Map()
			for i := range v {
    			if _, ok := aMap[v[i]]; ok {
					return true
    			}
			}
		}
	}
	return false
}

//Restricted middleware function for handling Access
func (pVal *PValidate) Restricted(next echo.HandlerFunc) echo.HandlerFunc {
	return func (c echo.Context) error {
		user := c.Get("token").(*jwt.Token)
		claims := user.Claims.(*AccessToken)
		if ok := claims.User.Permission.Validate(pVal); ok {
			return next(c)
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "permission_denied")
	}
}
