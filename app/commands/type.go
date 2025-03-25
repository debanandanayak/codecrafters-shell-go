package commands

import (
	"fmt"
	"os"
	"slices"

	"github.com/codecrafters-io/shell-starter-go/app/constants"
)

var built_in_list = []string{"exit", "echo", "type"}

func Type(command string) {

	is_builtin := slices.Contains(built_in_list, command)
	if is_builtin {
		fmt.Printf("%s is a shell builtin\n", command)
		return
	} else {
		paths := constants.GetPaths()

		for _, path := range paths {
			entries, _ := os.ReadDir(path)
			for _, entry := range entries {
				if entry.Name() == command {
					fmt.Printf("%s is %s/%s\n", command, path, command)
					return
				}
			}
		}

	}
	fmt.Printf("%s: not found\n", command)
}
