package main

import (
	"github.com/jcwillox/node-alias/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
