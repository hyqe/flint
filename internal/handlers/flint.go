package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/hyqe/flint/internal/cache"
)

type Flint struct {
	cache.Cacher
	Verbose bool
}

func (h *Flint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body []byte
	var err error
	switch r.Method {
	case http.MethodGet:
		if v, ok := h.Get(r.URL.Path); ok {
			content, ok := v.(Content)
			if !ok {
				http.Error(w, "failed to read cache", http.StatusInternalServerError)
				break
			}
			w.Header().Set(headerContentType, content.Type)
			w.Write(content.Body)
		} else {
			http.Error(w, "", http.StatusNotFound)
		}
	case http.MethodPut:
		if r.Body == nil {
			http.Error(w, "empty body", http.StatusBadRequest)
			break
		}
		body, err = io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to read body: %v", err), http.StatusInternalServerError)
			break
		}
		h.Put(r.URL.Path, Content{
			Body: body,
			Type: GetContentType(r),
		})
	case http.MethodDelete:
		if !h.Delete(r.URL.Path) {
			http.Error(w, "", http.StatusNotFound)
		}
	default:
		http.Error(w, "", http.StatusNotFound)
	}
	if h.Verbose {
		log.Println(r.Method, r.URL.Path, string(body))
	}
}

const (
	headerContentType = "Content-Type"
)

func GetContentType(r *http.Request) string {
	switch t := strings.TrimSpace(r.Header.Get(headerContentType)); t {
	case "":
		return "plain/text"
	default:
		return t
	}
}

type Content struct {
	Type string
	Body []byte
}
