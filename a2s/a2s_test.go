package a2s

import (
	"testing"
)

func TestSteamAPI(t *testing.T) {
	c, _ := NewClient("51.89.54.207", 27015)
	c.Info()
}
