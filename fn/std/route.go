package std

import "github.com/gofiber/fiber/v2"


func RouteDecision(api fiber.Router) {
	api.Post("/sort",sort).Name("sort")

}