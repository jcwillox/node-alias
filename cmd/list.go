package cmd

import (
	"github.com/jcwillox/node-alias/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:                "list",
	Aliases:            []string{"ls"},
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, remaining []string) {
		var args []string
		manager := utils.GetPackageManager()

		if manager == "bun" {
			args = []string{"pm", "ls"}
		} else {
			args = []string{"list"}
		}

		utils.RunCommand(manager, args, remaining...)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
