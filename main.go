package main

import (
	"yourate/handlers"
	"yourate/services"
	"yourate/util"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	util.SetupLogger(e)

	s := services.New()

	handler := handlers.New(s)

	e.Static("/assets", "assets")

	e.GET("/", handler.HomeHandler.GetHome)
	e.GET("/videos/dialog", handler.VideoHandler.GetVideoLogDialog)

	e.Logger.Fatal(e.Start(":1323"))
}
