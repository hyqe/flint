package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/hyqe/flint/internal/graceful"
	cli "github.com/urfave/cli/v2"
)

func main() {
	cliApp := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Value:   1389,
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Value:   false,
			},
		},
	}
	cliApp.Action = run
	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	handler := &Handler{
		Cache:   NewCache(),
		Verbose: c.Bool("verbose"),
	}
	return graceful.Run(graceful.NewServer(
		graceful.WithHandler(handler),
		graceful.WithPort(c.Int("port")),
	))
}

type Handler struct {
	Cache
	Verbose bool
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.Verbose {
		log.Println(r.Method, r.URL.Path)
	}
	switch r.Method {
	case http.MethodGet:
		if v, ok := h.Get(r.URL.Path); ok {
			content, ok := v.(Content)
			if !ok {
				Internal(w, "failed to read cache")
				return
			}
			w.Header().Set(HeaderContentType, content.Type)
			w.Write(content.Body)
		} else {
			NotFound(w)
		}
	case http.MethodPut:
		if r.Body == nil {
			Bad(w, "empty body")
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			Internal(w, fmt.Sprintf("failed to read body: %v", err))
			return
		}

		h.Put(r.URL.Path, Content{
			Body: body,
			Type: GetContentType(r),
		})
	case http.MethodDelete:
		if !h.Delete(r.URL.Path) {
			NotFound(w)
			return
		}
	default:
		NotFound(w)
	}
}

const (
	HeaderContentType = "Content-Type"
)

func Internal(w http.ResponseWriter, msg string) {
	http.Error(w, msg, http.StatusInternalServerError)
}
func Bad(w http.ResponseWriter, msg string) {
	http.Error(w, msg, http.StatusBadRequest)
}
func NotFound(w http.ResponseWriter) {
	http.Error(w, "", http.StatusNotFound)
}
func Ok(w http.ResponseWriter, v interface{}) {
	fmt.Fprint(w, v)
}

func GetContentType(r *http.Request) string {
	switch t := strings.TrimSpace(r.Header.Get(HeaderContentType)); t {
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

type Cache struct {
	sync.RWMutex
	db map[string]interface{}
}

func NewCache() Cache {
	return Cache{
		db: make(map[string]interface{}),
	}
}

func (c *Cache) Get(k string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	v, ok := c.db[k]
	return v, ok
}

func (c *Cache) Put(k string, v interface{}) {
	c.Lock()
	defer c.Unlock()
	c.db[k] = v
}

func (c *Cache) Delete(k string) bool {
	c.Lock()
	defer c.Unlock()
	_, ok := c.db[k]
	if !ok {
		return false
	}
	delete(c.db, k)
	return true
}
