package vcago

import (
	"context"
	"encoding/json"
	"os"

	"github.com/Viva-con-Agua/vcago/verr"
	"github.com/Viva-con-Agua/vcago/vmod"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var JWTStore = redis.Client{}

func NewJWTStore() {
	JWTStore = *redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})
}

func InsertJWT(t string, user vmod.User) (apiErr *verr.APIError) {
	u, _ := json.Marshal(user)
	err := JWTStore.Set(ctx, t, u, 0).Err()
	if err != nil {
		return verr.NewAPIError(err).InternalServerError()
	}
	return nil
}

func GetJWT(t string) (user *vmod.User, apiErr *verr.APIError) {
	u, err := JWTStore.Get(ctx, t).Result()
	if err != nil {
		return nil, verr.NewAPIError(err).InternalServerError()
	}
	user = new(vmod.User)
	json.Unmarshal([]byte(u), &user)
	return user, nil
}
