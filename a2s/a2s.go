package a2s

import (
	"fmt"
	"net"
)

type Client struct {
	addr   string
	port   int
	conn   net.Conn
	buffer []byte
}

func NewClient(addr string, port int) (*Client, error) {
	c := &Client{
		addr: addr,
		port: port,
	}
	address := fmt.Sprintf("%s:%d", c.addr, c.port)
	s, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		return nil, fmt.Errorf("Steam2Go A2S: Error %s", err)
	}
	c.conn, err = net.DialUDP("udp", nil, s)
	return c, nil
}

func (c *Client) send(data []byte) error {
	_, err := c.conn.Write(data)
	if err != nil {
		return fmt.Errorf("Steam2Go A2S: Error %s", err)
	}
	return nil
}

func (c *Client) read() error {
	c.buffer = make([]byte, 1400)
	size, err := c.conn.Read(c.buffer)
	if size <= 0 {
		return fmt.Errorf("Steam2Go A2S: Packet size 0")
	}
	if err != nil {
		return fmt.Errorf("Steam2Go A2S: Error %s", err)
	}
	return nil
}
