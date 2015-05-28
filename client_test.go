package http

import (
	"github.com/GoGoupil/assert"
	"testing"
)

func TestReturnCode(t *testing.T) {
	var client = Client{
		BaseURL: "http://www.google.fr",
	}

	code := client.Get("/")

	assert.AssertEqual(t, code, 200)
}
