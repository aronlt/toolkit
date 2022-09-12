package toolkit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseSlice(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	ReverseSlice(m)
	expected := []int{6, 5, 4, 3, 2, 1}
	unexpected := []int{6, 4, 5, 3, 2, 1}
	assert.Equal(t, m, expected)
	assert.NotEqual(t, m, unexpected)
}

func TestUniqueSlice(t *testing.T) {
	m := []int{1, 1, 3, 3, 5, 6}
	n := UniqueSlice(m)
	expected := []int{1, 3, 5, 6}
	unexpected := []int{1, 1, 3, 5, 6}
	assert.Equal(t, n, expected)
	assert.NotEqual(t, n, unexpected)
}

func TestCopySlice(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	n := CopySlice(m)
	assert.Equal(t, m, n)
}

func TestBinarySearch(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	idx := BinarySearch(m, 4)
	assert.Equal(t, idx, 3)

	m = []int{1, 2, 2, 2, 4, 4, 4, 5, 9}
	idx = BinarySearch(m, 4)
	assert.Equal(t, idx, 4)
	idx = BinarySearch(m, 3)
	assert.Equal(t, idx, -1)
}

func TestSliceMax(t *testing.T) {
	m := []int{10, 11, 12, 77, 21, 36, 34}
	n := SliceMax(m)
	assert.Equal(t, n, 77)
}

func TestSliceMin(t *testing.T) {
	m := []int{10, 11, 12, 77, 21, 36, 34}
	n := SliceMin(m)
	assert.Equal(t, n, 10)
}

func TestMax(t *testing.T) {
	m := Max(1, 2, 10, 18, 99, 10, 12)
	assert.Equal(t, m, 99)
}

func TestMin(t *testing.T) {
	m := Min(1, 2, 10, 18, 99, 10, 12)
	assert.Equal(t, m, 1)
}
