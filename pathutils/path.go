package pathutils

import (
	"os"
	"strings"

	"github.com/pkg/sftp"
)

// ConvertBackslashToSlash
// Convert "\\" to "/".
func ConvertBackslashToSlash(s string) string {
	ss := strings.Split(s, "\\")
	res := ""

	for i, t := range ss {
		res += t

		if i < len(ss)-1 {
			res += "/"
		}
	}

	return res
}

// GetParentPath
// get parent path.
func GetParentPath(p string) string {
	index := strings.LastIndex(p, "/")
	if index == -1 {
		return p
	}

	if index == len(p)-1 {
		p = p[0:index]
		index = strings.LastIndex(p, "/")
	}

	return p[0:index]
}

func IsDir(p string) (bool, error) {
	info, err := os.Stat(p)
	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
}

func IsDirWithSftp(p string, sftpClient *sftp.Client) (bool, error) {
	info, err := sftpClient.Stat(p)
	if err != nil {
		return false, err
	}

	return info.IsDir(), nil
}
