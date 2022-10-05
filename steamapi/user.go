package steamapi

import (
	"context"
	"fmt"
	"strings"
)

// GetFriendListResponseV1 stores the data from GetFriendListV1 ( version 1 ) method  that is part of ISteamUser interface.
type GetFriendListResponseV1 struct {
	Friendslist struct {
		Friends []struct {
			Steamid      string `json:"steamid"`
			Relationship string `json:"relationship"`
			FriendSince  int    `json:"friend_since"`
		} `json:"friends"`
	} `json:"friendslist"`
}

// GetPlayerBansResponseV1 stores the data from GetPlayerBansV1 ( version 1 ) method  that is part of ISteamUser interface.
type GetPlayerBansResponseV1 struct {
	Players []struct {
		SteamID          string `json:"SteamId"`
		CommunityBanned  bool   `json:"CommunityBanned"`
		VACBanned        bool   `json:"VACBanned"`
		NumberOfVACBans  int    `json:"NumberOfVACBans"`
		DaysSinceLastBan int    `json:"DaysSinceLastBan"`
		NumberOfGameBans int    `json:"NumberOfGameBans"`
		EconomyBan       string `json:"EconomyBan"`
	} `json:"players,omitempty"`
}

// GetFriendListV1 gets the specified user friend list.
// This call has a mandatory parameter - steamid.
// This call accepts additional parameter - Relationship.
// Response - GetFriendListResponseV1
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

// GetPlayerBansV1 gets the account ban information for specified steamids.
// This call has a mandatory parameter - steamids, passed as slice.
// This call has no additional parameters.
// Response - GetPlayerBansResponseV1
func (c *Client) GetPlayerBansV1(ctx context.Context, steamids []string) (*GetPlayerBansResponseV1, error) {
	const (
		apiVersion = 1
	)
	resp := GetPlayerBansResponseV1{}
	apiURL := fmt.Sprintf("%s/ISteamUser/GetPlayerBans/v%d?key=%s&steamids=%s", c.baseURL, apiVersion, c.apiKey, strings.Join(steamids[:], ","))
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}

// GetPlayerSummariesV1 gets the account ban information for specified steamids.
// This call has a mandatory parameter - steamids, passed as slice.
// This call has no additional parameters.
// Response - GetPlayerBansResponseV1
func (c *Client) GetPlayerSummariesV1(ctx context.Context, steamids []string) (*GetPlayerBansResponseV1, error) {
	const (
		apiVersion = 1
	)
	resp := GetPlayerBansResponseV1{}
	apiURL := fmt.Sprintf("%s/ISteamUser/GetPlayerBans/v%d?key=%s&steamids=%s", c.baseURL, apiVersion, c.apiKey, strings.Join(steamids[:], ","))
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}
