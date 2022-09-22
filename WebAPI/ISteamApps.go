package Steam2Go

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetAppListResponse struct {
	Applist Applist `json:"applist"`
}

type Applist struct {
	Apps App `json:"apps"`
}

type Apps struct {
	App []App `json:"app"`
}

type App struct {
	Appid int    `json:"appid"`
	Name  string `json:"name"`
}

func (c *Client) GetAppList(apiVersion int) (*GetAppListResponse, error) {
	var fullResponse GetAppListResponse
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/ISteamApps/GetAppList/v%d",
			c.BaseURL, apiVersion), nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return nil, err
	}
	if err := json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return nil, err
	}
	return &fullResponse, err
}
