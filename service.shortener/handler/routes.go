package handler

import (
	"net/http"

	"github.com/vhall1/shorturl/service.shortener/domain"
)

func RegisterRoutes(router *http.ServeMux, urlService *domain.UrlService) {
	router.Handle("POST /", handlePostShortenUrl(urlService))
	router.Handle("GET /s/{shortUrl}", handleGetRedirectUrl(urlService))
}
