package steamapi

import (
	"context"
	"fmt"
	"log"
)

type InventoryResponse struct {
	Assets []struct {
		Appid      int    `json:"appid"`
		Contextid  string `json:"contextid"`
		Assetid    string `json:"assetid"`
		Classid    string `json:"classid"`
		Instanceid string `json:"instanceid"`
		Amount     string `json:"amount"`
	} `json:"assets"`
	Descriptions []struct {
		Appid           int    `json:"appid"`
		Classid         string `json:"classid"`
		Instanceid      string `json:"instanceid"`
		Currency        int    `json:"currency"`
		BackgroundColor string `json:"background_color"`
		IconURL         string `json:"icon_url"`
		IconURLLarge    string `json:"icon_url_large,omitempty"`
		Descriptions    []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
			Color string `json:"color,omitempty"`
		} `json:"descriptions"`
		Tradable int `json:"tradable"`
		Actions  []struct {
			Link string `json:"link"`
			Name string `json:"name"`
		} `json:"actions,omitempty"`
		OwnerDescriptions []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
			Color string `json:"color,omitempty"`
		} `json:"owner_descriptions,omitempty"`
		Name           string `json:"name"`
		NameColor      string `json:"name_color"`
		Type           string `json:"type"`
		MarketName     string `json:"market_name"`
		MarketHashName string `json:"market_hash_name"`
		MarketActions  []struct {
			Link string `json:"link"`
			Name string `json:"name"`
		} `json:"market_actions,omitempty"`
		Commodity                 int `json:"commodity"`
		MarketTradableRestriction int `json:"market_tradable_restriction"`
		Marketable                int `json:"marketable"`
		Tags                      []struct {
			Category              string `json:"category"`
			InternalName          string `json:"internal_name"`
			LocalizedCategoryName string `json:"localized_category_name"`
			LocalizedTagName      string `json:"localized_tag_name"`
			Color                 string `json:"color,omitempty"`
		} `json:"tags"`
		Fraudwarnings []string `json:"fraudwarnings,omitempty"`
	} `json:"descriptions"`
	TotalInventoryCount int `json:"total_inventory_count"`
	Success             int `json:"success"`
	Rwgrsn              int `json:"rwgrsn"`
}

func (c *Client) GetInventory(ctx context.Context, appid int, steamid int) (*InventoryResponse, error) {
	c.baseURL = "https://steamcommunity.com"
	const (
		iface      = "inventory"
		apiVersion = 2
	)
	resp := InventoryResponse{}
	apiURL := fmt.Sprintf("%s/%s/%d/%d/%d", c.baseURL, iface, steamid, appid, apiVersion)
	log.Println(apiURL)
	err := c.get(ctx, apiURL, &resp)
	return &resp, err
}
