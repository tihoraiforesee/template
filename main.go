package main

import (
	"github.com/andreasaiforesee/template/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	var r routes.Routes
	r.SetupRoutes(app)

	app.Listen(":3000")
}
