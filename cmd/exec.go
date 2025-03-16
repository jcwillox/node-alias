package cmd

import (
	"github.com/jcwillox/node-alias/utils"
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
)

var execCmd = &cobra.Command{
	Use:                "exec",
	Aliases:            []string{"x"},
	DisableFlagParsing: true,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		if len(args) > 0 {
			return nil, cobra.ShellCompDirectiveDefault
		}

		pkg := utils.GetPackageJson()
		completions := make([]cobra.Completion, 0, len(pkg.Scripts))

		entries, err := os.ReadDir("./node_modules/.bin")
		if err != nil {
			return completions, cobra.ShellCompDirectiveError
		}

		for _, entry := range entries {
			if entry.IsDir() || filepath.Ext(entry.Name()) != "" {
				continue
			}
			completions = append(
				completions,
				cobra.CompletionWithDesc(entry.Name(), path.Join("node_modules/.bin", entry.Name())),
			)
		}

		return completions, cobra.ShellCompDirectiveDefault
	},
	Run: func(cmd *cobra.Command, remaining []string) {
		var args []string
		manager := utils.GetPackageManager()

		if len(remaining) > 0 {
			bin := path.Join("./node_modules/.bin/", remaining[0])

			if _, err := os.Stat(bin); err == nil {
				utils.RunAliasCommand(bin, remaining[0], remaining[1:])
				return
			}

			if manager == "npm" {
				manager = "npx"
			} else if manager == "pnpm" {
				args = []string{"exec"}
			}

			utils.RunCommand(manager, args, remaining...)
		}
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
