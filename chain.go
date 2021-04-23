package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func init() {
	// log "github.com/sirupsen/logrus"
	/* log.SetFormatter(&log.JSONFormatter{}) */
}

// LogMiddleware ...
func LogMiddleware(next http.Handler) http.Handler {
	log.WithField("middleware", "log").Print("installing")
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s", r.URL)
		next.ServeHTTP(w, r)

	}

	return http.HandlerFunc(fn)
}

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

// CommonHandlers ...
type CommonHandlers struct {
	handlers []func(http.Handler) http.Handler
}

// Then ....
func (ch *CommonHandlers) Then(h http.Handler) http.Handler {
	if h == nil {
		h = http.DefaultServeMux
	}

	for i := range ch.handlers {
		h = ch.handlers[len(ch.handlers)-1-i](h)
	}
	return h
}

// ThenFunc ....
func (ch *CommonHandlers) ThenFunc(h http.HandlerFunc) http.Handler {
	if h == nil {
		return ch.Then(nil)
	}
	return ch.Then(h)

}

// NewCommonHandlers ...
func NewCommonHandlers(handlers ...func(http.Handler) http.Handler) *CommonHandlers {
	return &CommonHandlers{handlers: handlers}
}
