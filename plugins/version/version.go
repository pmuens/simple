package plugin

import (
	"fmt"

	"github.com/pmuens/simple/plugins"
	"github.com/pmuens/simple/util"
)

const version = "0.1.0" // TODO read version from VERSION file

func init() {
	plugin := plugins.Plugin{}
	plugin.Name = "core:version"
	plugin.Command = "version"
	plugin.ShortDesc = "Print the version number of Simple"
	plugin.LongDesc = "Full version number of Simple"
	plugin.FuncToRun = func() {
		v := fmt.Sprintf("CLI v%s\n", version)
		util.Log(v)
	}

	plugins.Register(plugin)
}
