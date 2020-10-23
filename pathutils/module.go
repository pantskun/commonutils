package pathutils

import (
	"os"
	"strings"
)

// GetModulePath
// get current module path.
func GetModulePath(moduleName string) string {
	fp, err := os.Getwd()
	if err != nil {
		return ""
	}

	fp = ConvertBackslashToSlash(fp)
	fp = strings.SplitAfter(fp, moduleName)[0]

	return fp
}
