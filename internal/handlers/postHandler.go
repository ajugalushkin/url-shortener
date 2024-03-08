package handlers

import (
	"fmt"
	"github.com/ajugalushkin/url-shortener/internal/shorten"
	"github.com/ajugalushkin/url-shortener/internal/storage"
	"io"
	"net/http"
)

func PostHandler(wrt http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		GetHandler(wrt, req)
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

	UrlData, err := storage.GetByUrl(originalURL)
	if err != nil {
		UrlData = storage.URLData{
			Key: shorten.GenerateShortKey(),
			Url: originalURL}
		err = storage.Create(UrlData)
		if err != nil {
			http.Error(wrt, "ShortKey not created", http.StatusNotFound)
			return
		}
	}

	shortenedURL := fmt.Sprintf("http://localhost:8080/%s", UrlData.Key)

	wrt.Header().Set("Content-Type", "text/plan")
	wrt.WriteHeader(http.StatusCreated)
	wrt.Write([]byte(shortenedURL))
}
