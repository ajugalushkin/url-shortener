package app

import (
	"fmt"
	"github.com/ajugalushkin/url-shortener/internal/config"
	"github.com/ajugalushkin/url-shortener/internal/handlers/redirect"
	"github.com/ajugalushkin/url-shortener/internal/handlers/save"
	"github.com/ajugalushkin/url-shortener/internal/service"
	"github.com/ajugalushkin/url-shortener/internal/storage"
	"github.com/labstack/echo/v4"
)

func Run() error {
	server := echo.New()

	serviceAPI := service.NewService(storage.NewInMemory())

	server.POST("/", save.New(serviceAPI))
	server.GET("/:id", redirect.New(serviceAPI))

	fmt.Println("Running server on", config.FlagRunAddr)
	err := server.Start(config.FlagRunAddr)
	if err != nil {
		return err
	}

	return nil
}
