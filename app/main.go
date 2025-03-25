package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage

	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input", err)
			os.Exit(1)
		}

		command = command[:len(command)-1]
		args := strings.Split(command, " ")
		switch args[0] {
		case "exit":
			if args[1] == "0" {
				os.Exit(0)
			}
		case "echo":
			commands.Echo(command[5:])
		case "type":
			commands.Type(args[1])
		default:
			fmt.Println(command + ": command not found")
		}

	}
}
