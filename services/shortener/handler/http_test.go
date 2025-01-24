package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockShortenerService struct{}

func (s *MockShortenerService) ShortenUrl(ctx context.Context, longUrl string) (string, error) {
	return "14q60P", nil
}

func (s *MockShortenerService) GetRedirectUrl(ctx context.Context, shortUrl string) (string, error) {
	return "http://test.com/", nil
}

func TestShortenerHttpHandler(t *testing.T) {
	h := NewShortenerHttpHandler(&MockShortenerService{})

	t.Run("HandleGetLongUrl", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/s/14q60P", nil)

		h.HandleGetLongUrl(h.shortenerService).ServeHTTP(w, r)

		assertStatus(t, w.Code, http.StatusOK)
		assertHeader(t, w.Header().Get("Content-Type"), "application/json")

		var res map[string]interface{}
		if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
			t.Fatalf("failed decoding json: %v", err)
		}

		url, present := res["url"]
		if !present {
			t.Fatalf("response body does not contain key \"url\"")
		}

		want := "http://test.com/"
		if url != want {
			t.Fatalf("did not get correct redirect url, got %v, want %v", url, want)
		}
	})

	t.Run("HandlePostShortenUrl", func(t *testing.T) {
		body := map[string]interface{}{"url": "http://test.com/"}
		buff := bytes.NewBuffer(nil)

		if err := json.NewEncoder(buff).Encode(&body); err != nil {
			t.Fatalf("failed encoding json: %v", err)
		}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/", buff)

		h.HandlePostShortenUrl(h.shortenerService).ServeHTTP(w, r)

		assertStatus(t, w.Code, http.StatusOK)
		assertHeader(t, w.Header().Get("Content-Type"), "application/json")

		var res map[string]interface{}
		if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
			t.Fatalf("failed decoding json: %v", err)
		}

		url, present := res["url"]
		if !present {
			t.Fatalf("response body does not contain key \"url\"")
		}

		want := "14q60P"
		if url != want {
			t.Fatalf("did not get correct shortened url, got %v, want %v", url, want)
		}
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status code, got %v, want %v", got, want)
	}
}

func assertHeader(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct header value, got %v, want %v", got, want)
	}
}
