package a2s

import (
	"fmt"
	"net"
)

type Client struct {
	addr string
	port int
	conn net.Conn
	size int
}

type Request struct {
	msg      []byte
	response *Response
}

func NewClient(addr string, port int) (*Client, error) {
	c := &Client{
		addr: addr,
		port: port,
		size: 1400,
	}
	address := fmt.Sprintf("%s:%d", c.addr, c.port)
	s, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		return nil, fmt.Errorf("Steam2Go A2S: Error %s", err)
	}
	c.conn, err = net.DialUDP("udp", nil, s)
	return c, nil
}

func (c *Client) send(request *Request) error {
	_, err := c.conn.Write(request.msg)
	if err != nil {
		return fmt.Errorf("Steam2Go A2S: Error %s", err)
	}
	return nil
}

func (c *Client) read(request *Request) error {
	request.response = &Response{
		raw: make([]byte, c.size),
	}
	size, err := c.conn.Read(request.response.raw)
	if size <= 0 {
		return fmt.Errorf("Steam2Go A2S: Packet size 0")
	}
	if err != nil {
		return fmt.Errorf("Steam2Go A2S: Error %s", err)
	}
	return nil
}
func (c *Client) Do(request *Request) (*Response, error) {
	err := c.send(request)
	if err != nil {
		return nil, err
	}
	err = c.read(request)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return request.response, nil
}
