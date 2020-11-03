# vca-go
Package for handling vca-api

## vcago package
```
CONSTANTS

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

VARIABLES

var CORSConfig = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins:     strings.Split(os.Getenv("ALLOW_ORIGINS"), ","),
	AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	AllowCredentials: true,
})
    CORSConfig for api services. Can be configured via .env.


FUNCTIONS

func DeleteSession(c echo.Context)
    DeleteSession remove sessin for context c from Redis.

func GetSessionUser(c echo.Context) (u *vmod.User, contains bool)
    GetSessionUser return user for c or false if c has no user in Redis

func InitSession(c echo.Context, user *vmod.User)
    InitSession initial a session for a vmod.User via Redis

func JSONErrorHandler(c echo.Context, i interface{}) (rErr *verr.ResponseError)
    JSONErrorHandler formats JsonError to ResponseError

func LogAPIError(e *verr.ApiError, c echo.Context, i interface{})
    LogApiError via log.Print

func RedisSession() echo.MiddlewareFunc
    RedisSession middleware initial redis store for session handling

func ResponseErrorHandler(c echo.Context, apiErr *verr.ApiError, i interface{}) (rErr *verr.ResponseError)
    ResponseErrorHandler handles ApiError

func SessionAuth(next echo.HandlerFunc) echo.HandlerFunc
    SessionAuth go to next if the request has a session else return 401.

func formatRequestPrint(r *http.Request) string

TYPES

type Validator struct {
	Validator *validator.Validate
}
    Validator represents a Json validator

func (cv *Validator) Validate(i interface{}) error
    Validate interface i
```
