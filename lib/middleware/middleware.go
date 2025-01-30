package middleware

import "net/http"

// Takes a handler and applies all of the given middleware funcs
func ApplyMiddleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}
