package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

type Plan struct {
	Host    string
	Port    int
	Threads []*Thread
}

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

func (p *Plan) Run() {
	fmt.Printf("Running plan on %s:%d\n", p.Host, p.Port)
	wg := sync.WaitGroup{}
	for _, thread := range p.Threads {
		wg.Add(1)
		go func(t *Thread) {
			defer wg.Done()
			fmt.Printf("Running %d threads on route %s\n", t.Count, t.Route)
			t.Run(p.Host, p.Port)
		}(thread)
	}
	wg.Wait()
	fmt.Printf("Ending plan, computing results\n")
	p.DisplayResult()
}

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
		fmt.Printf("Error rate: %f%%\n", thread.ErrorRate)
	}
	fmt.Println()
}
