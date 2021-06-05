package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/hyqe/flint/internal/storage"
)

type Flint struct {
	storage.Storage
	Verbose bool
}

func (f *Flint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		f.put(w, r)
	case http.MethodGet:
		f.get(w, r)
	case http.MethodDelete:
		f.delete(w, r)
	default:
		http.Error(w, "", http.StatusNotFound)
	}
}

func (f *Flint) put(w http.ResponseWriter, r *http.Request) {
	contentType := GetContentType(r)

	if f.Verbose {
		log.Println("DEBUG:", r.Method, r.URL.Path, contentType)
	}

	if r.Body == nil {
		httpErr(w, "empty body", http.StatusBadRequest, f.Verbose)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		httpErr(w, fmt.Sprintf("failed to read body: %v", err), http.StatusInternalServerError, f.Verbose)
		return
	}

	err = f.Storage.Put(r.URL.Path, Content{Body: body, Type: contentType})
	if err != nil {
		httpErr(w, fmt.Sprintf("failed to write to storage: %v", err), storage.HttpStatus(err), f.Verbose)
	}
}

func (f *Flint) get(w http.ResponseWriter, r *http.Request) {
	if f.Verbose {
		log.Println("DEBUG:", r.Method, r.URL.Path)
	}
	var content Content
	if err := f.Storage.Get(r.URL.Path, &content); err != nil {
		httpErr(w, fmt.Sprintf("failed to read from storage: %v", err), storage.HttpStatus(err), f.Verbose)
	}
	w.Header().Set(headerContentType, content.Type)
	w.Write(content.Body)
}

func (f *Flint) delete(w http.ResponseWriter, r *http.Request) {
	if f.Verbose {
		log.Println("DEBUG:", r.Method, r.URL.Path)
	}
	if err := f.Storage.Delete(r.URL.Path); err != nil {
		httpErr(w, fmt.Sprintf("failed to delete from storage: %v", err), storage.HttpStatus(err), f.Verbose)
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

func httpErr(w http.ResponseWriter, err string, code int, verbose bool) {
	if verbose {
		switch code {
		case http.StatusNotFound:
			log.Println("DEBUG:", code, err)
		default:
			log.Println("ERROR:", code, err)
		}
	}
	http.Error(w, err, code)
}
