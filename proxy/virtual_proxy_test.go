package proxy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVirtualProxyBitmap(t *testing.T) {
	img := NewLazyBitmap("image")
	assert.Equal(t, img.Draw(), drwaing+"image")
}
