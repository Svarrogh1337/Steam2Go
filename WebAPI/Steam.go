package SteamGo

import (
	"net/http"
	"time"
)

const (
	BaseURLV1 = "https://api.steampowered.com"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func ApiClient(appKey string) *Client {
	return &Client{
		BaseURL: BaseURLV1,
		apiKey:  appKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
