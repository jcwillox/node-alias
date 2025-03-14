package cmd

import (
	"github.com/jcwillox/node-alias/utils"
	"github.com/spf13/cobra"
	"path/filepath"
	"slices"
)

var rootCmd = &cobra.Command{
	Use:                "node-alias",
	DisableSuggestions: true,
	DisableFlagParsing: true,
	SilenceUsage:       true,
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, remaining []string) {
		var args []string
		manager := utils.GetPackageManager()

		if len(remaining) == 0 {
			if utils.CmdExists("tsx") {
				manager = "tsx"
			} else if utils.CmdExists("bun") {
				manager = "bun"
				args = []string{"repl"}
			} else {
				manager = "node"
			}
		}

		utils.RunCommand(manager, args, remaining...)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// TODO does start or test need to be special?
	commands := map[string]string{
		"install":        "i",
		"list":           "ls",
		"start":          "s",
		"test":           "t",
		"lint":           "l",
		"approve-builds": "ab",
		"publish":        "publish",
		"pack":           "pack",
	}
	for use, alias := range commands {
		rootCmd.AddCommand(&cobra.Command{
			Use:                use,
			Aliases:            []string{alias},
			DisableFlagParsing: true,
			Run:                utils.DefaultRunCommand,
		})
	}
}
