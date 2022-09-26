package steamapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNews(t *testing.T) {
	c := NewClient("")
	newsV2, err := c.GetAppNewsV2(context.Background(), 730, Maxlength(5), Feeds("SteamDB"))
	newsV1, err2 := c.GetAppNewsV1(context.Background(), 730, Count(3), Enddate(time.Now().Unix()), Tags("patchnotes"))
	assert.NotNil(t, newsV1)
	assert.NotNil(t, newsV2)
	assert.Equal(t, 730, newsV2.Appnews.Newsitems[0].Appid)
	assert.Equal(t, 730, newsV2.Appnews.Newsitems[0].Appid)
	assert.Equal(t, 730, newsV1.Appnews.Newsitems.Newsitem[0].Appid)
	assert.Nil(t, err)
	assert.Nil(t, err2)
}
