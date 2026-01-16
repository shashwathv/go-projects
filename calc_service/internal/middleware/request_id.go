package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type RequestIDType struct{}

var RequestIDKey = RequestIDType{}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.NewString()

		ctx := context.WithValue(r.Context(), RequestIDKey, requestID)

		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(RequestIDKey).(string); ok {
		return id
	}
	return ""
}
