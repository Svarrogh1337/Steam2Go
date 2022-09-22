package Steam2Go

import (
	"net/http"
	"time"
)

const (
	BaseURL = "https://api.steampowered.com"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func ApiClient(appKey string) *Client {
	return &Client{
		BaseURL: BaseURL,
		apiKey:  appKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
func (c *Client) sendRequest(req *http.Request) (*http.Response, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return nil, err
	}
	return res, nil
}
