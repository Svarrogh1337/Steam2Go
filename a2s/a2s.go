package a2s

import (
	"bytes"
	"fmt"
	"net"
)

const (
	nonSplit      = 0xFFFFFFFF
	split         = 0xFFFFFFFE
	gameID        = 0x01
	steamID       = 0x10
	keywords      = 0x20
	sourceTV      = 0x40
	challenge     = 0x41
	player        = 0x44
	rules         = 0x45
	info          = 0x49
	infoRequest   = 0x54
	playerRequest = 0x55
	rulesRequest  = 0x56
	port          = 0x80
	theShip       = 2400
)

type Client struct {
	addr      string
	port      int
	conn      net.Conn
	size      int
	challenge []byte
}

type Request struct {
	msg         bytes.Buffer
	response    *Response
	requestType int
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
	if c.challenge != nil {
		if request.requestType == playerRequest || request.requestType == rulesRequest {
			request.msg.Truncate(5)
		}
		request.msg.Write(c.challenge)
	}
	_, err := c.conn.Write(request.msg.Bytes())
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
	switch PacketHeader := request.response.readUint32(); PacketHeader {
	case nonSplit:
	case split:
	}
	switch request.response.t = request.response.readUint8(); request.response.t {
	case challenge:
		c.challenge = request.response.readRawUint32()
		c.do(request)
	}
	request.response.position = 0
	if size <= 0 {
		return fmt.Errorf("Steam2Go A2S: Packet size 0")
	}
	if err != nil {
		return fmt.Errorf("Steam2Go A2S: Error %s", err)
	}
	return nil
}
func (c *Client) Do(request *Request) (*Response, error) {
	resp, err := c.do(request)
	return resp, err
}
func (c *Client) do(request *Request) (*Response, error) {
	err := c.send(request)
	if err != nil {
		return nil, err
	}
	err = c.read(request)
	if err != nil {
		return nil, err
	}
	return request.response, nil
}
