package http

import (
	"github.com/labstack/echo"
)

func (s *Server) parseToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("netauth")
		if err != nil {
			s.Trace("No cookie in request")
			return next(c)
		}

		claims, err := s.nacl.Validate(cookie.Value)
		if err != nil {
			s.Trace("Cookie present but invalid", "error", err)
			return next(c)
		}

		c.Set("claims", claims)
		c.Set("token", cookie.Value)
		s.Logger.Debug("Set cookie and claims for request")
		return next(c)
	}
}
