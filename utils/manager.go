package utils

import (
	"golang.org/x/sys/execabs"
	"os"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func CmdExists(cmd string) bool {
	path, err := execabs.LookPath(cmd)
	return err == nil && path != ""
}

func GuessPackageManager() string {
	if FileExists("pnpm-lock.yaml") {
		return "pnpm"
	} else if FileExists("bun.lockb") {
		return "bun"
	} else if FileExists("yarn.lock") {
		return "yarn"
	} else if FileExists("pyproject.toml") {
		return "poetry"
	} else if FileExists("package-lock.json") {
		return "npm"
	} else if CmdExists("pnpm") {
		return "pnpm"
	} else if CmdExists("bun") {
		return "bun"
	}
	return "npm"
}

func GetPackageManager() string {
	manager := GetCorepackManager()
	if manager != "" {
		return manager
	}
	return GuessPackageManager()
}
