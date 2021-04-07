package routes

import (
	"github.com/federicolivarez/challengeMeli/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route fiber.Router) {
	route.Post("/topsecret", controllers.GetLocation)
	// route.Post("", controllers.CreateTodo)
	// route.Put("/:id", controllers.UpdateTodo)
	// route.Delete("/:id", controllers.DeleteTodo)
	// route.Get("/:id", controllers.GetTodo)
}
