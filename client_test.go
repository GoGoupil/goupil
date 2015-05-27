package http

import (
	"testing"
	"github.com/GoGoupil/assert"
)

func TestReturnCode(t *testing.T) {
	var client = Client {
		baseURL: "http://www.google.fr",
	}
	
	code := client.Get("/")
	
	assert.AssertEqual(t, code, 200)
}
