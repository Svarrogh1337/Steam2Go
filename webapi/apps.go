package webapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.sendRequest(req)
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
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
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.sendRequest(req)
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	return &resp, err
}
