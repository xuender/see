package app

type Service struct{}

func NewService() *Service { return &Service{} }

func (p *Service) About() string {
	return "about"
}
