package main

import (
	"encoding/json"
	"fmt"
	"github.com/GoGoupil/http"
	"io/ioutil"
	"log"
	"sync"
)

type Plan struct {
	Name    string
	BaseURL string
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
	fmt.Printf("Running %s plan...\n", p.Name)
	defer p.DisplayResult()

	client := http.Client{
		BaseURL: p.BaseURL,
	}

	wg := sync.WaitGroup{}

	for _, thread := range p.Threads {
		wg.Add(1)
		go func(thread *Thread) {
			defer wg.Done()
			thread.Run(client)
		}(thread)
	}

	wg.Wait()
}

func (p *Plan) DisplayResult() {
	fmt.Println("Plan execution is terminated. Displaying results...\n")

	for _, thread := range p.Threads {
		fmt.Printf("Results for route %s:\n", thread.Route)
		fmt.Printf("---------------------\n")
		fmt.Printf("Error rate: %f", thread.ErrorRate)
		fmt.Printf("\n\n")
	}
}
