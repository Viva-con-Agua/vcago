# vcago
The package contains standard functions that are used in the Viva-con-Agua API services and is on the [echo web framework](https://github.com/labstack/echo)


## PACKAGE

### Basic Webserver


#### Setup in server.go

```
func main() {
    e := echo.New()
    e.HTTPErrorHandler = vcago.HTTPErrorHandler    
    e.Validator = vcago.JSONValidator
    e.Use(vcago.CORS.Init())
    e.Use(vcago.Logger.Init()) 
    ...

    ...
	appPort := vcago.Config.GetEnvString("APP_PORT", "n", "1323")
	e.Logger.Fatal(e.Start(":" + appPort))

}
```
#### edit the .env file

```
SERVICE_NAME=default  //service name, default default
APP_PORT=1323           // default port 1323
LOGGING_OUTPUT=strout  // pretty, nats default strout
ALLOW_ORIGINS="https://example.com,https://api.example.com"
...
```

#### output logger pretty

```
{
    "id": "",  //not implemented
    "service": "example",
    "time": "2022-02-07T19:50:08.971851362+01:00",
    "remote_ip": "127.0.0.1",
    "host": "localhost:1323",
    "method": "POST",
    "uri": "/example",
    "user_agent": "Mozilla/5.0 (X11; Linux x86_64; rv:96.0) Gecko/20100101 Firefox/96.0",
    "status": 500,
    "error": {
        "status_message": "example error",
        "status_type": "internal_error"
    },
    "latency": 2002915,
    "latency_human": "2.002915ms",
    "byte_in": "",
    "byte_out": "",
    "modified": {
        "updated": 1644259808,
        "created": 1644259808
    }
}
```

#### Use in handler

BindAndValidate follows  [this](https://echo.labstack.com/guide/request/#validate-data) documentation. 
In the event of an error, an `ValidationError` will always be returned, 
which can be processed by the `HTTPErrorHandler`.

```
func Example(c echo.Context) (err error) {
	...
	//Validate JSON
	body := new(dao.Example)
	if vcago.BindAndValidate(c, body); err != nil {
		return
	}
    ...
```


```
func main() {
	e := echo.New()
	e.Debug = false
	// Middleware
	e.Use(vcago.Logger.Init())
	e.Use(vcago.CORS.Init())

	//error
	e.HTTPErrorHandler = vcago.HTTPErrorHandler
	e.Validator = vcago.JSONValidator
	handlers.CreateClient()
	login := e.Group("/v1/auth")
	login.POST("/callback", handlers.CallbackHandler)

	//server
	appPort := vcago.Config.GetEnvString("APP_PORT", "n", "1323")
	e.Logger.Fatal(e.Start(":" + appPort))
```

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
