package app

import (
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

	server.Logger.Fatal(server.Start(":8080"))

	return nil
}
