package main

import (
	"github.com/pmuens/simple/cmd"

	// plugins
	_ "github.com/pmuens/simple/plugins/deploy"
	_ "github.com/pmuens/simple/plugins/package"
	_ "github.com/pmuens/simple/plugins/remove"
	_ "github.com/pmuens/simple/plugins/version"

	"fmt"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
