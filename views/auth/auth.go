package auth

type Config struct {
}

type Auth struct {
	config *Config

	state bool
}

func Init(config *Config) *Auth {
	return &Auth{config: config}
}

func (auth *Auth) State() bool {
	return auth.state
}
