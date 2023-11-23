package steamapi

import (
	"context"
	"encoding/json"
	"fmt"
)

type GetAssetPricesResponseV1 struct {
	Result struct {
		Success bool `json:"success"`
		Assets  []struct {
			Prices struct {
				Usd     int `json:"USD"`
				Gbp     int `json:"GBP"`
				Eur     int `json:"EUR"`
				Rub     int `json:"RUB"`
				Brl     int `json:"BRL"`
				Unknown int `json:"Unknown"`
				Jpy     int `json:"JPY"`
				Nok     int `json:"NOK"`
				Idr     int `json:"IDR"`
				Myr     int `json:"MYR"`
				Php     int `json:"PHP"`
				Sgd     int `json:"SGD"`
				Thb     int `json:"THB"`
				Vnd     int `json:"VND"`
				Krw     int `json:"KRW"`
				Try     int `json:"TRY"`
				Uah     int `json:"UAH"`
				Mxn     int `json:"MXN"`
				Cad     int `json:"CAD"`
				Aud     int `json:"AUD"`
				Nzd     int `json:"NZD"`
				Pln     int `json:"PLN"`
				Chf     int `json:"CHF"`
				Aed     int `json:"AED"`
				Clp     int `json:"CLP"`
				Cny     int `json:"CNY"`
				Cop     int `json:"COP"`
				Pen     int `json:"PEN"`
				Sar     int `json:"SAR"`
				Twd     int `json:"TWD"`
				Hkd     int `json:"HKD"`
				Zar     int `json:"ZAR"`
				Inr     int `json:"INR"`
				Ars     int `json:"ARS"`
				Crc     int `json:"CRC"`
				Ils     int `json:"ILS"`
				Kwd     int `json:"KWD"`
				Qar     int `json:"QAR"`
				Uyu     int `json:"UYU"`
				Kzt     int `json:"KZT"`
				Byn     int `json:"BYN"`
			} `json:"prices"`
			Name  string `json:"name"`
			Date  string `json:"date"`
			Class []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"class"`
			Classid string `json:"classid"`
		} `json:"assets"`
	} `json:"result"`
}

type GetAssetClassInfoResponseRawV1 struct {
	Result map[string]json.RawMessage `json:"result"`
}
type GetAssetClassInfoResponseV1 struct {
	IconURL                   string       `json:"icon_url,omitempty"`
	IconURLLarge              string       `json:"icon_url_large,omitempty"`
	IconDragURL               string       `json:"icon_drag_url,omitempty"`
	Name                      string       `json:"name,omitempty"`
	MarketHashName            string       `json:"market_hash_name,omitempty"`
	MarketName                string       `json:"market_name,omitempty"`
	NameColor                 string       `json:"name_color,omitempty"`
	BackgroundColor           string       `json:"background_color,omitempty"`
	Type                      string       `json:"type,omitempty"`
	Tradable                  string       `json:"tradable,omitempty"`
	Marketable                string       `json:"marketable,omitempty"`
	Commodity                 string       `json:"commodity,omitempty"`
	MarketTradableRestriction string       `json:"market_tradable_restriction,omitempty"`
	Fraudwarnings             string       `json:"fraudwarnings,omitempty"`
	Descriptions              Descriptions `json:"descriptions,omitempty"`
	OwnerDescriptions         string       `json:"owner_descriptions,omitempty"`
	Tags                      Tag          `json:"tags,omitempty"`
	Classid                   string       `json:"classid,omitempty"`
}
type Num0 struct {
	Type    string `json:"type,omitempty"`
	Value   string `json:"value,omitempty"`
	AppData string `json:"app_data,omitempty"`
}
type Num1 struct {
	Type    string `json:"type,omitempty"`
	Value   string `json:"value,omitempty"`
	AppData string `json:"app_data,omitempty"`
}
type AppData struct {
	Limited string `json:"limited,omitempty"`
}
type Num2 struct {
	Type    string  `json:"type,omitempty"`
	Value   string  `json:"value,omitempty"`
	Color   string  `json:"color,omitempty"`
	AppData AppData `json:"app_data,omitempty"`
}
type Descriptions struct {
	Num0 Num0 `json:"0,omitempty"`
	Num1 Num1 `json:"1,omitempty"`
	Num2 Num2 `json:"2,omitempty"`
}
type Num10 struct {
	InternalName string `json:"internal_name,omitempty"`
	Name         string `json:"name,omitempty"`
	Category     string `json:"category,omitempty"`
	CategoryName string `json:"category_name,omitempty"`
}
type Num11 struct {
	InternalName string `json:"internal_name,omitempty"`
	Name         string `json:"name,omitempty"`
	Category     string `json:"category,omitempty"`
	CategoryName string `json:"category_name,omitempty"`
}
type Num12 struct {
	InternalName string `json:"internal_name,omitempty"`
	Name         string `json:"name,omitempty"`
	Category     string `json:"category,omitempty"`
	Color        string `json:"color,omitempty"`
	CategoryName string `json:"category_name,omitempty"`
}
type Tag struct {
	Num0 Num10 `json:"0,omitempty"`
	Num1 Num11 `json:"1,omitempty"`
	Num2 Num12 `json:"2,omitempty"`
}

func (c *Client) GetAssetPricesV1(ctx context.Context, appid int, options ...RequestParameter) (*GetAssetPricesResponseV1, error) {
	const (
		iface      = "ISteamEconomy"
		apiVersion = 1
		method     = "GetAssetPrices"
	)
	resp := GetAssetPricesResponseV1{}
	apiURL := fmt.Sprintf("%s/%s/%s/v%d?key=%s&appid=%d", c.baseURL, iface, method, apiVersion, c.apiKey, appid)
	if params := getOptionalParameters(options...).urlParams.Encode(); params != "" {
		apiURL += "&" + params
	}
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}

func (c *Client) GetAssetClassInfoV1(ctx context.Context, appid int, classid int) (*GetAssetClassInfoResponseV1, error) {
	const (
		iface      = "ISteamEconomy"
		apiVersion = 1
		method     = "GetAssetClassInfo"
	)
	apiURL := fmt.Sprintf("%s/%s/%s/v%d?key=%s&appid=%d&class_count=1&classid0=%d", c.baseURL, iface, method, apiVersion, c.apiKey, appid, classid)
	raw := GetAssetClassInfoResponseRawV1{}
	resp := GetAssetClassInfoResponseV1{}
	err := c.get(ctx, apiURL, &raw)
	for _, object := range raw.Result {
		err := json.Unmarshal(object, &resp)
		if err != nil {
			continue
		}
	}
	return &resp, err
}
