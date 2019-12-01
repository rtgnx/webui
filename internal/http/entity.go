package http

import (
	"net/http"

	"github.com/labstack/echo"
)

func (s *Server) entityInfo(c echo.Context) error {
	entity, err := s.nacl.EntityInfo(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.Render(http.StatusInternalServerError, "internal-error", err)
	}

	keys, err := s.nacl.EntityKeys(c.Request().Context(), c.Param("id"), "READ", "*", "")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "internal-error", err)
	}

	kv, err := s.nacl.EntityUM(c.Request().Context(), c.Param("id"), "READ", "", "")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "internal-error", err)
	}

	d := make(map[string]interface{})
	d["entity"] = entity
	d["keys"] = keys
	d["kv"] = kv

	return c.Render(http.StatusOK, "entity-info", d)
}
