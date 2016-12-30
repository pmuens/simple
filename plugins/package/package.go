package plugin

import (
	"github.com/pmuens/simple/plugins"
	"github.com/pmuens/simple/plugins/package/sam"
	"github.com/pmuens/simple/util"
)

func init() {
	plugin := plugins.Plugin{}
	plugin.Name = "core:package"
	plugin.Command = "package"
	plugin.ShortDesc = "Packages your service"
	plugin.LongDesc = "Packages your Simple service"
	plugin.FuncToRun = packageService

	plugins.Register(plugin)
}

func packageService() {
	util.Log("Packaging...")

	serviceDir := util.GetServiceDir()

	funcs := GetFunctionDirectoryNames(serviceDir)

	sam := sam.NewSAM(funcs)
	createStackYaml := sam.GetCreateStackYAML()
	updateStackYaml := sam.GetUpdateStackYAML()

	CreateYAMLFile("create-stack.yml", createStackYaml)
	CreateYAMLFile("update-stack.yml", updateStackYaml)
	ZipService()

	util.Log("Successfully packaged service...")
}
