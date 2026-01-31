package cmd

import (
	"maps"
	"os"
	"strings"

	"github.com/jcwillox/node-alias/utils"
	"github.com/spf13/cobra"
)

var patchRemoveCmd = &cobra.Command{
	Use:                "patch-remove",
	DisableFlagParsing: true,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		if len(args) > 0 {
			return nil, cobra.ShellCompDirectiveDefault
		}
		pkg := utils.GetPackageJson()
		if pkg == nil {
			return nil, cobra.ShellCompDirectiveDefault
		}

		patches := make(map[string]string)
		maps.Copy(patches, pkg.PatchedDependencies)
		maps.Copy(patches, pkg.Pnpm.PatchedDependencies)

		// try to find in the patches directory
		if len(patches) == 0 {
			entries, _ := os.ReadDir("patches")
			for _, entry := range entries {
				name := entry.Name()
				patches[strings.TrimSuffix(name, ".patch")] = "patches/" + name
			}
		}

		completions := make([]cobra.Completion, 0, len(patches))
		for name, desc := range patches {
			completions = append(completions, cobra.CompletionWithDesc(name, desc))
		}
		return completions, cobra.ShellCompDirectiveDefault
	},
	Run: utils.DefaultRunCommand,
}

func init() {
	rootCmd.AddCommand(patchRemoveCmd)
}
