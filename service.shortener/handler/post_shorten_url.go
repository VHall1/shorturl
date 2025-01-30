package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/vhall1/shorturl/service.shortener/domain"
)

type postShortenUrlRequest struct {
	Url string `json:"url"`
}

func handlePostShortenUrl(urlService *domain.UrlService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req postShortenUrlRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		shortUrl, err := urlService.ShortenUrl(ctx, req.Url)
		if err != nil {
			log.Printf("got an unexpected error shortening url: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"url": shortUrl})
	})
}
