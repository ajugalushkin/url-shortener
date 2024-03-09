package app

import (
	"github.com/ajugalushkin/url-shortener/internal/handlers/middleware"
	"github.com/ajugalushkin/url-shortener/internal/handlers/redirect"
	"github.com/ajugalushkin/url-shortener/internal/handlers/save"
	"github.com/ajugalushkin/url-shortener/internal/service"
	"github.com/ajugalushkin/url-shortener/internal/storage"
	"net/http"
)

func Run() error {
	mux := http.NewServeMux()

	serviceAPI := service.NewService(storage.NewInMemory())

	mux.Handle("/", middleware.Switch(save.New(serviceAPI), redirect.New(serviceAPI)))

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		return err
	}
	return nil
}
