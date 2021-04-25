package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// RecoverMiddleware ...
func RecoverMiddleware(next http.Handler) http.Handler {
	log.WithField("middleware", "recover").Println("installing")
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %#v %+v", next, err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
