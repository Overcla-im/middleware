package middleware

import (
	"net/http"
)

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
