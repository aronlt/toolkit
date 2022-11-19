package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedMap(t *testing.T) {
	m := make(map[int]int, 0)
	for i := 0; i < 10; i++ {
		m[i] = i
	}
	sortedMap := NewSortedMap(m)
	for i := 0; i < 10; i++ {
		assert.Equal(t, sortedMap.Tuples[i].Key, i)
		assert.Equal(t, sortedMap.Tuples[i].Value, i)
	}
	for i := 0; i < 10; i++ {
		sortedMap.RawMap[i] = i + 1
	}
	sortedMap.Rebuild()
	for i := 0; i < 10; i++ {
		assert.Equal(t, sortedMap.Tuples[i].Key, i)
		assert.Equal(t, sortedMap.Tuples[i].Value, i+1)
	}
}
