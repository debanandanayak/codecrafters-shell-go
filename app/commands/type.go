package commands

import (
	"fmt"
	"slices"
)

func Type(args []string, builtin []string) {
	is_builtin := slices.Contains(builtin, args[1])
	if is_builtin {
		fmt.Println(args[1] + " is a shell builtin")
	} else {
		fmt.Println(args[1] + ": not found")
	}
}
