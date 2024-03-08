package handlers

import (
	"fmt"
	"github.com/ajugalushkin/url-shortener/internal/model"
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

	storageApi, errGetApi := storage.NewStorage()
	if errGetApi != nil {
		http.Error(wrt, "Storage not found!", http.StatusBadRequest)
		return
	}

	UrlData, errGet := storageApi.RetrieveByURL(originalURL)
	if errGet != nil {
		UrlData = model.URLData{
			Key: shorten.GenerateShortKey(),
			Url: originalURL}
		_, err = storageApi.Insert(UrlData)
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
