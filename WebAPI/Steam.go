package Steam2Go

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
	ApiVersion int
	HTTPClient *http.Client
}

func ApiClient(appKey string, apiVersion int) *Client {
	return &Client{
		BaseURL:    BaseURLV1,
		apiKey:     appKey,
		ApiVersion: apiVersion,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
