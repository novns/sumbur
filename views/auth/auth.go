package auth

type Auth struct {
	state bool
}

func (auth *Auth) State() bool {
	return auth.state
}
