package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func Logger(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			wrapper := &WrapperWriter{
				ResponseWriter: w,
				StatusCode:     http.StatusOK,
			}
			next.ServeHTTP(wrapper, r)
			end := time.Since(start)
			logger.Debug(r.URL.String(),
				slog.Int("status", wrapper.StatusCode),
				slog.String("time", end.String()),
				slog.String("method", r.Method),
			)
		})
	}
}
