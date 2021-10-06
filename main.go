package main

import (
	"github.com/andreasaiforesee/template/routes"

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
