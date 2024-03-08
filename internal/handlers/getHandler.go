package handlers

import (
	"github.com/ajugalushkin/url-shortener/internal/storage"
	"net/http"
	"strings"
)

func GetHandler(wrt http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(wrt, "Invalid request method", http.StatusBadRequest)
		return
	}

	key := strings.Replace(req.URL.Path, "/", "", -1)

	storageApi, errGetApi := storage.NewStorage()
	if errGetApi != nil {
		http.Error(wrt, "Storage not found!", http.StatusBadRequest)
		return
	}

	dataURL, err := storageApi.Retrieve(key)
	if err != nil {
		http.Error(wrt, err.Error(), http.StatusBadRequest)
		return
	}

	wrt.Header().Set("Location", dataURL.Url)
	wrt.WriteHeader(http.StatusTemporaryRedirect)
}
