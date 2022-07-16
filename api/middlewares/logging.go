package api

import (
	"net/http"

	"github.com/harvestcore/upgote/log"
)

// LoggingMiddleware Logs each request.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.AddRequest(r)
		next.ServeHTTP(w, r)
	})
}
