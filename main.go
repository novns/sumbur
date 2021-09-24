package main

import (
	"os"
	"sumbur/views"
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

			NotFoundView: http_errors.GetNotFound,
			PanicView:    http_errors.GetPanic,
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

	// Server & routes

	server := atreugo.New(config.Server)

	server.GET("/", func(ctx *atreugo.RequestCtx) error {
		views.WritePage(ctx, &views.BasePage{})
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
