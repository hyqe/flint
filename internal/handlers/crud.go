package handlers

import (
	"io"
	"net/http"

	"github.com/hyqe/flint/internal/storage/memory"
)

type Value struct {
	ContentType string
	Body        []byte
}

func HandleGet(m *http.ServeMux, lookup *memory.Lookup[*Value]) {
	m.HandleFunc("GET /{key}", func(w http.ResponseWriter, r *http.Request) {
		key := r.PathValue("key")
		v, ok := lookup.Get(key)
		if !ok {
			http.Error(w, "", http.StatusNotFound)
			return
		}
		if v.ContentType != "" {
			w.Header().Set("Content-Type", v.ContentType)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(v.Body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

const maxPutReadBytes = 1_000_000_000

func HandlePut(m *http.ServeMux, lookup *memory.Lookup[*Value]) {
	m.HandleFunc("PUT /{key}", func(w http.ResponseWriter, r *http.Request) {
		key := r.PathValue("key")
		if r.Body == nil {
			http.Error(w, "expected request body", http.StatusBadRequest)
			return
		}
		b, err := io.ReadAll(io.LimitReader(r.Body, maxPutReadBytes))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		value := Value{
			ContentType: w.Header().Get("Content-Type"),
			Body:        b,
		}
		lookup.Put(key, &value)
		w.WriteHeader(http.StatusCreated)
	})
}

func HandleDelete(m *http.ServeMux, lookup *memory.Lookup[*Value]) {
	m.HandleFunc("DELETE /{key}", func(w http.ResponseWriter, r *http.Request) {
		key := r.PathValue("key")
		lookup.Delete(key)
		w.WriteHeader(http.StatusOK)
	})
}
