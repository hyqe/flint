package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	slog.Info("starting main")
	ctx := context.Background()
	mux := defaultRoutes()
	if err := routes(mux); err != nil {
		slog.Error("failed to build routes", "reason", err)
		os.Exit(1)
	}
	if err := run(ctx, mux); err != nil {
		slog.Error("failed to run server", "reason", err)
		os.Exit(1)
	}
}
