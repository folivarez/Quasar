package routes

import (
	"github.com/federicolivarez/challengeMeli/controllers"
	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route fiber.Router) {
	route.Post("/topsecret", controllers.GetLocation)
}
