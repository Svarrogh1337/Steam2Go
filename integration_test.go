package Steam2Go

import (
	Steam2Go "github.com/Svarrogh1337/Steam2Go/WebAPI"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var options = &Steam2Go.GetAppNewsOptions{
	Maxlength: 5,
	Enddate:   time.Now().Unix(),
	Count:     5,
	Tags:      "patchnotes",
}

var options2 = &Steam2Go.GetAppNewsOptions{
	Count:   5,
	Enddate: time.Now().Unix(),
}

// If Maxcount = 0 ( default value ) SteamAPI returns either 500 or empty json
var options3 = &Steam2Go.GetCMListForConnectOptions{
	Cellid:   0,
	Maxcount: 5,
}

func TestApiClient(t *testing.T) {
	c := Steam2Go.ApiClient("")
	newsV2, err := c.GetAppNewsV2(730, options)
	newsV1, err2 := c.GetAppNewsV1(730, options2)
	appsv1, err3 := c.GetAppListV1()
	cmlist, err4 := c.GetCMListForConnectV1(options3)
	assert.NotNil(t, newsV1)
	assert.NotNil(t, newsV2)
	assert.NotNil(t, appsv1.Applist.Apps.App[0].Appid)
	assert.NotNil(t, cmlist)
	assert.Nil(t, err)
	assert.Nil(t, err2)
	assert.Nil(t, err3)
	assert.Nil(t, err4)
	assert.Equal(t, 730, newsV2.Appnews.Newsitems[0].Appid)
	assert.Equal(t, 730, newsV1.Appnews.Newsitems.Newsitem[0].Appid)
	assert.True(t, true, cmlist.Response.Success)
}
