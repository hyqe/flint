package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	_ "embed"

	"github.com/hyqe/graceful"
	cli "github.com/urfave/cli/v2"
)

func main() {
	cliApp := &cli.App{
		Name:        "Flint",
		Description: "\n" + Docs,
		Copyright:   "\n" + License,
		Version:     Version,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				EnvVars: []string{"FLINT_PORT"},
				Value:   2000,
			},
			&cli.BoolFlag{
				Name:    "verbose",
				EnvVars: []string{"FLINT_VERBOSE"},
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
	verbose := c.Bool("verbose")
	port := c.Int("port")

	handler := &Handler{
		Cache:   NewCache(),
		Verbose: verbose,
	}

	server := graceful.NewServer(
		graceful.WithHandler(handler),
		graceful.WithPort(port),
	)

	if verbose {
		log.Printf("starting Flint on port: %v\n", port)
	}
	return graceful.Run(server)
}

var (
	//go:embed LICENSE
	License string

	//go:embed VERSION
	Version string

	//go:embed docs.md
	Docs string
)

type Handler struct {
	Cache
	Verbose bool
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body []byte
	var err error
	switch r.Method {
	case http.MethodGet:
		if v, ok := h.Get(r.URL.Path); ok {
			content, ok := v.(Content)
			if !ok {
				Internal(w, "failed to read cache")
				break
			}
			w.Header().Set(HeaderContentType, content.Type)
			w.Write(content.Body)
		} else {
			NotFound(w)
		}
	case http.MethodPut:
		if r.Body == nil {
			Bad(w, "empty body")
			break
		}

		body, err = io.ReadAll(r.Body)
		if err != nil {
			Internal(w, fmt.Sprintf("failed to read body: %v", err))
			break
		}

		h.Put(r.URL.Path, Content{
			Body: body,
			Type: GetContentType(r),
		})
	case http.MethodDelete:
		if !h.Delete(r.URL.Path) {
			NotFound(w)
		}
	default:
		NotFound(w)
	}
	if h.Verbose {
		log.Println(r.Method, r.URL.Path, string(body))
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
