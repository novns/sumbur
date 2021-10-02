package main

import (
	"os"
	"sumbur/db"
	"sumbur/views"
	"sumbur/views/auth"
	"sumbur/views/http_errors"

	"github.com/savsgio/atreugo/v11"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Auth   auth.Config
	DB     string
	Server atreugo.Config
}

func main() {
	// Configuration

	config := Config{
		Server: atreugo.Config{
			NoDefaultContentType: true,
		},
	}

	file, err := os.ReadFile(os.Getenv("SUMBUR_CONFIG"))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	// Server

	au := auth.Init(&config.Auth)

	config.Server.MethodNotAllowedView = http_errors.NotFoundView(au)
	config.Server.NotFoundView = http_errors.NotFoundView(au)
	config.Server.PanicView = http_errors.PanicView(au)

	db.DSN = &config.DB

	server := atreugo.New(config.Server)

	server.UseBefore(au.Check())

	// Routes

	server.GET("/", func(ctx *atreugo.RequestCtx) error {
		views.WritePage(ctx, &views.BasePage{}, au)
		return nil
	})

	server.POST("/auth", auth.LoginView(au))

	server.GET("/restrict", auth.RestrictedView(au)).
		UseBefore(au.Restrict())

	server.GET("/panic", func(ctx *atreugo.RequestCtx) error {
		panic("A drum, a drum! Panic doth come.")
	})

	// Run

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
