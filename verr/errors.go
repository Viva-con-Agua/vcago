package verr

import (
	"errors"
)

var (
	ErrorInternalServer = errors.New("InternalServer")
	ErrorNotFound       = errors.New("NotFound")
	ErrorConflict       = errors.New("Conflict")
	ErrorPassword       = errors.New("Password")
	ErrorUnauthorized   = errors.New("Unauthorized")
	ErrorForbidden      = errors.New("Forbidden")
)
