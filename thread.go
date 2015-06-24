package main

import (
	"fmt"
	"sync"
	"time"
)

type Thread struct {
	Duration  int
	Gap       int
	Count     int
	Route     string
	Method    string
	Params    map[string]string
	Results   AverageResults
	ErrorRate float64
}

type AverageResults struct {
	AverageSendingTime           float64
	AverageReadingFirstBytesTime float64
	AverageReadingTotalTime      float64
	AverageTotalTime             float64
	MinTotalTime                 float64
	MaxTotalTime                 float64
}

func (ar *AverageResults) Cumulate(results Result) {
	ar.AverageSendingTime += results.TimeSending
	ar.AverageReadingFirstBytesTime += results.TimeReadingFirstBytes
	ar.AverageReadingTotalTime += results.TimeReadingTotal
	ar.AverageTotalTime += results.TimeTotal

	if results.TimeTotal < ar.MinTotalTime || ar.MinTotalTime == 0 {
		ar.MinTotalTime = results.TimeTotal
	}
	if results.TimeTotal > ar.MaxTotalTime || ar.MaxTotalTime == 0 {
		ar.MaxTotalTime = results.TimeTotal
	}
}

func (ar *AverageResults) Compute(count int) {
	ar.AverageSendingTime /= float64(count)
	ar.AverageReadingFirstBytesTime /= float64(count)
	ar.AverageReadingTotalTime /= float64(count)
	ar.AverageTotalTime /= float64(count)
}

func (t *Thread) Run(host string, port int) {
	wg := sync.WaitGroup{}
	clients := make([]Client, t.Count)

	for i, _ := range clients {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			clients[i].NewClient(host, port)
		}(i)
	}
	wg.Wait()

	reqTotal := t.Count * (t.Duration / t.Gap)
	for reqCount := 0 ; reqCount < reqTotal ; reqCount += t.Count {
		fmt.Printf("New wave... (%d/%d requests)\n", (reqCount + t.Count), reqTotal)
		for i, _ := range clients {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				results, code := clients[i].Get(t.Route, t.Method, t.Params)
				t.Results.Cumulate(results)
				if code != 200 && code != 302 {
					t.ErrorRate++
				}
			}(i)
		}
		time.Sleep(time.Duration(t.Gap) * time.Millisecond)
	}
	fmt.Println("Done. Waiting for remaining connections to finish... This can take a while according to your web server.")
	wg.Wait()
	
	for i, _ := range clients {
		clients[i].Close()
	}

	t.ComputeResult()
	t.Results.Compute(t.Count * (t.Duration / t.Gap))
}

func (t *Thread) ComputeResult() {
	t.ErrorRate = (t.ErrorRate * 100) / float64(t.Count * (t.Duration / t.Gap))
}
