package Steam2Go

import (
	"SteamGo/WebAPI"
	"context"
	"log"
	"testing"
	"time"
)

var options = &SteamGo.GetISteamNewsOptions{
	5,
	time.Now().Unix(),
	5,
	"patchnotes",
}

var options2 = &SteamGo.GetISteamNewsOptions{
	Count:   5,
	Enddate: time.Now().Unix(),
}

func TestApiClient(t *testing.T) {
	ctx := context.Background()
	c := SteamGo.ApiClient("")
	data, _ := c.GetISteamNews(ctx, 730, nil)
	log.Println(data.Appnews.Newsitems[5].Title)
	data2, _ := c.GetISteamNews(ctx, 440, options)
	log.Println(data2.Appnews.Newsitems[1].Title)
	data3, _ := c.GetISteamNews(ctx, 440, options2)
	if len(data3.Appnews.Newsitems) == options2.Count {
		log.Println("pass")
	}
	log.Println(data3.Appnews.Newsitems[1].Title)
}
