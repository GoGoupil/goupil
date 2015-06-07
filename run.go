package main

var runCommand = &Command{
	Usage:     "run plan.json",
	ShortDesc: "Execute given plan",
	LongDesc: `The plan parameter is a JSON filed describing the plan to execute.
Here is an example of plan description file.
You can check Plan structure defined in plan.go or check our GitHub to have more details :)

{
	"Host": "devatoria.info",
	"Port": 80,
	"Threads": [
		{
			"Count": 30,
			"Route": "/"
		},
		{
			"Count": 30,
			"Route": "/test"
		}
	]
}`,
	MinArgs: 1,
	Run:     run,
}

func run(args []string) {
	var plan = Plan{}
	plan.Load(args[0])
}
