package pkg

import "github.com/gofiber/fiber/v2"

type HandlerInterface interface {
	HandlerGetAllBook(ctx *fiber.Ctx) string
}

type Handler struct {
	Service ServiceInterface
}

func (h Handler) HandlerGetAllBook(ctx *fiber.Ctx) string {
	books := h.Service.GetBooks()
	return books
}
