package getArgs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn/platform/std"
)

type MyAction struct{}

func (a *MyAction) Run(c *fiber.Ctx) error {
	body, err := std.GetBodyAs[User](c)
	if err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}

	return std.Send(c, fiber.StatusOK, body)
}

func (a *MyAction) FnArguments(c *fiber.Ctx) error {
	// Return Form Application Model
	return std.Send(c, fiber.StatusOK, User{})
}
func (a *MyAction) Settings(c *fiber.Ctx) error {
	// Global Setting for Fn
	if c.Method() == fiber.MethodGet {

		return std.Send(c, fiber.StatusOK, "show settings fields and model")
	}
	if c.Method() == fiber.MethodPut {
		return std.Send(c, fiber.StatusOK, "new value")

	}
	return nil
}
