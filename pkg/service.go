package pkg

type ServiceInterface interface {
	GetBooks() string
}

type Service struct {
	repo RepositoryInterface
}

func (s *Service) GetBooks() string {
	books := s.repo.GetBooks()
	return books
}
