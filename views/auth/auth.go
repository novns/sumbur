package auth

import (
	"github.com/go-redis/redis/v8"
)

type Config struct {
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
