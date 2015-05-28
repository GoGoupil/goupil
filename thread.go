package main

import (
	"fmt"
	"github.com/GoGoupil/http"
	"sync"
	"time"
)

type Thread struct {
	Count       int
	Route       string
	ErrorRate   float64
	AverageTime float64
}

func (t *Thread) Run(c http.Client) {
	fmt.Printf("Launch %d threads requesting route %s...\n", t.Count, t.Route)
	defer t.ComputeResult()

	waitGroup := sync.WaitGroup{}

	for i := 0; i < t.Count; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()

			start := time.Now()
			code := c.Get(t.Route)
			t.AverageTime += time.Since(start).Seconds()

			if code != 200 {
				t.ErrorRate++
			}
		}()
	}

	waitGroup.Wait()
}

func (t *Thread) ComputeResult() {
	fmt.Printf("Threads for route %s are done. Calculating results...\n", t.Route)
	t.ErrorRate = (t.ErrorRate * 100) / float64(t.Count)
	t.AverageTime /= float64(t.Count)
}
