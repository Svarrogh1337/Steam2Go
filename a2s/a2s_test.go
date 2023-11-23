package a2s

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*func TestCSGO(t *testing.T) {
	c, _ := NewClient("87.121.112.34", 27015)
	reqInfo := RequestInfo()
	resInfo, errInfo := c.Do(reqInfo)
	assert.Nil(t, errInfo)
	resInfo.Decode()
	assert.NotNil(t, resInfo.decoded.(InfoResponse).Version)
	reqPlayer := RequestPlayer()
	res, err := c.Do(reqPlayer)
	assert.Nil(t, err)
	res.Decode()
	assert.NotNil(t, res.decoded.(PlayerResponse).PlayerCount)
}



func TestValheim(t *testing.T) {
	c, _ := NewClient("alfheim.unovalegends.net", 2457)
	reqInfo := RequestInfo()
	resInfo, errInfo := c.Do(reqInfo)
	assert.Nil(t, errInfo)
	resInfo.Decode()
	assert.NotNil(t, resInfo.decoded.(InfoResponse).Version)
}*/

func TestShips(t *testing.T) {
	c, _ := NewClient("46.4.48.226", 27021)
	/*	reqInfo := RequestInfo()
		resInfo, errInfo := c.Do(reqInfo)
		assert.Nil(t, errInfo)
		resInfo.Decode()
		assert.NotNil(t, resInfo.decoded.(InfoResponse).Version)*/
	reqPlayer := RequestPlayer()
	res, err := c.Do(reqPlayer)
	assert.Nil(t, err)
	res.Decode()
}
