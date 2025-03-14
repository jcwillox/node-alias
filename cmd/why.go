package cmd

import (
	"github.com/jcwillox/node-alias/utils"
	"github.com/spf13/cobra"
)

var whyCmd = &cobra.Command{
	Use:                "why",
	Aliases:            []string{"w"},
	DisableFlagParsing: true,
	ValidArgsFunction:  utils.CompleteDependencies,
	Run: func(cmd *cobra.Command, remaining []string) {
		var args []string
		manager := utils.GetPackageManager()

		if manager == "poetry" {
			args = []string{"show", "--tree", "--why"}
		} else {
			args = []string{"why"}
		}

		utils.RunCommand(manager, args, remaining...)
	},
}

func init() {
	rootCmd.AddCommand(whyCmd)
}
