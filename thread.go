package main

import (
	"github.com/GoGoupil/http"
	"sync"
)

type Thread struct {
	Count     int
	Route     string
	Results   AverageResults
	ErrorRate float64
}

type AverageResults struct {
	AverageSendingTime           float64
	AverageReadingFirstBytesTime float64
	AverageReadingTotalTime      float64
	AverageTotalTime             float64
}

func (ar *AverageResults) Cumulate(results http.Result) {
	ar.AverageSendingTime += results.TimeSending
	ar.AverageReadingFirstBytesTime += results.TimeReadingFirstBytes
	ar.AverageReadingTotalTime += results.TimeReadingTotal
	ar.AverageTotalTime += results.TimeTotal
}

func (ar *AverageResults) Compute(count int) {
	ar.AverageSendingTime /= float64(count)
	ar.AverageReadingFirstBytesTime /= float64(count)
	ar.AverageReadingTotalTime /= float64(count)
	ar.AverageTotalTime /= float64(count)
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
			results, code := c.Get(t.Route)
			t.Results.Cumulate(results)
			if code != 200 {
				t.ErrorRate++
			}
		}()
	}
	wg.Wait()
	t.ComputeResult()
	t.Results.Compute(t.Count)
}

func (t *Thread) ComputeResult() {
	t.ErrorRate = (t.ErrorRate * 100) / float64(t.Count)
}
