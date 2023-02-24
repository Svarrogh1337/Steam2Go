package a2s

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCSGO(t *testing.T) {
	c, _ := NewClient("51.89.54.207", 27015)
	request := RequestInfo()
	res, err := c.Do(request)
	res.Decode()
	log.Println(res.t)
	assert.Nil(t, err)
	assert.NotNil(t, res.decoded.(InfoResponse).Name)
	request2 := RequestPlayer()
	res2, _ := c.Do(request2)
	log.Println(res2.raw)
}

func TestShips(t *testing.T) {
	c, _ := NewClient("46.4.48.226", 27021)
	request := RequestInfo()
	res, err := c.Do(request)
	assert.Nil(t, err)
	err2 := res.Decode()
	assert.Nil(t, err2)
	log.Println(res.decoded.(InfoResponse).Name)
	log.Println(res.decoded.(InfoResponse).Map)
	log.Println(res.decoded.(InfoResponse).TheShipResponse.Mode)
	log.Println(res.decoded.(InfoResponse).Version)
	log.Println(res.decoded.(InfoResponse).PacketHeader)
	log.Println(res.decoded.(InfoResponse).PayloadHeader)
	log.Println(res.decoded.(InfoResponse).Protocol)
	request2 := RequestPlayer()
	res2, err2 := c.Do(request2)
	log.Println(res2, err2)
}

func TestValheim(t *testing.T) {
	c, _ := NewClient("alfheim.unovalegends.net", 2457)
	request := RequestInfo()
	res, err := c.Do(request)
	assert.Nil(t, err)
	err2 := res.Decode()
	assert.Nil(t, err2)
	log.Println(res.decoded.(InfoResponse).PacketHeader)
	log.Println(res.decoded.(InfoResponse).PayloadHeader)
	log.Println(res.decoded.(InfoResponse).Name)
	request2 := RequestPlayer()
	res2, err2 := c.Do(request2)
	log.Println(res2, err2)
	request3 := RequestPlayer()
	res3, err3 := c.Do(request3)
	log.Println(res3, err3)
	request4 := RequestPlayer()
	res4, err4 := c.Do(request4)
	log.Println(res4, err4)

}
