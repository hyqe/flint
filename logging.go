package main

import (
	"log/slog"
	"net/http"
	"os"
)

func init() {
	level := slog.LevelInfo
	// LEVEL is the only environment variable not controled by the config, as it
	// may need to set before the config is parsed.
	if v, ok := os.LookupEnv("LEVEL"); ok {
		if err := level.UnmarshalText([]byte(v)); err != nil {
			slog.Error("failed to set logging level", "invalid", v, "reason", err)
		}
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})))
}

func LogRequests(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Debug("request recieved",
			"method", r.Method,
			"path", r.URL.Path,
		)
		handler.ServeHTTP(w, r)
	}
}
