package handlers

import (
	"yourate/services"
)

type Handlers struct {
	VideoHandler
	HomeHandler
}

func New(s *services.Services) *Handlers {
	return &Handlers{
		VideoHandler: &videoHandler{s.Video},
		HomeHandler:  &homeHandler{},
	}
}
