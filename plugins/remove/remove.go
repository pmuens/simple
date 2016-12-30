package plugin

import (
	"github.com/pmuens/simple/plugins"
	"github.com/pmuens/simple/util"
)

func init() {
	plugin := plugins.Plugin{}
	plugin.Name = "core:remove"
	plugin.Command = "remove"
	plugin.ShortDesc = "Removes your whole service"
	plugin.LongDesc = "Removes your whole service with all its dependencies from AWS"
	plugin.FuncToRun = removeService

	plugins.Register(plugin)
}

func removeService() {
	util.Log("Removing (this might take a few seconds)...")

	EmptyBucket()
	DeleteStack()

	util.Log("Successfully removed service...")
}
