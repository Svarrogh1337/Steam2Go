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
	assert.Nil(t, err)
	assert.NotNil(t, res.decoded.(InfoResponse).Name)
}

func TestShips(t *testing.T) {
	c, _ := NewClient("46.4.48.226", 27021)
	request := RequestInfo()
	res, err := c.Do(request)
	res.Decode()
	assert.Nil(t, err)
	log.Println(res.decoded.(InfoResponse).Name)
	log.Println(res.decoded.(InfoResponse).Map)
	log.Println(res.decoded.(InfoResponse).TheShipResponse.Mode)
	log.Println(res.decoded.(InfoResponse).Version)
}
