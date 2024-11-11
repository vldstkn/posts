package main

import (
	"log/slog"
	"net/http"
	"os"
	"p1/internal/app"
	"p1/internal/config"
	"p1/pkg/logger"
)

func main() {
	opts := logger.PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	logHandler := logger.NewPrettyHandler(os.Stdout, opts)
	log := slog.New(logHandler)

	conf := config.Load()

	app := app.App(log, conf)

	server := http.Server{
		Addr:    ":6543",
		Handler: app,
	}

	log.Info("Server run", slog.String("Address", conf.Addr), slog.String("Env", conf.Env))
	server.ListenAndServe()
}
