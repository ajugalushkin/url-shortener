package main

import "github.com/ajugalushkin/url-shortener/internal/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
