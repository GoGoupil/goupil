package http

import (
	"github.com/GoGoupil/assert"
	"sync"
	"testing"
)

func TestReturnCode(t *testing.T) {
	var client = Client{
		BaseURL: "http://devatoria.info",
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			code := client.Get("/")
			assert.AssertEqual(t, code, 200)
		}(i)
	}

	wg.Wait()
}
