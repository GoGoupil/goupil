package main

import (
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
			clients[i].Open(host, port)
		}(i)
	}
	wg.Wait()

	for i, _ := range clients {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer clients[i].Close()
			results, code := clients[i].Get(t.Route)
			t.Results.Cumulate(results)
			if code != 200 {
				t.ErrorRate++
			}
		}(i)
	}
	wg.Wait()

	t.ComputeResult()
	t.Results.Compute(t.Count)
}

func (t *Thread) ComputeResult() {
	t.ErrorRate = (t.ErrorRate * 100) / float64(t.Count)
}
