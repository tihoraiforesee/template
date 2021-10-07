package pkg

import (
	"github.com/andreasaiforesee/template/models"
	"github.com/gofiber/fiber/v2"
)

type HandlerInterface interface {
	HandlerGetAllBook(ctx *fiber.Ctx) string
	HandlerGetBook(ctx *fiber.Ctx) models.Book
}

type Handler struct {
	Service ServiceInterface
}

func (h Handler) HandlerGetAllBook(ctx *fiber.Ctx) string {
	books := h.Service.GetBooks()
	return books
}

func (h Handler) HandlerGetBook(ctx *fiber.Ctx) models.Book {
	book := models.Book{}
	book.Name = "buku"
	return book
}
