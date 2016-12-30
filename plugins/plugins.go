package plugins

import (
	"github.com/pmuens/simple/cmd"
	"github.com/spf13/cobra"
)

type Plugin struct {
	Name      string
	Command   string
	ShortDesc string
	LongDesc  string
	FuncToRun func()
}

func Register(plugin Plugin) {
	addToPluginsArray(plugin)
	addCommandToCli(plugin)
}

var plugins = make(map[string]Plugin)

func addToPluginsArray(plugin Plugin) {
	plugins[plugin.Name] = plugin
}

func addCommandToCli(plugin Plugin) {
	cmmd := &cobra.Command{
		Use:   plugin.Command,
		Short: plugin.ShortDesc,
		Long:  plugin.LongDesc,
		Run: func(cmd *cobra.Command, args []string) {
			plugin.FuncToRun()
		},
	}

	cmd.RootCmd.AddCommand(cmmd)
}
