package main

import "fmt"

type Command struct {
	Run       func([]string)
	Usage     string
	ShortDesc string
	LongDesc  string
	MinArgs   int
}

func (command Command) String() string {
	return fmt.Sprintf(`Usage: goupil %s

%s

%s`, command.Usage, command.ShortDesc, command.LongDesc)
}
