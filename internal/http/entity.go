package http

import (
	"net/http"
	"strconv"

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

	kv, err := s.nacl.EntityUM(c.Request().Context(), c.Param("id"), "READ", "*", "")
	if err != nil {
		return c.Render(http.StatusInternalServerError, "internal-error", err)
	}

	groups, err := s.nacl.EntityGroups(c.Request().Context(), c.Param("id"))
	if err != nil {
		return c.Render(http.StatusInternalServerError, "internal-error", err)
	}

	d := echo.Map{}
	d["entity"] = entity
	d["keys"] = keys
	d["kv"] = kv
	d["groups"] = groups

	return c.Render(http.StatusOK, "entity-info", d)
}

func (s *Server) entitySearch(c echo.Context) error {
	d := echo.Map{}
	query := c.QueryParam("query")

	if query != "" {
		res, err := s.nacl.EntitySearch(c.Request().Context(), query)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "internal-error", err)
		}
		d["result"] = res
	}
	d["query"] = query
	d["name"] = "Entity"

	return c.Render(http.StatusOK, "search", d)
}

func (s *Server) entityCreateForm(c echo.Context) error {
	return c.Render(http.StatusOK, "entity-create", nil)
}

func (s *Server) entityCreateHandler(c echo.Context) error {
	ctx, err := authorizedContext(c)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/auth/login?next="+c.Path())
	}

	id := c.FormValue("id")
	secret := c.FormValue("secret")

	number, err := strconv.Atoi(c.FormValue("number"))
	if err != nil {
		// For bad input let NetAuth choose the next number.
		number = -1
	}

	if err := s.nacl.EntityCreate(ctx, id, secret, number); err != nil {
		s.Warn("Error creating entity", "error", err)
		return c.Render(http.StatusUnprocessableEntity, "internal-error", err)
	}

	s.Trace("Entity creation with values:",
		"entity", id)
	return c.Redirect(http.StatusSeeOther, "/entity/info/"+id)
}
