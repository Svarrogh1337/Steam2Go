package steamapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEconomy(t *testing.T) {
	key, _ := os.LookupEnv("apikey")
	c := NewClient(key)
	pricing, err := c.GetAssetPricesV1(context.Background(), 730)
	assert.NotNil(t, pricing.Result.Assets[0].Prices.Eur)
	assert.Nil(t, err)
	assets, _ := c.GetAssetClassInfoV1(context.Background(), 730, 4901045083)
	assert.NotNil(t, assets.Tags.Num0.Name)
	assert.NotNil(t, assets.Descriptions.Num2.AppData.Limited)
	assert.Nil(t, err)
}
