package vcago

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/Viva-con-Agua/vcago/redisstore"
	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

//SessionRedisStore initial session store via Redis and return session.Middleware.
//Use with echo framework like echo.Echo.Use(SessionRedisStore())
func SessionRedisStore() echo.MiddlewareFunc {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})
	redis, err := redisstore.NewRedisStore(client)

	if err != nil {
		log.Fatal("failed to create redis store: ", err)
	}
	gob.Register(&vmod.User{})
	log.Println("Redis successfully connected!")
	return session.Middleware(redis)
}

//NewSession creates a cookie using the `COOKIE_SECURE` and `SAME_SITE` variables from the os environment.
//For that use .env file in your project. The cookie is always be httpOnly and have a durancy of 7 days.
//The `user` will be stored in that session as []byte generated with json.Marshal(user).
func NewSession(c echo.Context, user *vmod.User) {
	secure := true
	if os.Getenv("COOKIE_SECURE") == "false" {
		secure = false
	}
	sameSite := http.SameSiteNoneMode
	if os.Getenv("SAME_SITE") == "lax" {
		sameSite = http.SameSiteLaxMode
	}
	if os.Getenv("SAME_SITE") == "none" {
		sameSite = http.SameSiteNoneMode
	}
	if os.Getenv("SAME_SITE") == "strict" {
		sameSite = http.SameSiteStrictMode
	}
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		SameSite: sameSite,
		Secure:   secure,
	}
	sessionUser, _ := json.Marshal(user)
	sess.Values["valid"] = true
	sess.Values["user"] = &sessionUser
	sess.Save(c.Request(), c.Response())
}

//GetSession selects u from the session storage by using c. In case there is no user contains is set false.
func GetSession(c echo.Context) (apiErr *verr.APIError) {
	sess, _ := session.Get("session", c)
	val := sess.Values["user"]
	var user []byte
	user, contains := val.([]byte)
	if contains == false {
		return verr.NewAPIError(errors.New("no user found in session storage")).Unauthorized()
	}
	u := new(vmod.User)
	json.Unmarshal(user, &u)
	c.Set("user", u)
	return nil

}

//DeleteSession removes session from storage using c.
func DeleteSession(c echo.Context) {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	sess.Values["valid"] = nil
	sess.Values["user"] = nil
	sess.Save(c.Request(), c.Response())
}

//SessionAuth go to next if the request has a session else return 401.
func SessionAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiErr := GetSession(c)
		if apiErr != nil {
			apiErr.Log(c)
			return echo.NewHTTPError(http.StatusUnauthorized, "")
		}
		return next(c)
	}
}
