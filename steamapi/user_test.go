package steamapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	key, _ := os.LookupEnv("apikey")
	c := NewClient(key)
	c2 := NewClient("")
	userV1, err := c.GetFriendListV1(context.Background(), 76561198068163231)
	_, err2 := c2.GetFriendListV1(context.Background(), 76561198068163231)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond))
	defer cancel()
	_, err3 := c2.GetFriendListV1(ctx, 76561198068163231)
	_, err4 := c2.GetFriendListV1(nil, 76561198068163231)
	assert.Nil(t, err)
	assert.NotNil(t, userV1.Friendslist.Friends[0].Steamid)
	assert.NotNil(t, err2)
	assert.NotNil(t, err3)
	assert.NotNil(t, err4)
}
