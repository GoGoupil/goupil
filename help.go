package main

import "fmt"

var helpCommand = &Command{
	Usage:     "help [command]",
	ShortDesc: "Display help message",
	LongDesc:  `If command is passed, the output will display how to use the passed command.`,
}

func init() {
	helpCommand.Run = help
}

func help(args []string) {
	if len(args) == 0 {
		fmt.Println("Goupil is a web load testing tool designed in Go")
		fmt.Println("Here is the list of available commands")
		fmt.Println("Use goupil help [command] to get more info about a specific command\n")

		for name, command := range commands {
			fmt.Printf("%s - %s\n", name, command.ShortDesc)
		}
	} else {
		if command, ok := commands[args[0]]; ok {
			fmt.Println(command)
		} else {
			fmt.Printf("Command %s doesn't exist.", args[0])
		}
	}
}
