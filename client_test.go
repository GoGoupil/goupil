package http

import (
	"log"
	"sync"
	"testing"
)

type AverageResults struct {
	AverageSendingTime           float64
	AverageReadingFirstBytesTime float64
	AverageReadingTotalTime      float64
	AverageTotalTime             float64
}

func TestReturnCode(t *testing.T) {
	socketCount := 100
	socketHost := "devatoria.info"
	socketPort := 80
	remainingSockets := 0
	clients := make([]Client, socketCount)
	averageResults := AverageResults{}
	var totalError int
	wg := sync.WaitGroup{}

	for i, _ := range clients {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			clients[i].Open(socketHost, socketPort)
			remainingSockets++
			log.Printf("%d sockets opened...\n", remainingSockets)
		}(i)
	}
	wg.Wait()

	log.Printf("Opened %d sockets on %s:%d\n", socketCount, socketHost, socketPort)

	for i, _ := range clients {
		wg.Add(1)
		go func(i int) {
			defer func() {
				wg.Done()
				remainingSockets--
				log.Printf("%d sockets remaining...\n", remainingSockets)
			}()
			results, code := clients[i].Get("/")
			averageResults.AverageSendingTime += results.TimeSending
			averageResults.AverageReadingFirstBytesTime += results.TimeReadingFirstBytes
			averageResults.AverageReadingTotalTime += results.TimeReadingTotal
			averageResults.AverageTotalTime += results.TimeTotal
			if code != 200 {
				totalError++
				log.Println(code)
			}
		}(i)
	}

	wg.Wait()

	for i, _ := range clients {
		clients[i].Close()
	}

	averageResults.AverageSendingTime /= float64(socketCount)
	averageResults.AverageReadingFirstBytesTime /= float64(socketCount)
	averageResults.AverageReadingTotalTime /= float64(socketCount)
	averageResults.AverageTotalTime /= float64(socketCount)

	log.Printf("Average sending time: %fms\n", averageResults.AverageSendingTime)
	log.Printf("Average receiving first bytes time: %fms\n", averageResults.AverageReadingFirstBytesTime)
	log.Printf("Average receiving total time: %fms\n", averageResults.AverageReadingTotalTime)
	log.Printf("Average total time: %fms\n", averageResults.AverageTotalTime)
	log.Printf("Error percentile: %f%%\n", float64((totalError/socketCount)*100))
}
