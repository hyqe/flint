package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	slog.Info("starting Flint")
	ctx := context.Background()
	cfg := newConfig()
	if err := parseConfig(&cfg); err != nil {
		slog.Error("failed to parse config", "reason", err)
		os.Exit(1)
	}
	mux := defaultRoutes()
	if err := routes(mux); err != nil {
		slog.Error("failed to build routes", "reason", err)
		os.Exit(1)
	}
	if err := runServer(ctx, cfg.Server, mux); err != nil {
		slog.Error("failed to run server", "reason", err)
		os.Exit(1)
	}
}
