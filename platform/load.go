package platform

import (
	"github.com/gofiber/fiber/v2"
	inflowV1 "github.com/mehdi-shokohi/inflow-fn/platform/inflowFnV1"
)


func RegisterFunctions(api fiber.Router) {
	apiGroup:=api.Group("/fn")
	inflowV1.RegisterInflowFnV1(apiGroup)
}

