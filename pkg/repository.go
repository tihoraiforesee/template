package pkg

import (
	"fmt"

	"github.com/andreasaiforesee/template/models"
)

type RepositoryInterface interface {
	GetBooks() string
}

type Repository struct {
}

func (r *Repository) GetBooks() string {
	books := models.Book{}
	books.Name = "testing"
	fmt.Println(books)
	return "all books"
}
