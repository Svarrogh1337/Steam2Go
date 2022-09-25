// Package Steam2Go.steamapi provides low level bindings for the
// Steamworks Web API interfaces .
package steamapi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	BaseURL = "https://api.steampowered.com"
)

type Client struct {
	baseURL string
	apiKey  string
	http    *http.Client
}

type requestParameters struct {
	urlParams url.Values
}

type RequestParameter func(*requestParameters)

func NewClient(apiKey string) *Client {
	return &Client{
		baseURL: BaseURL,
		apiKey:  apiKey,
		http: &http.Client{
			Timeout: time.Minute,
		},
	}
}
func (c *Client) sendRequest(req *http.Request) (*http.Response, error) {
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusForbidden {
		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("Steam2Go: Unknown error %s", res.Status)
		}
		return nil, fmt.Errorf("Steam2Go: Response %s: %s",
			res.Status, string(responseBody))
	}
	if res.StatusCode != http.StatusOK {
		return nil, err
	}
	return res, nil
}

func getOptionalParameters(options ...RequestParameter) requestParameters {
	o := requestParameters{
		urlParams: url.Values{},
	}
	for _, opt := range options {
		opt(&o)
	}
	return o
}

func Maxlength(amount int) RequestParameter {
	return func(o *requestParameters) {
		o.urlParams.Set("maxlength", strconv.Itoa(amount))
	}
}

func Count(amount int) RequestParameter {
	return func(o *requestParameters) {
		o.urlParams.Set("count", strconv.Itoa(amount))
	}
}
