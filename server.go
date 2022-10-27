package traineeEVOFintech

import (
	"context"
	"net/http"
)

// Server struct is superstructure over http.Server
type Server struct {
	httpServer *http.Server
}

// Run is Server method, that launches the server on the specified port
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	return s.httpServer.ListenAndServe()
}

// Shutdown is Server method, that stops the server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
