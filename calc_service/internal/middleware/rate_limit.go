package middleware

import (
	"net"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

type rateLimiterStore struct {
	limiters map[string]*rate.Limiter
	mu       sync.Mutex
}

func newRateLimiterStore() *rateLimiterStore {
	return &rateLimiterStore{
		limiters: make(map[string]*rate.Limiter),
	}
}

func getClientIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}
func RateLimit(requestPerSecond float64, burst int) func(http.Handler) http.Handler {
	store := newRateLimiterStore()
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := getClientIP(r)

			store.mu.Lock()
			limiter, exists := store.limiters[ip]
			if !exists {
				limiter = rate.NewLimiter(rate.Limit(requestPerSecond), burst)
				store.limiters[ip] = limiter
			}
			store.mu.Unlock()

			if !limiter.Allow() {
				http.Error(w, "rate Limit exceeded", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
