package main

import (
	"os"
	"sumbur/views"
	"sumbur/views/auth"
	"sumbur/views/http_errors"

	"github.com/savsgio/atreugo/v11"
	"gopkg.in/yaml.v2"
)

type Config struct {
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

	server := atreugo.New(config.Server)

	au := &auth.Auth{}

	config.Server.MethodNotAllowedView = http_errors.NotFoundView(au)
	config.Server.NotFoundView = http_errors.NotFoundView(au)
	config.Server.PanicView = http_errors.PanicView(au)

	// Routes

	server.GET("/", func(ctx *atreugo.RequestCtx) error {
		views.WritePage(ctx, &views.BasePage{}, au)
		return nil
	})

	server.GET("/panic", func(ctx *atreugo.RequestCtx) error {
		panic("A drum, a drum! Panic doth come.")
	})

	// Run

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
