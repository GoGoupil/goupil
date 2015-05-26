package main

import "fmt"

var helpCommand = Command {
	Name: "help",
	Usage: "help",
	ShortDesc: "Display help message",
	LongDesc: "Long help message",
}

func (command *Command) List(commands []Command) {
	for _, value := range commands {
		fmt.Printf(`%s - %s`, value.Name, value.ShortDesc)
		fmt.Println()
	}
}
