package models

type Book struct {
	ID   int64
	Name string
}

type ServiceInterface interface {
	GetBooks() string
}

type HandlerInterface interface {
}
