package main

import "fmt"

var helpCommand = Command {
	Usage: "help [command]",
	ShortDesc: "Display help message",
	LongDesc: `If command is passed, the output will display how to use the passed command.`,
}

func (command *Command) Run(args []string) {
	if len(args) == 0 {
	} else {
		if command, ok := commands[args[0]]; ok {
			fmt.Println(command)
		} else {
			fmt.Printf("Command %s doesn't exist.", args[0])
		}
	}
}
