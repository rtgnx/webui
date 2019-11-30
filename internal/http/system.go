package http

import (
	"net/http"

	"github.com/labstack/echo"
)

func (s *Server) systemStatus(c echo.Context) error {
	status, err := s.nacl.SystemStatus(c.Request().Context())
	if err != nil {
		return c.Render(http.StatusInternalServerError, "internal-error", err)
	}

	return c.Render(http.StatusOK, "system-status", status)
}
