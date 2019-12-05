package http

import (
	"net/http"

	"github.com/labstack/echo"
)

func (s *Server) groupInfo(c echo.Context) error {
	group, manages, err := s.nacl.GroupInfo(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.Render(http.StatusInternalServerError, "internal-error", err)
	}

	kv, err := s.nacl.GroupUM(c.Request().Context(), c.Param("id"), "READ", "*", "")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "internal-error", err)
	}

	d := echo.Map{}
	d["group"] = group
	d["manages"] = manages
	d["kv"] = kv

	return c.Render(http.StatusOK, "group-info", d)
}

func (s *Server) groupMembers(c echo.Context) error {
	members, err := s.nacl.GroupMembers(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.Render(http.StatusInternalServerError, "internal-error", err)
	}

	d := echo.Map{}
	d["members"] = members
	d["name"] = c.Param("id")

	return c.Render(http.StatusOK, "group-members", d)
}
