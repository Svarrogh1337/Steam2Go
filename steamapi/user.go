package steamapi

import (
	"context"
	"fmt"
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
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}
