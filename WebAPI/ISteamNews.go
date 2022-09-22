package Steam2Go

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func (c *Client) GetAppNewsV2(appId int, options *GetAppNewsOptions) (*GetAppNewsResponseV2, error) {
	const (
		apiVersion = 2
	)
	resp := GetAppNewsResponseV2{}
	maxlength := 20
	enddate := time.Now().Unix()
	count := 20
	tags := ""
	if options != nil {
		maxlength = options.Maxlength
		enddate = options.Enddate
		count = options.Count
		tags = options.Tags
	}
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/ISteamNews/GetNewsForApp/v%d?appid=%d&maxlength=%d&enddate=%d&count=%d&tags=%s",
			c.BaseURL, apiVersion, appId, maxlength, enddate, count, tags), nil)
	if err != nil {
		return nil, err
	}
	res, err := c.sendRequest(req)
	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		log.Println(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	return &resp, err
}

func (c *Client) GetAppNewsV1(appId int, options *GetAppNewsOptions) (*GetAppNewsResponseV1, error) {
	const (
		apiVersion = 1
	)
	resp := GetAppNewsResponseV1{}
	maxlength := 20
	enddate := time.Now().Unix()
	count := 20
	tags := ""
	if options != nil {
		maxlength = options.Maxlength
		enddate = options.Enddate
		count = options.Count
		tags = options.Tags
	}
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/ISteamNews/GetNewsForApp/v%d?appid=%d&maxlength=%d&enddate=%d&count=%d&tags=%s",
			c.BaseURL, apiVersion, appId, maxlength, enddate, count, tags), nil)
	if err != nil {
		return nil, err
	}
	res, err := c.sendRequest(req)
	if err = json.NewDecoder(res.Body).Decode(&resp); err != nil {
		log.Println(err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	return &resp, err
}
