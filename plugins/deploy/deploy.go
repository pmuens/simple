package plugin

import (
	"github.com/pmuens/simple/plugins"
	"github.com/pmuens/simple/util"
)

func init() {
	plugin := plugins.Plugin{}
	plugin.Name = "core:deploy"
	plugin.Command = "deploy"
	plugin.ShortDesc = "Deploy your packaged service"
	plugin.LongDesc = "Deploys your previously packaged service to the cloud provider"
	plugin.FuncToRun = deployService

	plugins.Register(plugin)
}

func deployService() {
	util.Log("Deploying (this might take a few seconds)...")

	CreateStack()
	UpdateStack()

	util.Log("Done...")
}
