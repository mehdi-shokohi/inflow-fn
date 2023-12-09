package fn

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn/fn/std"
)


func RegisterInflowFnV1(api fiber.Router) {

	api.Group("/v1")
	std.RouteDecision(api)
}