package http

import (
	"net/http"

	"github.com/hashicorp/go-hclog"
	"github.com/labstack/echo"
)

// Server serves the user interface over http using Echo.
type Server struct {
	hclog.Logger

	*echo.Echo
}

// New initializes and returns a new http.Server.
func New() (*Server, error) {
	s := Server{
		hclog.L().Named("http"),
		echo.New(),
	}

	r, err := newRenderer("tpl", s.Logger)
	if err != nil {
		return nil, err
	}
	s.Renderer = r

	s.GET("/meta/ok", s.metaOK)
	s.GET("/meta/about", s.metaAbout)

	return &s, nil
}

// Serve commences serving and blocks forever.
func (s *Server) Serve(bind string) error {
	return s.Start(bind)
}

func (s *Server) metaOK(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func (s *Server) metaAbout(c echo.Context) error {
	return c.Render(http.StatusOK, "meta-about", "foo")
}
