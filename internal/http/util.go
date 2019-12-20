package http

import (
	"context"
	"errors"

	"github.com/labstack/echo"
	"github.com/netauth/netauth/pkg/netauth"
)

func authorizedContext(c echo.Context) (context.Context, error) {
	t, ok := c.Get("token").(string)
	if !ok {
		return nil, errors.New("No valid token was present in the request")
	}

	ctx := netauth.Authorize(c.Request().Context(), t)
	return ctx, nil
}
