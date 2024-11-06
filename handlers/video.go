package handlers

import (
	"yourate/components"
	"yourate/services"

	"github.com/labstack/echo/v4"
)

type (
	VideoHandler interface {
		GetVideoLogDialog(c echo.Context) error
	}

	videoHandler struct {
		services.VideoService
	}
)

func (h *videoHandler) GetVideoLogDialog(c echo.Context) error {
	url := c.FormValue("url")

	video, err := h.VideoService.GetVideoDetails(url)

	return components.VideoLogDialog(video, err).Render(c.Request().Context(), c.Response().Writer)
}
