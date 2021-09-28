package auth

import (
	"github.com/go-redis/redis/v8"
	"github.com/savsgio/atreugo/v11"
)

type Config struct {
	Cookie string
}

type Auth struct {
	config *Config
	rdb    *redis.Client

	state bool
}

func Init(config *Config) *Auth {
	return &Auth{
		config: config,
		rdb:    redis.NewClient(&redis.Options{}),
	}
}

func (auth *Auth) State() bool {
	return auth.state
}

func (auth *Auth) Check() atreugo.Middleware {
	return func(ctx *atreugo.RequestCtx) error {
		session_b := ctx.Request.Header.Cookie(auth.config.Cookie)
		if session_b == nil {
			return ctx.Next()
		}

		err := auth.rdb.Ping(ctx).Err()
		if err != nil {
			panic(err)
		}

		auth.state, err = auth.rdb.Get(ctx, string(session_b)).Bool()
		auth.state = auth.state && err == nil

		return ctx.Next()
	}
}
