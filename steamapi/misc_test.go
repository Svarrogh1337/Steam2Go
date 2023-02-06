package steamapi

import (
	"context"
	"log"
	"os"
	"testing"
)

func TestMiscInventory(t *testing.T) {
	key, _ := os.LookupEnv("apikey")
	log.Println(key)
	c := NewClient(key)
	x, _ := c.GetInventory(context.Background(), 730, 76561198068163231)
	log.Println(x.Assets[0])
}
