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

	originUrl, err := storage.GetByKey(key)
	if err != nil {
		http.Error(wrt, err.Error(), http.StatusBadRequest)
		return
	}

	wrt.Header().Set("Location", originUrl)
	wrt.WriteHeader(http.StatusTemporaryRedirect)
}
