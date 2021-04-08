package main

import (
	"github.com/federicolivarez/challengeMeli/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	// "log"
	"fmt"
	"github.com/federicolivarez/challengeMeli/config"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Welcome to Meli Challege Quasar 🚀",
		})
	})

	api := app.Group("/api/v1")

	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Welcome to Meli Challege Quasar 🚀",
		})
	})

	routes.TodoRoute(api.Group("/"))
}

func main() {
	
	configuration := config.GetConfiguration(".")

	app := fiber.New()
	app.Use(logger.New())

	
	setupRoutes(app)
	fmt.Println("port", configuration.Server.Port)

	er := app.Listen(":" + configuration.Server.Port)

	if er != nil {
		panic(er)
	}
}
