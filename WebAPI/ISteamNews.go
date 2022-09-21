package Steam2Go

import (
	"encoding/json"
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

func (c *Client) GetISteamNews(appId int, options *GetAppNewsOptions) (*GetAppNewsResponse, error) {
	var fullResponse GetAppNewsResponse
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
		fmt.Sprintf("%s/ISteamNews/GetNewsForApp/v2?appid=%d&maxlength=%d&enddate=%d&count=%d&tags=%s",
			c.BaseURL, appId, maxlength, enddate, count, tags), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return nil, err
	}

	return &fullResponse, err
}
