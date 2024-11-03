package handlers

import (
	"yourate/views"

	"github.com/labstack/echo/v4"
)

func (h *Handler) HandleHome(c echo.Context) error {
	return views.Home().Render(c.Request().Context(), c.Response().Writer)
}
