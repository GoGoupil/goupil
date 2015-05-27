package http

import (
	"net/http"
	"log"
)

type Client struct {
	baseURL string
}

func (c *Client) Get(route string) int {
	result, err := http.Get(c.baseURL + route)
	if err != nil {
		log.Fatal(err)
	}
	
	return result.StatusCode
}