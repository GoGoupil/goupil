package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Client struct determining a package of connected sockets
// having the same configuration.
type Client struct {
	Sockets []*net.Conn
	Host   string
	Port   int
	Https  bool
}

// Result struct containing client computed results.
type Result struct {
	TimeSending           float64
	TimeReadingFirstBytes float64
	TimeReadingTotal      float64
	TimeTotal             float64
}

// NewClient function initializing a new client.
func (c *Client) NewClient(host string, port int, https bool) {
	c.Host = host
	c.Port = port
	c.Https = https
}

// Get function sending an HTTP GET request through the sockets
// and returning time results and HTTP code.
// This function manages different type of HTTP response (chunked or not).
func (c *Client) Get(route string, method string, params map[string]string) (Result, int) {
	// Generate a new socket.
	socket, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		panic(err)
	}
	c.Sockets = append(c.Sockets, &socket)
	
	if method != "GET" && method != "POST" {
		panic(fmt.Sprintf("Wrong method %s", method))
	}
	
	// Prepare parameters.
	data := url.Values{}
	for key, value := range params {
		data.Set(key, value)
	}
	
	// Prepare HTTP request.
	var url string
	if c.Https {
		url = fmt.Sprintf("https://%s%s", c.Host, route)
	} else {
		fmt.Sprintf("http://%s%s", c.Host, route)
	}

	req, err := http.NewRequest(method, url, bytes.NewBufferString(data.Encode()))
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpRequest(req, false)
	if err != nil {
		panic(err)
	}

	// Write/Read HTTP request/response.
	results := Result{}
	startSending := time.Now()
	fmt.Fprintf(socket, string(dump))
	results.TimeSending = time.Since(startSending).Seconds() * 1000
	startReading := time.Now()
	reader := bufio.NewReader(socket)
	headers := make(map[string]string)
	headers["HTTP"], err = reader.ReadString('\n')
	header, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	for header != "\r\n" {
		parsed := strings.Split(header, ":")
		if len(parsed) == 2 {
			headers[parsed[0]] = strings.Trim(parsed[1], "\r\n ")
		}
		header, err = reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
	}
	results.TimeReadingFirstBytes = time.Since(startReading).Seconds() * 1000
	if val, ok := headers["Content-Length"]; ok {
		contentLength, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		for i := 0; i < contentLength; i++ {
			_, err = reader.ReadByte()
			if err != nil {
				panic(err)
			}
		}
	} else {
		if val, ok := headers["Transfer-Encoding"]; ok {
			if val == "chunked" {
				var data []byte
				hexSize, _, _ := reader.ReadLine()
				size, _ := strconv.ParseInt("0x"+string(hexSize), 0, 64)
				for size > 0 {
					for i := 0; i < int(size); i++ {
						bt, err := reader.ReadByte()
						if err != nil {
							panic(err)
						}
						data = append(data, bt)
					}
					hexSize, _, _ := reader.ReadLine()
					hexSize, _, _ = reader.ReadLine()
					size, _ = strconv.ParseInt("0x"+string(hexSize), 0, 64)
				}
			}
		}
	}
	results.TimeReadingTotal = time.Since(startReading).Seconds() * 1000
	results.TimeTotal = time.Since(startSending).Seconds() * 1000

	// Parse code.
	re := regexp.MustCompile("HTTP/1.[0-1] ([0-9]{3}).*")
	submatches := re.FindStringSubmatch(headers["HTTP"])
	if len(submatches) == 0 {
		fmt.Println(headers)
		panic("Can't find HTTP response code")
	}
	code, err := strconv.Atoi(submatches[1])
	if err != nil {
		panic(err)
	}

	return results, code
}

// Close function to close
// all sockets.
func (c *Client) Close() {
	for i, _ := range c.Sockets {
		(*c.Sockets[i]).Close()
	}
}
