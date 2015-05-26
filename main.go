package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	
	if len(args) > 0 {
		switch args[0] {
		default:
			fmt.Println("Wrong command name given")
		}
	} else {
		fmt.Println("No command name given")
	}
}
