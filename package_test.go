package Steam2Go

import (
	"github.com/Svarrogh1337/Steam2Go/steamapi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImports(t *testing.T) {
	c := steamapi.NewClient("")
	assert.NotNil(t, c)
}
