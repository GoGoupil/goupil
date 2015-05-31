package http

import (
	"net"
	"fmt"
	"io/ioutil"
)

type Client struct {
	Socket *net.Conn
}

func (c *Client) Open(host string, port int) {
	socket, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}
	
	c.Socket = &socket
}

func (c *Client) Get(route string) {
	if c.Socket == nil {
		panic("Socket not opened")
	}
	
	(*c.Socket).Write([]byte(fmt.Sprintf("GET %s HTTP/1.0\r\n\r\n")))
	_, err := ioutil.ReadAll(*c.Socket)
	if err != nil {
		panic(err)
	}
}

func (c *Client) Close() {
	(*c.Socket).Close()
}
