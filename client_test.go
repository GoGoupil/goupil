package http

import (
	"testing"
)

func TestReturnCode(t *testing.T) {
	client := Client{}
	client.Open("devatoria.info", 80)
	client.Get("/")
	client.Close()
}
