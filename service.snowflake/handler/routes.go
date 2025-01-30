package handler

import (
	"net/http"

	"github.com/vhall1/shorturl/service.snowflake/domain"
)

func RegisterRoutes(router *http.ServeMux, snowflakeService *domain.SnowflakeService) {
	router.Handle("GET /", handleGetSnowflake(snowflakeService))
}
