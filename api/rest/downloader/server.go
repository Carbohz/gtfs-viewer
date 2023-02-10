package downloader

import (
	"context"
	"fmt"
	"net/http"
)

type APIServer struct {
	httpServer *http.Server
}

func NewAPIServer() (*APIServer, error) {
	r := setupRouter()

	srv := &APIServer{
		httpServer: &http.Server{
			Addr: "localhost:8080",
			Handler: r,
		},
	}
	return srv, nil
}

func (s *APIServer) Run(ctx context.Context) error {

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("listen: %s", err)
	}

	return nil
}