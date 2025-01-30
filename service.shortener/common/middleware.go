package common

import (
	"log"
	"net/http"
	"time"
)

// Takes a handler and applies all of the given middleware funcs
func ApplyMiddleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *responseRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rec, r)

		log.Printf("%s %s %d - %v",
			r.Method,
			r.URL.Path,
			rec.statusCode,
			time.Since(start),
		)
	})
}
