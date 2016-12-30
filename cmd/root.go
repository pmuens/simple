package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "simple",
	Short: "Simple CLI tool",
	Long:  "Simple helps you manage and deploy serverless applications accross multiple providers",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
