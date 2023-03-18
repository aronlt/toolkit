package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSlice(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	m2 := FpMapSlice(m, func(m []int, i int) int {
		return m[i] + 1
	})

	assert.Equal(t, m2, []int{2, 3, 4, 5, 6, 7})
}

func TestMapMap(t *testing.T) {
	m := map[int]int{1: 2, 2: 3, 3: 4}

	m2 := FpMapMap(m, func(_ map[int]int, k int, v int) (int, int) {
		return k, v + 1
	})

	assert.Equal(t, m2, map[int]int{1: 3, 2: 4, 3: 5})
}
