package http

import (
	"net/http"

	"github.com/labstack/echo"
)

// Server serves the user interface over http using Echo.
type Server struct {
	*echo.Echo
}

// New initializes and returns a new http.Server.
func New() (*Server, error) {
	s := Server{
		echo.New(),
	}

	s.GET("/meta/ok", s.metaOK)

	return &s, nil
}

// Serve commences serving and blocks forever.
func (s *Server) Serve(bind string) error {
	return s.Start(bind)
}

func (s *Server) metaOK(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
