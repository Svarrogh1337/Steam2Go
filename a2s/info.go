package a2s

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
)

type InfoResponse struct {
	PacketHeader  uint32
	PayloadHeader byte
	Protocol      byte
	Name          string
	Map           string
	Folder        string
	Game          string
	ID            int32
	Players       byte
	MaxPlayers    byte
	Bots          byte
	ServerType    byte
	Environment   byte
	Visibility    byte
	VAC           byte
}

func (c *Client) Info() {
	var msg bytes.Buffer
	msg.Write([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x54})
	msg.WriteString("Source Engine Query")
	msg.WriteByte(0)
	c.send(msg.Bytes())
	c.read()
	resp := InfoResponse{}
	//resp.Header = binary.LittleEndian.Uint32(c.buffer[:4])
	resp.PayloadHeader = c.buffer[5]
	resp.Protocol = c.buffer[6]
	pos := 6
	start := pos
	v := reflect.Indirect(reflect.ValueOf(resp))
	for i := 0; i < v.NumField(); i++ {
		if v.Type().Field(i).Type.String() == "string" {
			fmt.Println(v.Type().Field(i).Name)
		}
	}
	for {
		if c.buffer[pos] == 0 {
			pos++
			break
		}
		pos++
	}
	log.Println(string(c.buffer[start : pos-1]))
}

/*func (c *Client) ReadString(start int) int {
	for {
		pos := start
		if c.buffer[pos] == 0 {
			pos++
			return 1
		}
		pos++
		return ReadString(pos - 1)
	}
}*/
