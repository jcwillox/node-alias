package cmd

import (
	"github.com/jcwillox/node-alias/utils"
	"github.com/spf13/cobra"
)

var dlxCmd = &cobra.Command{
	Use:                "dlx",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, remaining []string) {
		var args []string
		manager := utils.GetPackageManager()

		if manager == "npm" {
			manager = "npx"
		} else if manager == "bun" {
			manager = "bunx"
		} else {
			args = []string{"dlx"}
		}

		utils.RunCommand(manager, args, remaining...)
	},
}

func init() {
	rootCmd.AddCommand(dlxCmd)
}
