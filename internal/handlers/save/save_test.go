package save

import (
	"github.com/ajugalushkin/url-shortener/internal/service"
	"github.com/ajugalushkin/url-shortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostHandler(t *testing.T) {
	type want struct {
		code           int
		body           string
		reqcontentType string
		retcontentType string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "Test StatusCreated",
			want: want{
				code:           http.StatusCreated,
				body:           "https://practicum.yandex.ru/",
				reqcontentType: "text/plain",
				retcontentType: "text/plain",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body := strings.NewReader(test.want.body)
			request := httptest.NewRequest(http.MethodPost, "/", body)
			request.Header.Set("Content-Type", test.want.reqcontentType)

			writer := httptest.NewRecorder()

			handler := New(service.NewService(storage.NewInMemory()))
			handler.ServeHTTP(writer, request)

			result := writer.Result()

			assert.Equal(t, test.want.code, result.StatusCode)
			assert.Equal(t, test.want.retcontentType, result.Header.Get("Content-Type"))
		})
	}

	tests = []struct {
		name string
		want want
	}{
		{
			name: "Test Wrong Content-Type",
			want: want{
				code:           http.StatusBadRequest,
				body:           "https://practicum.yandex.ru/",
				reqcontentType: "application/json",
			},
		},
		{
			name: "Test Wrong Request",
			want: want{
				code: http.StatusBadRequest,
				body: "https://practicum.yandex.ru/",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			body := strings.NewReader(test.want.body)
			request := httptest.NewRequest(http.MethodGet, "/", body)
			request.Header.Set("Content-Type", test.want.reqcontentType)

			writer := httptest.NewRecorder()

			handler := New(service.NewService(storage.NewInMemory()))
			handler.ServeHTTP(writer, request)

			result := writer.Result()

			assert.Equal(t, test.want.code, result.StatusCode)
		})
	}
}
