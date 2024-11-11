package server_error

import (
	"log/slog"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, logger *slog.Logger, msg string, err error) {
	logger.Error(msg, slog.String("err", err.Error()))
	http.Error(w, "internal server error", http.StatusBadRequest)
}

func BadRequest(w http.ResponseWriter, logger *slog.Logger, msg string, err error) {
	logger.Error(msg, slog.String("err", err.Error()))
	http.Error(w, "internal server error", http.StatusBadRequest)
}
