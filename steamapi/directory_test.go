package steamapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirectory(t *testing.T) {
	c := NewClient("")
	cmlist, err := c.GetCMListForConnectV1(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, cmlist)
	assert.True(t, true, cmlist.Response.Success)
}
