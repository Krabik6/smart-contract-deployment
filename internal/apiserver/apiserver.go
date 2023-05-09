package apiserver

import (
	"context"
	"net/http"
	"time"
)

// Server ...
type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{}
}

// Run ...
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		MaxHeaderBytes:    1 << 20,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

// ShutDown ...
func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)

}
