package Steam2Go

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

// If Maxcount = 0 ( default value ) SteamAPI returns either 500 or empty json
var options3 = &GetCMListForConnectOptions{
	Cellid:   0,
	Maxcount: 5,
}

func TestApiClient(t *testing.T) {
	c := WebApiClient("")

	newsV2, err := c.GetAppNewsV2(context.Background(), 730, Maxlength(5))
	newsV1, err2 := c.GetAppNewsV1(context.Background(), 730)
	frsV1, err5 := c.GetFriendListV1(context.Background(), 76561198068163231)
	appsv1, err3 := c.GetAppListV1(context.Background())
	cmlist, err4 := c.GetCMListForConnectV1(context.Background())
	assert.NotNil(t, frsV1)
	assert.NotNil(t, newsV1)
	assert.NotNil(t, newsV2)
	assert.NotNil(t, appsv1.Applist.Apps.App[0].Appid)
	assert.NotNil(t, cmlist)
	assert.Nil(t, err)
	assert.Nil(t, err2)
	assert.Nil(t, err3)
	assert.Nil(t, err4)
	assert.Nil(t, err5)
	assert.Equal(t, 730, newsV2.Appnews.Newsitems[0].Appid)
	assert.Equal(t, 730, newsV2.Appnews.Newsitems[0].Appid)
	assert.Equal(t, 730, newsV1.Appnews.Newsitems.Newsitem[0].Appid)
	assert.True(t, true, cmlist.Response.Success)
}