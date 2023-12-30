package std

import "github.com/gofiber/fiber/v2"


var actions map[string]func(*fiber.Ctx)error


func RegisterCommand(actionId string,action func(ctx *fiber.Ctx)error){
	if actions==nil{
		actions = make(map[string]func(*fiber.Ctx) error)
	}
	// validate action name
	
	actions[actionId]=action

}

func GetActionFunc(key string)func(*fiber.Ctx)error{
	if actions==nil{
		actions = make(map[string]func(*fiber.Ctx) error)

	}
	if f,ok:= actions[key];ok{
		return f
	}
	return nil
}
func GetActions()[]string{
	if actions==nil{
		actions = make(map[string]func(*fiber.Ctx) error)

	}
	actList:=make([]string,0)
	for  k,_:=range actions{
		actList = append(actList, k)
	}
	return actList
}