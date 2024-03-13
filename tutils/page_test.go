package tutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryPage(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	v1 := MemoryPage(v, 10, 1)
	v2 := MemoryPage(v, 9, 1000)
	v3 := MemoryPage(v, -1, 1000)
	v4 := MemoryPage(v, 1, 5)
	v5 := MemoryPage(v, 1000, 10)
	assert.Equal(t, v1, []int{})
	assert.Equal(t, v2, []int{10})
	assert.Equal(t, v3, []int{})
	assert.Equal(t, v4, []int{2, 3, 4, 5, 6})
	assert.Equal(t, v5, []int{})
}
