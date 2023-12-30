package inflowV1

import "github.com/gofiber/fiber/v2"


func RegisterInflowFnV1(api fiber.Router) {
	api.Post("/:action",actionHandlers)
	api.Get("",getActions)
}