package main

import (
	"fmt"
	"github.com/GoGoupil/http"
	"sync"
)

type Thread struct {
	Count int
	Route string
}

func (t *Thread) Run(c http.Client) {
	fmt.Printf("Launch %d threads requesting route %s...\n", t.Count, t.Route)
	defer fmt.Printf("Threads for route %s are done...\n", t.Route)

	waitGroup := sync.WaitGroup{}

	for i := 0; i < t.Count; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			c.Get(t.Route)
		}()
	}

	waitGroup.Wait()
}
