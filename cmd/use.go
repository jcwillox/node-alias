package cmd

import (
	"github.com/jcwillox/node-alias/utils"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:                "use",
	Aliases:            []string{"u"},
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, remaining []string) {
		manager := "nvm"
		args := []string{"use"}

		//if (($remainingArgs.Length -eq 0) -and (Test-Path .\.nvmrc)) {
		//	$cmdArgs += "$(Get-Content .nvmrc)"
		//}

		utils.RunCommand(manager, args, remaining...)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
