package main

var runCommand = &Command{
	Usage:     "run plan.json",
	ShortDesc: "Execute given plan",
	LongDesc: `The plan parameter is a JSON filed describing the plan to execute.
Here is an example of plan description file.
You can check Plan structure defined in plan.go or check our GitHub to have more details :)

{
	"Name": "test",
	"BaseURL": "http://www.google.com",
	"Threads": [
		{
			"Count": 10,
			"Route": "/"
		},
		{
			"Count": 10,
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
