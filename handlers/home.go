package handlers

import (
	"yourate/views"

	"github.com/labstack/echo/v4"
)

type (
	HomeHandler interface {
		GetHome(c echo.Context) error
	}

	homeHandler struct{}
)

func (h *homeHandler) GetHome(c echo.Context) error {
	return views.Home().Render(c.Request().Context(), c.Response().Writer)
}
