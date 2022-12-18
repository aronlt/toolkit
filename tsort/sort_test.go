package tsort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortSlice(t *testing.T) {
	m := []int{1, 3, 2, 4, 1, 5, 6, 9, 8}
	SortSlice(m)
	n := []int{1, 1, 2, 3, 4, 5, 6, 8, 9}
	assert.Equal(t, m, n)
	m = []int{1, 3, 2, 4, 1, 5, 6, 9, 8}
	SortSlice(m, true)
	n = []int{9, 8, 6, 5, 4, 3, 2, 1, 1}
	assert.Equal(t, m, n)
}

func TestSortSliceWithComparator(t *testing.T) {
	m := []int{1, 3, 2, 4, 1, 5, 6, 9, 8}
	SortSliceWithComparator(m, func(i, j int) bool {
		return m[i] < m[j]
	})
}

func TestIsSorted(t *testing.T) {
	m := []int{1, 3, 2, 4, 1, 5, 6, 9, 8}
	ok := IsSorted(m)
	assert.False(t, ok)
	n := []int{1, 1, 2, 3, 4, 5, 6, 8, 9}
	ok = IsSorted(n)
	assert.True(t, ok)
}
