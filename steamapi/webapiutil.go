package steamapi

import (
	"context"
	"fmt"
)

type GetServerInfoResponseV1 struct {
	Servertime       int    `json:"servertime"`
	Servertimestring string `json:"servertimestring"`
}

func (c *Client) GetServerInfoV1(ctx context.Context) (*GetServerInfoResponseV1, error) {
	const (
		apiVersion = 1
	)
	resp := GetServerInfoResponseV1{}
	apiURL := fmt.Sprintf("%s/ISteamWebAPIUtil/GetServerInfo/v%d?key=%s", c.baseURL, apiVersion, c.apiKey)
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}
