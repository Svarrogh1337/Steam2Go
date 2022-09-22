package Steam2Go

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetCMListForConnectV1(options *GetCMListForConnectOptions) (*GetCMListForConnectResponseV1, error) {
	const (
		apiVersion = 1
	)
	resp := GetCMListForConnectResponseV1{}
	cellid := 1
	cmtype := ""
	realm := ""
	maxcount := 5
	if options != nil {
		cellid = options.Cellid
		cmtype = options.Cmtype
		realm = options.Realm
		maxcount = options.Maxcount
	}
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/ISteamDirectory/GetCMListForConnect/v%d?cellid=%d&cmtype=%s&realm=%s&maxcount=%d",
			c.BaseURL, apiVersion, cellid, cmtype, realm, maxcount), nil)
	log.Printf("%s/ISteamDirectory/GetCMListForConnect/v%d?cellid=%d&cmtype=%s&realm=%s&maxcount=%d",
		c.BaseURL, apiVersion, cellid, cmtype, realm, maxcount)
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
