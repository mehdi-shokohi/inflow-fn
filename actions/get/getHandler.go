package getArgs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mehdi-shokohi/inflow-fn/platform/std"
)



func GetHandler(c *fiber.Ctx)error{
	body,err := std.GetBodyAs[User](c)
	if err!=nil{
		return c.JSON(fiber.ErrBadRequest)
	}
	header,err := std.GetHeadersAs[map[string]interface{}](c)
	if err!=nil{
		return c.JSON(fiber.ErrBadRequest)
	}
	return c.JSON(struct{
		Body *User
		Header map[string]interface{}}{body,*header})
}