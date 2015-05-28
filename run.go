package main

var runCommand = &Command{
	Usage:     "run",
	ShortDesc: "Run goupil",
	LongDesc:  "TODO",
	MinArgs:   1,
	Run:       run,
}

func run(args []string) {
	var plan = Plan{}
	plan.Load(args[0])
}
