package tcache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemCache(t *testing.T) {
	m := NewMemCache[int, string]()
	_, ok := m.Get(1)
	assert.False(t, ok)

	m.Load(1, "ok")
	v, ok := m.Get(1)
	assert.True(t, ok)
	assert.Equal(t, v, "ok")
}
