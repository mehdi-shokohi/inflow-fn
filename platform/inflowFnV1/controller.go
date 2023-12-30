package inflowV1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn/platform/std"
) 


func actionHandlers(c *fiber.Ctx)error{
	action:=c.Params("action")
	fn:=std.GetActionFunc(action)
	if fn!=nil{
		return fn(c)
	}
	return std.Send(c,fiber.StatusNotFound,nil,fiber.ErrNotFound)
}

func getActions(c *fiber.Ctx)error{
	return c.JSON(std.GetActions())
}