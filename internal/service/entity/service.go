package entity

type Service struct{}

func NewService() *Service {

	return &Service{}
}

func (s *Service) List() []Entity {
	return allEntities
}
