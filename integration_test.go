package Steam2Go

import (
	Steam2Go "github.com/Svarrogh1337/Steam2Go/WebAPI"
	"log"
	"testing"
	"time"
)

var options = &Steam2Go.GetAppNewsOptions{
	5,
	time.Now().Unix(),
	5,
	"patchnotes",
}

var options2 = &Steam2Go.GetAppNewsOptions{
	Count:   5,
	Enddate: time.Now().Unix(),
}

func TestApiClient(t *testing.T) {
	c := Steam2Go.ApiClient("", 1)
	data, _ := c.GetAppNews(730, nil)
	log.Println(data.Appnews.Newsitems[5].Title)
	data2, _ := c.GetAppNews(440, options)
	log.Println(data2.Appnews.Newsitems[1].Title)
	data3, _ := c.GetAppNews(440, options2)
	if len(data3.Appnews.Newsitems) == options2.Count {
		log.Println("pass")
	}
	log.Println(data3.Appnews.Newsitems[1].Title)
}
