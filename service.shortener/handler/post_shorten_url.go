package handler

import (
	"encoding/json"
	"net/http"
)

type postShortenUrlRequest struct {
	Url string `json:"url"`
}

func handlePostShortenUrl() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req postShortenUrlRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"url": "14q60P"})
	})
}
