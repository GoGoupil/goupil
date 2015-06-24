package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

// Plan structure defining an execution plan
// constitued of host, port, protocol and a set of threads.
type Plan struct {
	Host    string
	Port    int
	Https   bool
	Threads []*Thread
}

// Load function loading JSON file to create the structure.
func (p *Plan) Load(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, p)
	if err != nil {
		log.Fatal(err)
	}

	p.Run()
}

// Run function running all threads.
func (p *Plan) Run() {
	fmt.Printf("Running plan on %s:%d\n", p.Host, p.Port)
	wg := sync.WaitGroup{}
	for _, thread := range p.Threads {
		wg.Add(1)
		go func(t *Thread) {
			defer wg.Done()
			fmt.Printf("Running %d threads sending a new request each %dms during %dms on route %s\n", t.Count, t.Gap, t.Duration, t.Route)
			t.Run(p.Host, p.Port, p.Https)
		}(thread)
	}
	wg.Wait()
	fmt.Printf("Ending plan, computing results\n")
	p.DisplayResult()
}

// DisplayResult function synthetizing the results
// to display them to user.
func (p *Plan) DisplayResult() {
	fmt.Println()
	fmt.Printf("Results:\n")
	for _, thread := range p.Threads {
		fmt.Printf("---------------------------------------\n")
		fmt.Printf("Route: %s\n", thread.Route)
		fmt.Printf("Average sending time: %fms\n", thread.Results.AverageSendingTime)
		fmt.Printf("Average reading first bytes time: %fms\n", thread.Results.AverageReadingFirstBytesTime)
		fmt.Printf("Average reading total time: %fms\n", thread.Results.AverageReadingTotalTime)
		fmt.Printf("Average total time: %fms\n", thread.Results.AverageTotalTime)
		fmt.Printf("Min total time: %fms\n", thread.Results.MinTotalTime)
		fmt.Printf("Max total time: %fms\n", thread.Results.MaxTotalTime)
		fmt.Printf("Error rate: %f%%\n", thread.ErrorRate)
	}
	fmt.Println()
}
