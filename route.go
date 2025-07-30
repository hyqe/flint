package main

import (
	"net/http"

	"github.com/hyqe/flint/internal/handlers"
	"github.com/hyqe/flint/internal/storage/memory"
)

func defaultRoutes() *http.ServeMux {
	return http.NewServeMux()
}

func routes(m *http.ServeMux) error {
	lookup := memory.NewLookup[*handlers.Value]()
	handlers.HandlePut(m, lookup)
	handlers.HandleGet(m, lookup)
	handlers.HandleDelete(m, lookup)
	return nil
}
