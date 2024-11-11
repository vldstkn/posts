package middleware

import (
	"context"
	"net/http"
	"p1/pkg/jwt"
	"strings"
)

func IsAuthor(secret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authedHeader := r.Header.Get("Authorization")
			if authedHeader == "" || !strings.HasPrefix(authedHeader, "Bearer ") {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
			token := strings.TrimPrefix(authedHeader, "Bearer ")
			isValid, data := jwt.NewJWT(secret).Parse(token)
			if !isValid {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
			if data.Role != "author" {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
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
