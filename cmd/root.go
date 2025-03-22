package cmd

import (
	"github.com/jcwillox/node-alias/constants"
	. "github.com/jcwillox/node-alias/utils"
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
		manager := GetPackageManager()

		if len(remaining) == 0 {
			if CmdExists("tsx") {
				manager = "tsx"
			} else if CmdExists("bun") {
				manager = "bun"
				args = []string{"repl"}
			} else {
				manager = "node"
			}
		}

		if len(remaining) > 0 {
			if ext := filepath.Ext(remaining[0]); ext != "" {
				if slices.Contains(constants.NodeExtensions, ext) {
					if CmdExists("tsx") {
						manager = "tsx"
					} else if CmdExists("bun") {
						manager = "bun"
					} else {
						manager = "node"
					}
				} else if ext == ".py" {
					if CmdExists("uv") {
						manager = "uv"
						args = []string{"run"}
					} else if CmdExists("python") {
						manager = "python"
					}
				}
			}
		}

		RunCommand(manager, args, remaining...)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:                "help",
		Aliases:            []string{"h"},
		DisableFlagParsing: true,
		Run:                DefaultRunCommand,
	})

	commands := map[string]string{
		"install":        "i",
		"approve-builds": "ab",
		"publish":        "publish",
		"pack":           "pack",
	}

	for use, alias := range commands {
		rootCmd.AddCommand(&cobra.Command{
			Use:                use,
			Aliases:            []string{alias},
			DisableFlagParsing: true,
			Run:                DefaultRunCommand,
		})
	}

	scriptCommands := map[string]string{
		"dev":   "d",
		"start": "s",
		"test":  "t",
		"lint":  "l",
	}

	for use, alias := range scriptCommands {
		rootCmd.AddCommand(&cobra.Command{
			Use:                use,
			Aliases:            []string{alias},
			DisableFlagParsing: true,
			Run:                RunScriptCommand,
		})
	}
}
