package steamapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	steamids := []string{"76561198068163231", "76561197982426067"}
	steamids2 := []string{""}
	key, _ := os.LookupEnv("apikey")
	c := NewClient(key)
	c2 := NewClient("")
	userV1, err := c.GetFriendListV1(context.Background(), 76561198068163231, Relationship("friend"))
	assert.NotNil(t, userV1.Friendslist.Friends[0].Steamid)
	assert.Nil(t, err)
	_, err1 := c2.GetFriendListV1(context.Background(), 76561198068163231)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond))
	defer cancel()
	assert.NotNil(t, err1)
	_, err2 := c2.GetFriendListV1(ctx, 76561198068163231)
	assert.NotNil(t, err2)
	_, err3 := c2.GetFriendListV1(nil, 76561198068163231)
	assert.NotNil(t, err3)
	pb, err4 := c.GetPlayerBansV1(context.Background(), steamids)
	assert.NotNil(t, pb.Players[1].EconomyBan)
	assert.Nil(t, err4)
	pb2, err5 := c.GetPlayerBansV1(context.Background(), steamids2)
	assert.Empty(t, pb2.Players)
	assert.Nil(t, err5)
}
