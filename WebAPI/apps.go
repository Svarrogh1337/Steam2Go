package Steam2Go

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetAppListResponseV1 struct {
	Applist ApplistV1 `json:"applist"`
}

type ApplistV1 struct {
	Apps AppsV1 `json:"apps"`
}

type AppsV1 struct {
	App []AppV1 `json:"app"`
}

type AppV1 struct {
	Appid int    `json:"appid"`
	Name  string `json:"name"`
}

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
