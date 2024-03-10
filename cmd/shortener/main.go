package main

import (
	"github.com/ajugalushkin/url-shortener/internal/app"
	"github.com/ajugalushkin/url-shortener/internal/config"
)

func main() {
	config.ParseFlags()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
