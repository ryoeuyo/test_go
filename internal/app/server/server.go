package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"ryoeuyo/testgo/internal/domain/models"
	"time"
)

type Server struct {
	srv *http.Server
}

type ServerOptions struct {
	Timeout time.Duration
	Port    int
}

func New(opts *ServerOptions, endpoints ...models.EndPoint) *Server {
	mux := http.NewServeMux()

	for _, ep := range endpoints {
		mux.HandleFunc(ep.Target, ep.Handler)
	}

	return &Server{
		srv: &http.Server{
			Handler:      mux,
			ReadTimeout:  opts.Timeout,
			WriteTimeout: opts.Timeout,
			Addr:         fmt.Sprintf(":%d", opts.Port),
		},
	}
}

func (s *Server) MustStartServer() {
	if err := s.srv.ListenAndServe(); err != nil {
		log.Fatalf("couldn't start the server: %s", err.Error())
	}
}

func (s *Server) GracefulStop() error {
	return s.srv.Shutdown(context.Background())
}
