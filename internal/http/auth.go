package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (s *Server) loginForm(c echo.Context) error {
	return c.Render(http.StatusOK, "auth-login", nil)
}

func (s *Server) loginGetToken(c echo.Context) error {
	token, err := s.nacl.AuthGetToken(
		c.Request().Context(),
		c.FormValue("id"),
		c.FormValue("secret"),
	)
	if err != nil {
		return c.Render(http.StatusUnauthorized, "internal-error", err)
	}
	cookie := new(http.Cookie)
	cookie.Name = "netauth"
	cookie.Value = token
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(s.sessionTimeout)
	c.SetCookie(cookie)

	next := c.QueryParam("next")
	s.Trace("Next location from login", "location", next)
	if next == "" {
		next = "/entity/info/" + c.FormValue("id")
	}

	return c.Redirect(http.StatusSeeOther, next)
}

func (s *Server) authLogout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "netauth"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(time.Hour * 24 * -2)
	cookie.Path = "/"
	c.SetCookie(cookie)
	return c.Redirect(http.StatusSeeOther, "/")
}
