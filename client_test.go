package http

import (
	"log"
	"sync"
	"testing"
)

func TestReturnCode(t *testing.T) {
	socketCount := 100
	socketHost := "devatoria.info"
	socketPort := 80
	remainingSockets := 0
	clients := make([]Client, socketCount)
	var totalElapsed float64
	var totalError int
	wg := sync.WaitGroup{}

	for i, _ := range clients {
		clients[i].Open(socketHost, socketPort)
		remainingSockets++
		log.Printf("%d sockets opened...\n", remainingSockets)
	}

	log.Printf("Opened %d sockets on %s:%d\n", socketCount, socketHost, socketPort)

	for i, _ := range clients {
		wg.Add(1)
		go func(i int) {
			defer func() {
				wg.Done()
				remainingSockets--
				log.Printf("%d sockets remaining...\n", remainingSockets)
			}()
			elapsed, code := clients[i].Get("/")
			totalElapsed += elapsed
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

	log.Printf("Elapsed: %fms\n", totalElapsed/float64(socketCount))
	log.Printf("Error percentile: %f%%\n", float64((totalError/socketCount)*100))
}
