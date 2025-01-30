package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/vhall1/shorturl/service.shortener/domain"
)

func handleGetRedirectUrl(urlService *domain.UrlService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		shortUrl := r.PathValue("shortUrl")

		// unlikely to happen, but good to validate regardless
		if shortUrl == "" {
			http.Error(w, "Invalid short URL", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		longUrl, err := urlService.GetRedirectUrl(ctx, shortUrl)
		if err != nil {
			log.Printf("got an unexpected error getting redirect url: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"url": longUrl})
	})
}
