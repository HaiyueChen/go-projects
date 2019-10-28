package utils

import (
	"strings"
)

// ParseFilenameFromPath return string
func ParseFilenameFromPath(path string) string {
	if len(strings.TrimSpace(path)) == 0 {
		return ""
	}

	pathComponets := strings.Split(path, "/")

	filename := strings.Split(pathComponets[len(pathComponets)-1], ".")[0]
	return filename
}
