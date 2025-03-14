package cmd

import (
	"github.com/jcwillox/node-alias/utils"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:                "add",
	Aliases:            []string{"a"},
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, remaining []string) {
		var args []string
		manager := utils.GetPackageManager()

		if manager == "npm" {
			args = []string{"install"}
		} else {
			args = []string{"add"}
		}

		if len(remaining) > 0 {
			if remaining[0] == "-d" {
				remaining[0] = "-D"
			}
			if remaining[0] == "-D" {
				// TODO: auto determine if dev-deps are exact
				args = append(args, "-E")
			}
		}

		utils.RunCommand(manager, args, remaining...)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
