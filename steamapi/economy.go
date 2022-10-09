package steamapi

import (
	"context"
	"fmt"
	"log"
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

const (
	iface = "ISteamEconomy"
)

func (c *Client) GetAssetPricesV1(ctx context.Context, appid int, options ...RequestParameter) (*GetAssetPricesResponseV1, error) {
	const (
		apiVersion = 1
		method     = "GetAssetPrices"
	)
	resp := GetAssetPricesResponseV1{}
	apiURL := fmt.Sprintf("%s/%s/%s/v%d?key=%s&appid=%d", c.baseURL, iface, method, apiVersion, c.apiKey, appid)
	if params := getOptionalParameters(options...).urlParams.Encode(); params != "" {
		apiURL += "&" + params
	}
	log.Printf(apiURL)
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}
