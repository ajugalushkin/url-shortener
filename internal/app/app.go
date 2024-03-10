package app

import (
	"fmt"
	"github.com/ajugalushkin/url-shortener/internal/config"
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

	fmt.Println("Running server on", config.FlagRunAddr)
	err := http.ListenAndServe(config.FlagRunAddr, mux)
	if err != nil {
		return err
	}
	return nil
}
