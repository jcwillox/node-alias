package utils

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/sys/execabs"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func logErr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func RunCommand(name string, args []string, remaining ...string) {
	RunAliasCommand(name, name, args, remaining...)
}

func RunAliasCommand(name string, alias string, args []string, remaining ...string) {
	args = append(args, remaining...)

	_, _ = fmt.Fprintf(os.Stderr, "\x1b[90m$ %s %s\x1b[0m\n", alias, strings.Join(args, " "))

	if runtime.GOOS != "windows" {
		exePath, err := execabs.LookPath(name)
		if err != nil {
			logErr(err)
		}
		err = syscall.Exec(exePath, append([]string{name}, args...), os.Environ())
		if err != nil {
			logErr(err)
		}
	} else {
		cmd := execabs.Command(name, args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Start()
		if err != nil {
			logErr(err)
		}

		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

		go func() {
			<-sigs
			time.Sleep(time.Millisecond * 100)
			err := cmd.Process.Kill()
			if err != nil {
				logErr(err)
			}
		}()

		err = cmd.Wait()
		if err != nil {
			var exitErr *execabs.ExitError
			if errors.As(err, &exitErr) {
				os.Exit(exitErr.ExitCode())
			} else {
				logErr(err)
			}
		}
	}
}

func DefaultRunCommand(cmd *cobra.Command, remaining []string) {
	RunCommand(GetPackageManager(), []string{cmd.Use}, remaining...)
}

func RunScriptCommand(cmd *cobra.Command, remaining []string) {
	RunCommand(GetPackageManager(), []string{"run", cmd.Use}, remaining...)
}
