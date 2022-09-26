package steamapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSteamAPI(t *testing.T) {
	c := NewClient("")
	err := c.get(context.Background(), "http://httpstat.us/100", GetFriendListResponseV1{})
	assert.NotNil(t, err)
}
