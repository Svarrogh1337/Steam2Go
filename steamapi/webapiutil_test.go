package steamapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWebAPIUtils(t *testing.T) {
	c := NewClient("")
	sinfo, err := c.GetServerInfoV1(context.Background())
	assert.NotNil(t, sinfo.Servertime)
	assert.Nil(t, err)
}
