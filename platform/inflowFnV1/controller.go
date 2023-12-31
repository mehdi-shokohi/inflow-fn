package inflowV1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn/platform/std"
) 


func actionHandlers(c *fiber.Ctx)error{
	action:=c.Params("action")
	fn:=std.GetAction(action)
	if fn==nil{
		return std.Send(c,fiber.StatusNotFound,nil,fiber.ErrNotFound)
	}
	return fn.Run(c)

}

func getActions(c *fiber.Ctx)error{
	return c.JSON(std.GetActions())
}