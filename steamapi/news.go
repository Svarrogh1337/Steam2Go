package steamapi

import (
	"context"
	"fmt"
)

type GetAppNewsResponseV1 struct {
	Appnews AppNewsV1 `json:"appnews"`
}

type AppNewsV1 struct {
	Appid     int         `json:"appid"`
	Newsitems NewsItemsV1 `json:"newsitems"`
	Count     int         `json:"count"`
}

type NewsItemsV1 struct {
	Newsitem []NewsItem `json:"newsitem"`
}

type NewsItem struct {
	Gid           string `json:"gid"`
	Title         string `json:"title"`
	URL           string `json:"url"`
	IsExternalURL bool   `json:"is_external_url"`
	Author        string `json:"author"`
	Contents      string `json:"contents"`
	Feedlabel     string `json:"feedlabel"`
	Date          int    `json:"date"`
	Feedname      string `json:"feedname"`
	FeedType      int    `json:"feed_type"`
	Appid         int    `json:"appid"`
}

type GetAppNewsResponseV2 struct {
	Appnews AppNewsV2 `json:"appnews"`
}

type AppNewsV2 struct {
	Appid     int           `json:"appid"`
	Newsitems []NewsItemsV2 `json:"newsitems"`
	Count     int           `json:"count"`
}

type NewsItemsV2 struct {
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

func (c *Client) GetAppNewsV1(ctx context.Context, appid int, options ...RequestParameter) (*GetAppNewsResponseV1, error) {
	const (
		apiVersion = 1
	)
	resp := GetAppNewsResponseV1{}
	apiURL := fmt.Sprintf("%s/ISteamNews/GetNewsForApp/v%d?key=%s&appid=%d", c.baseURL, apiVersion, c.apiKey, appid)
	if params := getOptionalParameters(options...).urlParams.Encode(); params != "" {
		apiURL += "&" + params
	}
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}

func (c *Client) GetAppNewsV2(ctx context.Context, appid int, options ...RequestParameter) (*GetAppNewsResponseV2, error) {
	const (
		apiVersion = 2
	)
	resp := GetAppNewsResponseV2{}
	apiURL := fmt.Sprintf("%s/ISteamNews/GetNewsForApp/v%d?key=%s&appid=%d", c.baseURL, apiVersion, c.apiKey, appid)
	if params := getOptionalParameters(options...).urlParams.Encode(); params != "" {
		apiURL += "&" + params
	}
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}
