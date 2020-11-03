package vcago

import (
	"encoding/gob"
	"log"
	"os"

	"github.com/Viva-con-Agua/vcago/redisstore"
	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

//RedisSession middleware initial redis store for session handling
func RedisSession() echo.MiddlewareFunc {
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
