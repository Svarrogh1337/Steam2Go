package steamapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestUser(t *testing.T) {
	key, _ := os.LookupEnv("apikey")
	c := NewClient(key)
	userV1, err := c.GetFriendListV1(context.Background(), 76561198068163231)
	assert.Nil(t, err)
	assert.NotNil(t, userV1)
}
