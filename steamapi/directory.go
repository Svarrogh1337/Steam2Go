package steamapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetCMListForConnectResponseV1 struct {
	Response ResponseV1 `json:"response"`
}

type ResponseV1 struct {
	Serverlist []ServerlistV1 `json:"serverlist"`
	Success    bool           `json:"success"`
	Message    string         `json:"message"`
}

type ServerlistV1 struct {
	Endpoint       string  `json:"endpoint"`
	LegacyEndpoint string  `json:"legacy_endpoint"`
	Type           string  `json:"type"`
	Dc             string  `json:"dc"`
	Realm          string  `json:"realm"`
	Load           int     `json:"load"`
	WtdLoad        float64 `json:"wtd_load"`
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
