package handler

import "net/http"

func RegisterRoutes(router *http.ServeMux) {
	router.Handle("GET /hello", handleGetHello())
}
