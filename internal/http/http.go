package http

import (
	"net/http"

	"github.com/hashicorp/go-hclog"
	"github.com/labstack/echo"

	"github.com/netauth/netauth/pkg/netauth"
	// At least one token cache must be registered for netauth to
	// work correctly.
	_ "github.com/netauth/netauth/pkg/netauth/memory"
)

// Server serves the user interface over http using Echo.
type Server struct {
	hclog.Logger

	*echo.Echo

	nacl *netauth.Client
}

// New initializes and returns a new http.Server.
func New() (*Server, error) {
	s := Server{
		Logger: hclog.L().Named("http"),
		Echo:   echo.New(),
	}

	client, err := netauth.New()
	if err != nil {
		return nil, err
	}
	s.nacl = client

	r, err := newRenderer("tpl", s.Logger)
	if err != nil {
		return nil, err
	}
	s.Renderer = r

	s.GET("/meta/ok", s.metaOK)
	s.GET("/meta/about", s.metaAbout)

	s.GET("/system/status", s.systemStatus)

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
