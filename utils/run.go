package utils

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/sys/execabs"
	"os"
	"strings"
)

func RunCommand(name string, args []string, remaining ...string) {
	RunAliasCommand(name, name, args, remaining...)
}

func RunAliasCommand(name string, alias string, args []string, remaining ...string) {
	args = append(args, remaining...)

	fmt.Printf("\x1b[90m$ %s %s\x1b[0m\n", alias, strings.Join(args, " "))

	cmd := execabs.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		var exitErr *execabs.ExitError
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.ExitCode())
		} else {
			fmt.Println("error:", err)
			os.Exit(1)
		}
	}
}

func DefaultRunCommand(cmd *cobra.Command, remaining []string) {
	RunCommand(GetPackageManager(), []string{cmd.Use}, remaining...)
}
