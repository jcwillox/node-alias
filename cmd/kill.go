package cmd

import (
	"github.com/jcwillox/node-alias/utils"
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:                "kill",
	Aliases:            []string{"k"},
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, remaining []string) {
		if utils.CmdExists("npkill") {
			utils.RunCommand("npkill", remaining)
		} else {
			utils.RunCommand(utils.GetPackageManager(), []string{"kill"}, remaining...)
		}
	},
}

func init() {
	rootCmd.AddCommand(killCmd)
}
