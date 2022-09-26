package steamapi

import (
	"context"
	"fmt"
)

type GetAppListResponseV1 struct {
	Applist struct {
		Apps struct {
			App []struct {
				Appid int    `json:"appid"`
				Name  string `json:"name"`
			} `json:"app"`
		} `json:"apps"`
	} `json:"applist"`
}

type GetAppListResponseV2 struct {
	Applist struct {
		Apps []struct {
			Appid int    `json:"appid"`
			Name  string `json:"name"`
		} `json:"apps"`
	} `json:"applist"`
}

// GetAppListV1 gets the complete list of public apps.
// This call has no additional parameters.
// This version is no longer officially supported. It will continue to be usable, but it's highly recommended that you use the latest version.
// Response - GetAppListResponseV1
func (c *Client) GetAppListV1(ctx context.Context) (*GetAppListResponseV1, error) {
	const (
		apiVersion = 1
	)
	resp := GetAppListResponseV1{}
	apiURL := fmt.Sprintf("%s/ISteamApps/GetAppList/v%d?key=%s", c.baseURL, apiVersion, c.apiKey)
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}

// GetAppListV2 gets the complete list of public apps.
// This call has no additional parameters.
// This method has previous versions which are no longer officially supported. They will continue to be usable, but it's highly recommended that you use the latest version.
// Change history:
// Version 2 - Removed the redundant "app" field.
// Response - GetAppListResponseV2
func (c *Client) GetAppListV2(ctx context.Context) (*GetAppListResponseV2, error) {
	const (
		apiVersion = 2
	)
	resp := GetAppListResponseV2{}
	apiURL := fmt.Sprintf("%s/ISteamApps/GetAppList/v%d?key=%s", c.baseURL, apiVersion, c.apiKey)
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}
