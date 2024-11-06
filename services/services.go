package services

import "github.com/labstack/echo/v4"

type Services struct {
	Video VideoService
}

func New(l echo.Logger) *Services {
	return &Services{
		Video: &videoService{logger: l},
	}
}
