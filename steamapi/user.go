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

// GetPlayerSummariesResponseV1 stores the data from GetPlayerSummariesV1 ( version 1 ) method  that is part of ISteamUser interface.
type GetPlayerSummariesResponseV1 struct {
	Response struct {
		Players struct {
			Player []struct {
				Steamid                  string `json:"steamid"`
				Communityvisibilitystate int    `json:"communityvisibilitystate"`
				Profilestate             int    `json:"profilestate"`
				Personaname              string `json:"personaname"`
				Commentpermission        int    `json:"commentpermission"`
				Profileurl               string `json:"profileurl"`
				Avatar                   string `json:"avatar"`
				Avatarmedium             string `json:"avatarmedium"`
				Avatarfull               string `json:"avatarfull"`
				Avatarhash               string `json:"avatarhash"`
				Lastlogoff               int    `json:"lastlogoff"`
				Personastate             int    `json:"personastate"`
				Primaryclanid            string `json:"primaryclanid"`
				Timecreated              int    `json:"timecreated"`
				Personastateflags        int    `json:"personastateflags"`
				Loccountrycode           string `json:"loccountrycode"`
				Locstatecode             string `json:"locstatecode"`
				Loccityid                int    `json:"loccityid"`
			} `json:"player"`
		} `json:"players"`
	} `json:"response"`
}

// GetPlayerSummariesResponseV2 stores the data from GetPlayerSummariesV2 ( version 2 ) method  that is part of ISteamUser interface.
type GetPlayerSummariesResponseV2 struct {
	Response struct {
		Players []struct {
			Steamid                  string `json:"steamid"`
			Communityvisibilitystate int    `json:"communityvisibilitystate"`
			Profilestate             int    `json:"profilestate"`
			Personaname              string `json:"personaname"`
			Commentpermission        int    `json:"commentpermission"`
			Profileurl               string `json:"profileurl"`
			Avatar                   string `json:"avatar"`
			Avatarmedium             string `json:"avatarmedium"`
			Avatarfull               string `json:"avatarfull"`
			Avatarhash               string `json:"avatarhash"`
			Lastlogoff               int    `json:"lastlogoff"`
			Personastate             int    `json:"personastate"`
			Primaryclanid            string `json:"primaryclanid"`
			Timecreated              int    `json:"timecreated"`
			Personastateflags        int    `json:"personastateflags"`
			Loccountrycode           string `json:"loccountrycode"`
			Locstatecode             string `json:"locstatecode"`
			Loccityid                int    `json:"loccityid"`
		} `json:"players"`
	} `json:"response"`
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
// This call is no longer officially supported. It will continue to be usable, but it's highly recommended that you use the latest version.
// Response - GetPlayerBansResponseV1
func (c *Client) GetPlayerSummariesV1(ctx context.Context, steamids []string) (*GetPlayerSummariesResponseV1, error) {
	const (
		apiVersion = 1
	)
	resp := GetPlayerSummariesResponseV1{}
	apiURL := fmt.Sprintf("%s/ISteamUser/GetPlayerSummaries/v%d?key=%s&steamids=%s", c.baseURL, apiVersion, c.apiKey, strings.Join(steamids[:], ","))
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}

// GetPlayerSummariesV2 gets the account ban information for specified steamids.
// This call has a mandatory parameter - steamids, passed as slice.
// This call has no additional parameters.
// This call has previous versions which are no longer officially supported. They will continue to be usable but it's highly recommended that you use the latest version.
// Change history:
// Version 2 - Removes element names from arrays
// Response - GetPlayerSummariesResponseV2
func (c *Client) GetPlayerSummariesV2(ctx context.Context, steamids []string) (*GetPlayerSummariesResponseV2, error) {
	const (
		apiVersion = 2
	)
	resp := GetPlayerSummariesResponseV2{}
	apiURL := fmt.Sprintf("%s/ISteamUser/GetPlayerSummaries/v%d?key=%s&steamids=%s", c.baseURL, apiVersion, c.apiKey, strings.Join(steamids[:], ","))
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}
