package services

type Services struct {
	Video VideoService
}

func New() *Services {
	return &Services{
		Video: &videoService{},
	}
}
