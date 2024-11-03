package handlers

import (
	"fmt"
	"yourate/components"
	"yourate/services"

	"github.com/labstack/echo/v4"
)

func (h *Handler) HandleDialogVideo(c echo.Context) error {
	url := c.FormValue("url")

	if url == "" {
		return components.VideoLogDialog(nil).Render(c.Request().Context(), c.Response().Writer)
	}

	video, err := services.GetVideoDetails(url)

	if err != nil {
		return err
	}

	fmt.Println(video)

	return components.VideoLogDialog(video).Render(c.Request().Context(), c.Response().Writer)
}
