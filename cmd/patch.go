package cmd

import (
	"github.com/jcwillox/node-alias/utils"
	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:                "patch",
	DisableFlagParsing: true,
	ValidArgsFunction:  utils.CompleteDependencies,
	Run:                utils.DefaultRunCommand,
}

func init() {
	rootCmd.AddCommand(patchCmd)
}
