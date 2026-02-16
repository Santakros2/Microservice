package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWT(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			auth := r.Header.Get("Authorization")

			if auth == "" {
				http.Error(w, "missing token", 401)
				return
			}

			tokenStr := strings.TrimPrefix(auth, "Bearer ")

			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "invalid token", 401)
				return
			}

			next.ServeHTTP(w, r)

		})
	}
}
