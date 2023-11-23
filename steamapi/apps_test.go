package steamapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApps(t *testing.T) {
	c := NewClient("")
	appsv1, err := c.GetAppListV1(context.Background())
	assert.NotNil(t, appsv1.Applist.Apps.App[0].Appid)
	assert.Nil(t, err)
	appsv2, err := c.GetAppListV2(context.Background())
	assert.NotNil(t, appsv2.Applist.Apps[0].Appid)
	assert.Nil(t, err)
	sdr1, err := c.GetSDRConfigV1(context.Background(), 730)
	assert.NotNil(t, sdr1.Pops.Ams.Relays[0].Ipv4)
	assert.Nil(t, err)
	server, err := c.GetServersAtAddressV1(context.Background(), "94.225.213.234:25560")
	assert.True(t, server.Response.Success)
	assert.Nil(t, err)
	server2, err := c.GetServersAtAddressV1(context.Background(), "94.225.213.234")
	assert.NotNil(t, server2.Response.Servers[0].Steamid)
	assert.Nil(t, err)
	up2date, err := c.UpToDateCheckV1(context.Background(), 730, 13843)
	assert.True(t, up2date.Response.Success)
	assert.Nil(t, err)
	up2date2, err := c.UpToDateCheckV1(context.Background(), 730, 1)
	assert.True(t, up2date2.Response.Success)
	assert.Nil(t, err)
}
