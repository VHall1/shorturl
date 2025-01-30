package handler

import (
	"encoding/json"
	"net/http"

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

		// ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		// defer cancel()

		// shortUrl, err := urlService.ShortenUrl(ctx, req.Url)
		// if err != nil {
		// 	log.Printf("got an unexpected error shortening url: %v", err)
		// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
		// 	return
		// }

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"url": "test.com"})
	})
}
