package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Plan struct {
	Name string
}

func (p *Plan) Load(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, p)
}
