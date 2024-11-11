package middleware

import (
	"context"
	"net/http"
	"p1/internal/config"
	"p1/pkg/jwt"
	"strings"
)

func IsAuthed(config *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authedHeader := r.Header.Get("Authorization")
			if authedHeader == "" || !strings.HasPrefix(authedHeader, "Bearer ") {
				writeUnauthed(w)
				return
			}
			token := strings.TrimPrefix(authedHeader, "Bearer ")
			isValid, data := jwt.NewJWT(config.JWTSecret).Parse(token)
			if !isValid {
				writeUnauthed(w)
				return
			}
			ctx := context.WithValue(r.Context(), "authInfo", AuthInfo{
				Id:   data.Id,
				Role: data.Role,
			})
			req := r.WithContext(ctx)
			next.ServeHTTP(w, req)
		})
	}
}

func writeUnauthed(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}
