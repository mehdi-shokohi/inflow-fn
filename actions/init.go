package actions

import (
	getArgs "github.com/mehdi-shokohi/inflow-fn/actions/get"
	"github.com/mehdi-shokohi/inflow-fn/platform/std"
)

func init() {
	std.RegisterCommand("args",&getArgs.MyAction{},std.Describe{})
}
