package pathutils

import "strings"

// GetURLPath
// "aaaa://bbbb" => "bbbb".
func GetURLPath(url string) string {
	res := strings.Split(url, "://")
	return res[1]
}
