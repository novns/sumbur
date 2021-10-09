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

	Redis redis.Options
}

var (
	config *Config
	rdb    *redis.Client
)

func Init(cfg *Config) {
	config = cfg
	rdb = redis.NewClient(&config.Redis)
}

func Check(ctx *atreugo.RequestCtx) error {
	views.AuthState = false

	session_b := ctx.Request.Header.Cookie(config.Cookie)
	if session_b == nil {
		return ctx.Next()
	}

	err := rdb.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}

	views.AuthState, err = rdb.Get(ctx, string(session_b)).Bool()
	views.AuthState = views.AuthState && err == nil

	return ctx.Next()
}

// Login

type Login struct {
	*views.BasePage
}

func AuthPost(ctx *atreugo.RequestCtx) error {
	// Session

	var session string

	session_b := ctx.Request.Header.Cookie(config.Cookie)

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

	cookie.SetKey(config.Cookie)
	cookie.SetValue(session)

	if string(ctx.FormValue("password")) == config.Password {
		cookie.SetExpire(time.Now().Add(time.Hour * config.TTL))
		rdb.SetEX(ctx, session, true, time.Hour*config.TTL)
	} else {
		cookie.SetExpire(time.Now().Add(time.Second))
		rdb.Del(ctx, session)
	}

	ctx.Response.Header.SetCookie(cookie)

	// Redirect

	return ctx.RedirectResponse(string(ctx.Referer()), 302)
}

// Restriction

type Forbidden struct {
	*views.BasePage
}

func Restrict(ctx *atreugo.RequestCtx) error {
	if views.AuthState {
		return ctx.Next()
	}

	ctx.SetStatusCode(403)
	views.WritePage(ctx, &Forbidden{})

	return nil
}

type Restricted struct {
	*views.BasePage
}

func RestrictedGet(ctx *atreugo.RequestCtx) error {
	views.WritePage(ctx, &Restricted{})
	return nil
}
