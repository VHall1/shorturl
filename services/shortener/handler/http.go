package handler

import (
	"net/http"

	"github.com/vhall1/shorturl/lib/util"
	"github.com/vhall1/shorturl/services/shortener/types"
)

type ShortenerHttpHandler struct {
	shortenerService types.ShortenerService
}

func NewShortenerHttpHandler(shortenerService types.ShortenerService) *ShortenerHttpHandler {
	return &ShortenerHttpHandler{shortenerService: shortenerService}
}

func (h *ShortenerHttpHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("GET /s/{shortUrl}", h.HandleGetLongUrl(h.shortenerService))
	mux.Handle("POST /", h.HandlePostShortenUrl(h.shortenerService))
}

func (*ShortenerHttpHandler) HandleGetLongUrl(shortenerService types.ShortenerService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		shortUrl := r.PathValue("shortUrl")

		longUrl, err := shortenerService.GetRedirectUrl(r.Context(), shortUrl)
		if err != nil {
			util.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		util.WriteJSON(w, http.StatusOK, &map[string]interface{}{"url": longUrl})
	})
}

func (*ShortenerHttpHandler) HandlePostShortenUrl(shortenerService types.ShortenerService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req map[string]string
		if err := util.ParseJSON(r, &req); err != nil {
			util.WriteError(w, http.StatusBadRequest, err)
			return
		}

		shortUrl, err := shortenerService.ShortenUrl(r.Context(), req["url"])
		if err != nil {
			util.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		util.WriteJSON(w, http.StatusOK, &map[string]interface{}{"url": shortUrl})
	})
}
