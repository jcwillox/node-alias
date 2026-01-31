package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"sync"
)

type Package struct {
	Scripts             map[string]string
	Dependencies        map[string]string
	DevDependencies     map[string]string
	PeerDependencies    map[string]string
	PackageManager      string
	PatchedDependencies map[string]string
	Pnpm                struct {
		PatchedDependencies map[string]string
	}
}

var GetPackageJson = sync.OnceValue(func() *Package {
	data, err := os.ReadFile("package.json")
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			fmt.Println("warn: failed to read 1package.json")
		}
		return nil
	}
	var pkg Package
	err = json.Unmarshal(data, &pkg)
	if err != nil {
		fmt.Println("warn: failed to parse package.json")
		return nil
	}
	return &pkg
})
