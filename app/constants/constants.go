package constants

import (
	"os"
	"strings"
)

func GetPaths() []string {
	pathString := os.Getenv("PATH")
	paths := strings.Split(pathString, ":")
	return paths
}
