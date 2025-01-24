package handler

import (
	"net/http"

	"github.com/vhall1/shorturl/lib/util"
	"github.com/vhall1/shorturl/services/snowflake/types"
)

type SnowflakeHttpHandler struct {
	snowflakeService types.Snowflake
}

func NewSnowflakeHttpHandler(snowflakeService types.Snowflake) *SnowflakeHttpHandler {
	return &SnowflakeHttpHandler{snowflakeService: snowflakeService}
}

func (h *SnowflakeHttpHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("GET /", h.HandleGetSnowflake(h.snowflakeService))
}

func (*SnowflakeHttpHandler) HandleGetSnowflake(snowflakeService types.Snowflake) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := snowflakeService.Generate()
		util.WriteJSON(w, http.StatusOK, &map[string]interface{}{"id": id})
	})
}
