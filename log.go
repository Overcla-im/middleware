package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Log ...
func Log(next http.Handler) http.Handler {
	log.WithField("middleware", "log").Print("installing")
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s", r.URL)
		next.ServeHTTP(w, r)

	}

	return http.HandlerFunc(fn)
}
