package main

import "fmt"

type Runner interface {
	Run([]string)
}

type Command struct {
	Usage     string
	ShortDesc string
	LongDesc  string
}

func (command Command) String() string {
	return fmt.Sprintf(`Usage: goupil %s

%s

%s`, command.Usage, command.ShortDesc, command.LongDesc)
}
