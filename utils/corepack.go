package utils

import "strings"

func GetCorepackManager() string {
	pkg := GetPackageJson()
	if pkg == nil {
		return ""
	}
	parts := strings.SplitN(pkg.PackageManager, "@", 2)
	if len(parts) == 0 {
		return ""
	}
	return parts[0]
}
