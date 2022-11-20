package a2s

import (
	"bytes"
	"log"
	"testing"
)

func TestSteamAPI(t *testing.T) {
	c, _ := NewClient("51.89.54.207", 27015)
	msg := struct {
		bytes.Buffer
	}{}
	msg.Write([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x54})
	msg.WriteString("Source Engine Query")
	msg.WriteByte(0)
	c.send(msg.Bytes())
	c.read()
	log.Println(string(c.buffer))
}
