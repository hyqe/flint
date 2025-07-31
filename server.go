package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os/signal"
	"sync"
)

type serverConfig struct {
	Port string `env:"PORT" envDefault:"80"`
}

func runServer(ctx context.Context, c serverConfig, handler http.Handler) error {
	ctx, cancel := signal.NotifyContext(ctx)
	defer cancel()
	server := http.Server{
		Addr:    net.JoinHostPort("", c.Port),
		Handler: LogRequests(handler),
	}
	var wg sync.WaitGroup
	var errShutdown error
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			errShutdown = err
		}
	}()
	slog.Info("server listening", "address", server.Addr)
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to list and serve: %w", err)
	}
	cancel()
	wg.Wait()
	return errShutdown
}
