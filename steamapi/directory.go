package steamapi

import (
	"context"
	"fmt"
)

type GetCMListForConnectResponseV1 struct {
	Response struct {
		Serverlist []struct {
			Endpoint       string  `json:"endpoint"`
			LegacyEndpoint string  `json:"legacy_endpoint"`
			Type           string  `json:"type"`
			Dc             string  `json:"dc"`
			Realm          string  `json:"realm"`
			Load           int     `json:"load"`
			WtdLoad        float64 `json:"wtd_load"`
		} `json:"serverlist"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"response"`
}

func (c *Client) GetCMListForConnectV1(ctx context.Context, options ...RequestParameter) (*GetCMListForConnectResponseV1, error) {
	const (
		apiVersion = 1
	)
	resp := GetCMListForConnectResponseV1{}
	apiURL := fmt.Sprintf("%s/ISteamDirectory/GetCMListForConnect/v%d?key=%s", c.baseURL, apiVersion, c.apiKey)
	if params := getOptionalParameters(options...).urlParams.Encode(); params != "" {
		apiURL += "&" + params
	}
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}
