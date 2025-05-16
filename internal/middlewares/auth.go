package middlewares

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement auth middleware
		// 1. Get the token from the request
		// 2. Verify the token
		// 3. verify user has permission to access the resource

		next.ServeHTTP(w, r)
	})
}
