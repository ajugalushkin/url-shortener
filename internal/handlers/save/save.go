package save

import (
	"fmt"
	"github.com/ajugalushkin/url-shortener/internal/config"
	"github.com/ajugalushkin/url-shortener/internal/model"
	"github.com/ajugalushkin/url-shortener/internal/service"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func New(serviceAPI *service.Service) echo.HandlerFunc {
	return func(context echo.Context) error {
		if context.Request().Method != http.MethodPost {
			return context.String(http.StatusBadRequest, "Wrong type request")
		}

		body, err := io.ReadAll(context.Request().Body)
		if err != nil {
			return context.String(http.StatusBadRequest, "URL parse error")
		}

		originalURL := string(body)
		if originalURL == "" {
			return context.String(http.StatusBadRequest, "URL parameter is missing")
		}

		shortenURL, err := serviceAPI.Shorten(model.ShortenInput{RawURL: originalURL})
		if err != nil {
			return context.String(http.StatusBadRequest, "URL not shortening")
		}

		shortenedURL := fmt.Sprintf("http://%s/%s", config.BaseURL, shortenURL.Key)

		context.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlain)
		context.Response().Status = http.StatusCreated
		context.Response().Write([]byte(shortenedURL))

		return context.String(http.StatusCreated, "")
	}
}

/*func New(serviceAPI *service.Service) http.HandlerFunc {
	return func(wrt http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(wrt, "Invalid request method", http.StatusBadRequest)
			return
		}

		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(wrt, "URL parse error", http.StatusBadGateway)
			return
		}

		originalURL := string(body)
		if originalURL == "" {
			http.Error(wrt, "URL parameter is missing", http.StatusBadRequest)
			return
		}

		shortenURL, err := serviceAPI.Shorten(model.ShortenInput{RawURL: originalURL})
		if err != nil {
			http.Error(wrt, "URL not shortening", http.StatusBadRequest)
			return
		}

		shortenedURL := fmt.Sprintf("http://localhost:8080/%s", shortenURL.Key)

		wrt.Header().Set("Content-Type", "text/plain")
		wrt.WriteHeader(http.StatusCreated)
		wrt.Write([]byte(shortenedURL))
	}
}*/
