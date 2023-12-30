package actions

import (
	action "github.com/mehdi-shokohi/inflow-fn/actions/get"
	"github.com/mehdi-shokohi/inflow-fn/platform/std"
)
func init(){
	std.RegisterCommand("args",action.GetHandler)
}