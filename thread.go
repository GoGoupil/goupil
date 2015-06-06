package main

import (
	"github.com/GoGoupil/http"
	"sync"
)

type Thread struct {
	Count       int
	Route       string
	AverageTime float64
	ErrorRate   float64
}

func (t *Thread) Run(host string, port int) {
	wg := sync.WaitGroup{}
	for i := 0; i < t.Count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := http.Client{}
			c.Open(host, port)
			defer c.Close()
			elapsed, code := c.Get(t.Route)
			t.AverageTime += elapsed
			if code != 200 {
				t.ErrorRate++
			}
		}()
	}
	wg.Wait()
	t.ComputeResult()
}

func (t *Thread) ComputeResult() {
	t.AverageTime /= float64(t.Count)
	t.ErrorRate = (t.ErrorRate * 100) / float64(t.Count)
}
