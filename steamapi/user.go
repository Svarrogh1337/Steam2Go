package steamapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetFriendListResponseV1 struct {
	Friendslist struct {
		Friends []struct {
			Steamid      string `json:"steamid"`
			Relationship string `json:"relationship"`
			FriendSince  int    `json:"friend_since"`
		} `json:"friends"`
	} `json:"friendslist"`
}

func (c *Client) GetFriendListV1(ctx context.Context, steamid int, options ...RequestParameter) (*GetFriendListResponseV1, error) {
	const (
		apiVersion = 1
	)
	resp := GetFriendListResponseV1{}
	apiURL := fmt.Sprintf("%s/ISteamUser/GetFriendList/v%d?key=%s&steamid=%d", c.baseURL, apiVersion, c.apiKey, steamid)
	if params := getOptionalParameters(options...).urlParams.Encode(); params != "" {
		apiURL += "&" + params
	}
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}
	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	return &resp, err
}
