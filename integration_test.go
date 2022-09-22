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
	data, err := c.GetAppNewsV2(730, options)
	data2, err2 := c.GetAppNewsV1(730, options2)
	assert.NotNil(t, data)
	assert.NotNil(t, data2)
	assert.Nil(t, err)
	assert.Nil(t, err2)
	assert.Equal(t, 730, data.Appnews.Newsitems[1].Appid)
	assert.Equal(t, 730, data2.Appnews.Newsitems.Newsitem[1].Appid)
}
