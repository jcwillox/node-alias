package cmd

import (
	"github.com/jcwillox/node-alias/utils"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:                "remove",
	Aliases:            []string{"rm"},
	DisableFlagParsing: true,
	ValidArgsFunction:  utils.CompleteDependencies,
	Run:                utils.DefaultRunCommand,
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
