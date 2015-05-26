package main

import (
	"fmt"
	"os"
)

type Command struct {
	Name string
	Usage string
	ShortDesc string
	LongDesc string
}

func (command Command) String() string {
	return fmt.Sprintf(`Usage: %s

%s

%s`, command.Usage, command.ShortDesc, command.LongDesc)
}

var commands = []Command{
	helpCommand,
}

func main() {
	args := os.Args[1:]
	
	if len(args) > 0 {
		switch args[0] {
		case "help":
			helpCommand.List(commands)
		default:
			fmt.Println("Unexpected command given. Use help command to display command list.")
		}
	} else {
		fmt.Println("No command name given. Use help command to display command list.")
	}
}
