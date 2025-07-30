package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"sync"
)

func run(ctx context.Context, handler http.Handler) error {
	ctx, cancel := signal.NotifyContext(ctx)
	defer cancel()
	server := http.Server{
		Addr:    net.JoinHostPort("", "8080"),
		Handler: handler,
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
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to list and serve: %w", err)
	}
	cancel()
	wg.Wait()
	return errShutdown
}
