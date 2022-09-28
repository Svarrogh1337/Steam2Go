// Package steamapi provides low level bindings for the
// Steamworks Web API interfaces .
package steamapi

import (
	"context"
	"encoding/json"
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

func (c *Client) get(ctx context.Context, apiURL string, v interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return err
	}
	res, err := c.http.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Steam2Go: Unknown error %s", res.Status)
		}
		return fmt.Errorf("Steam2Go: Response %s: %s",
			res.Status, string(responseBody))
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	return nil
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

func Maxcount(amount uint32) RequestParameter {
	return func(o *requestParameters) {
		o.urlParams.Set("maxcount", strconv.FormatUint(uint64(amount), 10))
	}
}
func Enddate(amount int64) RequestParameter {
	return func(o *requestParameters) {
		o.urlParams.Set("enddate", strconv.FormatUint(uint64(amount), 10))
	}
}
func Count(amount uint32) RequestParameter {
	return func(o *requestParameters) {
		o.urlParams.Set("count", strconv.FormatUint(uint64(amount), 10))
	}
}
func Cellid(amount uint32) RequestParameter {
	return func(o *requestParameters) {
		o.urlParams.Set("cellid", strconv.FormatUint(uint64(amount), 10))
	}
}

// Available only for ISteamNewsV2 API
func Feeds(amount string) RequestParameter {
	return func(o *requestParameters) {
		o.urlParams.Set("feeds", amount)
	}
}
func Tags(amount string) RequestParameter {
	return func(o *requestParameters) {
		o.urlParams.Set("Tags", amount)
	}
}

func Relationship(amount string) RequestParameter {
	return func(o *requestParameters) {
		o.urlParams.Set("relationship", amount)
	}
}
