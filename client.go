package http

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	Socket *net.Conn
	Host   string
	Port   int
}

type Result struct {
	TimeSending           float64
	TimeReadingFirstBytes float64
	TimeReadingTotal      float64
	TimeTotal             float64
}

func (c *Client) Open(host string, port int) {
	socket, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}

	c.Socket = &socket
	c.Host = host
	c.Port = port
}

func (c *Client) Get(route string) (Result, int) {
	if c.Socket == nil {
		panic("Socket not opened")
	}

	// Prepare HTTP request.
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s%s", c.Host, route), nil)
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
	fmt.Fprintf((*c.Socket), string(dump))
	results.TimeSending = time.Since(startSending).Seconds() * 1000
	startReading := time.Now()
	reader := bufio.NewReader((*c.Socket))
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
		panic("Can't find HTTP response code")
	}
	code, err := strconv.Atoi(submatches[1])
	if err != nil {
		panic(err)
	}

	return results, code
}

func (c *Client) Close() {
	(*c.Socket).Close()
}
