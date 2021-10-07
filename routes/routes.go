package routes

import (
	handler "github.com/andreasaiforesee/template/pkg"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	handler handler.HandlerInterface
}

func (r *Routes) SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/books", func(c *fiber.Ctx) error {
		res := r.handler.HandlerGetAllBook(c)
		return c.SendString(res)
	})

	app.Get("/book", func(c *fiber.Ctx) error {
		res := r.handler.HandlerGetBook(c)
		return c.SendString(res.Name)
	})

	app.Post("/books", func(c *fiber.Ctx) error {
		return c.SendString("insert book")
	})

}
