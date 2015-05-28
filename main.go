package main

import (
	"fmt"
	"os"
)

var commands = map[string]*Command{
	"help": helpCommand,
	"run":  runCommand,
}

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		if command, ok := commands[args[0]]; ok {
			if len(args[1:]) >= command.MinArgs {
				command.Run(args[1:])
			} else {
				fmt.Printf("Too few arguments. Use goupil help %s to display command usage.\n", args[0])
			}
		} else {
			fmt.Println("Unexpected command given. Use help command to display command list.")
		}
	} else {
		fmt.Println("No command name given. Use help command to display command list.")
	}
}
