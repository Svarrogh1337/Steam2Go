package steamapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApps(t *testing.T) {
	c := NewClient("")
	appsv1, err := c.GetAppListV1(context.Background())
	assert.NotNil(t, appsv1.Applist.Apps.App[0].Appid)
	assert.Nil(t, err)
	appsv2, err := c.GetAppListV2(context.Background())
	assert.NotNil(t, appsv2.Applist.Apps[0].Appid)
	assert.Nil(t, err)
}
