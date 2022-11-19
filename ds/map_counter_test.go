package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCounterMap(t *testing.T) {
	data := []int{1, 2, 3, 1, 2, 3, 4, 5}
	counter := NewCounterMap(data)
	assert.Equal(t, 2, counter[2])
	assert.Equal(t, 1, counter[4])
}
