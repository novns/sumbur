package main

import (
	"os"
	"sumbur/db"
	"sumbur/views/auth"
	"sumbur/views/blog"
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

			MethodNotAllowedView: http_errors.NotFoundView,
			NotFoundView:         http_errors.NotFoundView,
			PanicView:            http_errors.PanicView,
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

	auth.Init(&config.Auth)

	db.DSN = &config.DB

	server := atreugo.New(config.Server)

	server.UseBefore(auth.Check)

	// Routes

	server.GET("/", blog.BlogGet)
	server.GET("/tag/{tag}", blog.BlogGet)

	server.GET(`/edit/{article_id:\d+}`, blog.EditGet).
		UseBefore(auth.Restrict)

	server.POST(`/edit/{article_id:\d+}`, blog.EditPost).
		UseBefore(auth.Restrict)

	server.POST("/auth", auth.AuthPost)

	server.GET("/restrict", auth.RestrictedGet).
		UseBefore(auth.Restrict)

	server.GET("/panic", func(ctx *atreugo.RequestCtx) error {
		panic("A drum, a drum! Panic doth come.")
	})

	server.Static("/static", "static")

	// Run

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
