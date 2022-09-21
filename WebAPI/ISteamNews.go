package Steam2Go

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type GetAppNewsResponse struct {
	Appnews AppNews `json:"appnews"`
}

type AppNews struct {
	Appid     int         `json:"appid"`
	Newsitems []NewsItems `json:"newsitems"`
	Count     int         `json:"count"`
}

type NewsItems struct {
	Gid           string   `json:"gid"`
	Title         string   `json:"title"`
	URL           string   `json:"url"`
	IsExternalURL bool     `json:"is_external_url"`
	Author        string   `json:"author"`
	Contents      string   `json:"contents"`
	Feedlabel     string   `json:"feedlabel"`
	Date          int      `json:"date"`
	Feedname      string   `json:"feedname"`
	FeedType      int      `json:"feed_type"`
	Appid         int      `json:"appid"`
	Tags          []string `json:"tags,omitempty"`
}
type GetAppNewsOptions struct {
	Maxlength int    `json:"Maxlength"`
	Enddate   int64  `json:"Enddate"`
	Count     int    `json:"count"`
	Tags      string `json:"tags,omitempty"`
}

func (c *Client) GetAppNews(ctx context.Context, appId int, options *GetAppNewsOptions) (*AppNews, error) {
	maxlength := 20
	enddate := time.Now().Unix()
	count := 20
	tags := ""
	if options != nil {
		maxlength = options.Maxlength
		enddate = options.Enddate
		count = options.Count
		tags = options.Tags
	}
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/ISteamNews/GetNewsForApp/v%d?appid=%d&maxlength=%d&enddate=%d&count=%d&tags=%s",
			c.BaseURL, c.ApiVersion, appId, maxlength, enddate, count, tags), nil)
	if err != nil {
		return nil, err
	}
	res := AppNews{}
	if err := c.sendRequest(ctx, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
