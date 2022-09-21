package Steam2Go

import (
	"context"
	Steam2Go "github.com/Svarrogh1337/Steam2Go/WebAPI"
	"log"
	"testing"
	"time"
)

var options = &Steam2Go.GetISteamNewsOptions{
	5,
	time.Now().Unix(),
	5,
	"patchnotes",
}

var options2 = &Steam2Go.GetISteamNewsOptions{
	Count:   5,
	Enddate: time.Now().Unix(),
}

func TestApiClient(t *testing.T) {
	ctx := context.Background()
	c := Steam2Go.ApiClient("")
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
