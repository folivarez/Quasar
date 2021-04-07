package main

import (
	"github.com/federicolivarez/challengeMeli/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	// give response when at /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint 😉",
		})
	})

	// api group
	api := app.Group("/api/v1")

	// give response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint 😉",
		})
	})

	// connect todo routes
	routes.TodoRoute(api.Group("/"))
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	err := app.Listen(":8080")

	if err != nil {
		panic(err)
	}
}
