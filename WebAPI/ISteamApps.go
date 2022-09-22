package Steam2Go

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetAppListV1() (*GetAppListResponseV1, error) {
	const (
		apiVersion = 1
	)
	resp := GetAppListResponseV1{}
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/ISteamApps/GetAppList/v%d",
			c.BaseURL, apiVersion), nil)
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
