package handler

import (
	"encoding/json"
	"net/http"

	"github.com/vhall1/shorturl/service.snowflake/domain"
)

func handleGetSnowflake(snowflakeService *domain.SnowflakeService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := snowflakeService.Generate()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int64{"id": id})
	})
}
