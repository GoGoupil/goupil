package main

import (
	"encoding/json"
	"fmt"
	"github.com/GoGoupil/http"
	"io/ioutil"
	"log"
)

type Plan struct {
	Name    string
	BaseURL string
	Threads []Thread
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

	client := http.Client{
		BaseURL: p.BaseURL,
	}

	for _, thread := range p.Threads {
		thread.Run(client)
	}
}
