package app

import (
	"github.com/ajugalushkin/url-shortener/internal/handlers"
	"net/http"
)

func Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.PostHandler)
	//mux.HandleFunc("/{id}", handlers.GetHandler)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		return err
	}
	return nil
}
