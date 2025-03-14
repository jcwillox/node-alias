package cmd

import (
	"github.com/jcwillox/node-alias/utils"
	"github.com/spf13/cobra"
)

var outdatedCmd = &cobra.Command{
	Use:                "outdated",
	Aliases:            []string{"o"},
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, remaining []string) {
		var args []string
		manager := utils.GetPackageManager()

		if utils.CmdExists("taze") {
			manager = "taze"
			if len(remaining) == 0 {
				args = []string{"-l", "-I", "-r"}
			}
		}

		utils.RunCommand(manager, args, remaining...)
	},
}

func init() {
	rootCmd.AddCommand(outdatedCmd)
}
