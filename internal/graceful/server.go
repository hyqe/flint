package graceful

import (
	"context"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	port        int
	srv         *http.Server
	shutTimeout time.Duration
	handler     http.Handler
}

func NewServer(opts ...Option) *Server {
	s := &Server{
		port:        8080,
		shutTimeout: time.Second * 3,
	}
	s.Apply(opts...)
	return s
}

func (s *Server) Apply(opts ...Option) {
	for _, opt := range opts {
		opt(s)
	}
}

func (s *Server) Start() error {
	s.srv = &http.Server{
		Addr:    addr(s.port),
		Handler: s.handler,
	}
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutTimeout)
	defer cancel()
	return s.srv.Shutdown(ctx)
}

func addr(port int) string {
	return ":" + strconv.FormatInt(int64(port), 10)
}

type Option func(*Server)

func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func WithHandler(handler http.Handler) Option {
	return func(s *Server) {
		s.handler = handler
	}
}
