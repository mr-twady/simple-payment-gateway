package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sony/gobreaker"
)

// CircuitBreakerMiddleware wraps my HTTP handlers with circuit breaker logic i.e to support resilience for gateway connections and possible issues
func CircuitBreakerMiddleware(cb *gobreaker.CircuitBreaker) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := cb.Execute(func() (interface{}, error) {
				next.ServeHTTP(w, r) // Call the next handler
				return nil, nil
			})
			if err != nil {
				http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
			}
		})
	}
}
