package main

import (
	"yourate/handlers"
	"yourate/util"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	util.SetupLogger(e)

	handler := handlers.New()

	e.Static("/css", "css")
	e.Static("/assets", "assets")

	e.GET("/", handler.HandleHome)
	e.GET("/videos/dialog", handler.HandleDialogVideo)

	e.Logger.Fatal(e.Start(":1323"))
}
