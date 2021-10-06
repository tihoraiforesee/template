package pkg

type RepositoryInterface interface {
	GetBooks() string
}

type Repository struct {
}

func (r *Repository) GetBooks() string {
	return "all books"
}
