package auth

import (
	"crypto/rand"
	"encoding/hex"
	"sumbur/views"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

type Config struct {
	Cookie   string
	Password string
	TTL      time.Duration
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

// Login

type Login struct {
	*views.BasePage
}

func LoginView(auth *Auth) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		// Session

		var session string

		session_b := ctx.Request.Header.Cookie(auth.config.Cookie)

		if session_b == nil {
			random := make([]byte, 16)
			rand.Read(random)
			session = hex.EncodeToString(random)
		} else {
			session = string(session_b)
		}

		// Cookie

		cookie := fasthttp.AcquireCookie()
		defer fasthttp.ReleaseCookie(cookie)

		cookie.SetKey(auth.config.Cookie)
		cookie.SetValue(session)

		if string(ctx.FormValue("password")) == auth.config.Password {
			cookie.SetExpire(time.Now().Add(time.Hour * auth.config.TTL))
			auth.rdb.SetEX(ctx, session, true, time.Hour*auth.config.TTL)
		} else {
			cookie.SetExpire(time.Now().Add(time.Second))
			auth.rdb.Del(ctx, session)
		}

		ctx.Response.Header.SetCookie(cookie)

		// Redirect

		return ctx.RedirectResponse(string(ctx.Referer()), 302)
	}
}

// Restriction

type Forbidden struct {
	*views.BasePage
}

func (auth *Auth) Restrict() atreugo.Middleware {
	return func(ctx *atreugo.RequestCtx) error {
		if auth.state {
			return ctx.Next()
		}

		ctx.SetStatusCode(403)
		views.WritePage(ctx, &Forbidden{}, auth)

		return nil
	}
}

type Restricted struct {
	*views.BasePage
}

func RestrictedView(auth *Auth) atreugo.View {
	return func(ctx *atreugo.RequestCtx) error {
		views.WritePage(ctx, &Restricted{}, auth)
		return nil
	}
}
