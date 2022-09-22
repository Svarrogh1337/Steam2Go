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

func TestApiClient(t *testing.T) {
	c := Steam2Go.ApiClient("")
	newsV2, err := c.GetAppNewsV2(730, options)
	newsV1, err2 := c.GetAppNewsV1(730, options2)
	appsv1, err3 := c.GetAppListV1()
	assert.NotNil(t, newsV1)
	assert.NotNil(t, newsV2)
	assert.NotNil(t, appsv1.Applist.Apps.App[0].Appid)
	assert.Nil(t, err)
	assert.Nil(t, err2)
	assert.Nil(t, err3)
	assert.Equal(t, 730, newsV2.Appnews.Newsitems[0].Appid)
	assert.Equal(t, 730, newsV1.Appnews.Newsitems.Newsitem[0].Appid)
}
