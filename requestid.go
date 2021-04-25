package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxRequest int

const (
	RequestIDHeader            = "X-Request-Id"
	RequestIDKey    ctxRequest = iota
)

func Request(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get(RequestIDHeader)
		if requestID == "" {
			requestID = uuid.NewString()
		}
		ctx = context.WithValue(ctx, RequestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func GetRequestID(ctx context.Context) string {
	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
		return reqID
	}
	return ""
}
