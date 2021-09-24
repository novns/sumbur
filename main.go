package main

import (
	"os"
	"sumbur/views"

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

	// Server & routes

	server := atreugo.New(config.Server)

	server.GET("/", func(ctx *atreugo.RequestCtx) error {
		views.WritePage(ctx, &views.BasePage{})
		return nil
	})

	// Run

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
