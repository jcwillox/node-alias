package cmd

import (
	"github.com/jcwillox/node-alias/utils"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:                "run",
	Aliases:            []string{"r"},
	DisableFlagParsing: true,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		pkg := utils.GetPackageJson()
		completions := make([]cobra.Completion, 0, len(pkg.Scripts))

		for script, desc := range pkg.Scripts {
			completions = append(completions, cobra.CompletionWithDesc(script, desc))
		}

		return completions, cobra.ShellCompDirectiveDefault
	},
	Run: utils.DefaultRunCommand,
}

func init() {
	rootCmd.AddCommand(runCmd)
}
