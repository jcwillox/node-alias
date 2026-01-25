package utils

import (
	"bufio"
	"os"
	"runtime"
	"strings"
)

// GetShebang returns the interpreter and args from a file's shebang line.
// If no shebang is present, returns empty manager and nil args.
func GetShebang(filePath string) (string, []string) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	prefix := make([]byte, 2)
	_, err = reader.Read(prefix)
	if err != nil || string(prefix) != "#!" {
		return "", nil
	}

	line, _ := reader.ReadString('\n')
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return "", nil
	}

	manager := parts[0]
	args := parts[1:]

	if runtime.GOOS == "windows" && manager == "/usr/bin/env" && len(args) > 0 {
		manager = args[0]
		args = args[1:]
	}

	return manager, args
}
