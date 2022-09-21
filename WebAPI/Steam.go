package Steam2Go

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	BaseURLV1 = "https://api.steampowered.com"
)

type genericResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

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

func (c *Client) sendRequest(ctx context.Context, req *http.Request, data interface{}) error {
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return err
	}
	log.Println(res.Body)
	fullResponse := genericResponse{
		Data: data,
	}
	if err := json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}
	log.Println(fullResponse.Data)
	return nil
}
