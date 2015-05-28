package http

import (
	"log"
	"net/http"
)

type Client struct {
	BaseURL string
}

func (c *Client) Get(route string) int {
	result, err := http.Get(c.BaseURL + route)
	if err != nil {
		log.Fatal(err)
	}

	return result.StatusCode
}
