package cmd

import (
	. "github.com/jcwillox/node-alias/utils"
	"github.com/spf13/cobra"
	"os"
)

var useCmd = &cobra.Command{
	Use:                "use",
	Aliases:            []string{"u"},
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, remaining []string) {
		var args []string
		manager := GetPackageManager()

		if CmdExists("nvm") {
			manager = "nvm"
			args = []string{"use"}
			if len(remaining) == 0 {
				if data, err := os.ReadFile(".nvmrc"); err == nil {
					args = append(args, string(data))
				}
			}
		} else if CmdExists("fnm") {
			manager = "fnm"
			args = []string{"use"}
		}

		RunCommand(manager, args, remaining...)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
