package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceIndex(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	i := SliceIndex(m, 4)
	assert.Equal(t, i, 3)
	i = SliceIndex(m, 8)
	assert.Equal(t, i, -1)
}

func TestSliceInclude(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	ok := SliceInclude(m, 4)
	assert.Equal(t, ok, true)
	ok = SliceInclude(m, 9)
	assert.Equal(t, ok, false)
}

func TestSliceExclude(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	ok := SliceExclude(m, 4)
	assert.Equal(t, ok, false)
	ok = SliceExclude(m, 8)
	assert.Equal(t, ok, true)
}

func TestSliceFilter(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	v := SliceFilter(m, func(i int) bool {
		return m[i] > 4
	})
	assert.Equal(t, v, []int{5, 6})
}

func TestSliceMap(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	SliceMap(m, func(i int) {
		m[i] += 1
	})
	assert.Equal(t, m, []int{2, 3, 4, 5, 6, 7})
}

func TestSliceAbsoluteEqual(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	n := []int{1, 3, 4, 2, 5, 6}
	h := []int{1, 2, 3, 4, 5, 6}
	ok := SliceAbsoluteEqual(m, n)
	assert.Equal(t, ok, false)
	ok = SliceAbsoluteEqual(m, h)
	assert.Equal(t, ok, true)
}

func TestSliceLogicalEqual(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	n := []int{1, 3, 4, 2, 5, 6}
	h := []int{1, 2, 3, 4, 5, 7}
	ok := SliceLogicalEqual(m, n)
	assert.Equal(t, ok, true)
	ok = SliceLogicalEqual(m, h)
	assert.Equal(t, ok, false)
}

func TestSliceReverseCopy(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	k := []int{6, 5, 4, 3, 2, 1}
	n := SliceReverseCopy(m)
	assert.Equal(t, n, k)
	n[0] = 9
	assert.Equal(t, m[0], 1)
}

func TestSliceRemove(t *testing.T) {
	m := []int{1, 2, 3, 3, 2, 1, 2, 3, 4, 2}
	k := []int{1, 3, 3, 1, 3, 4}
	SliceRemove(&m, 2)
	assert.Equal(t, k, m)
}

func TestSliceReplace(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	n := []int{1, 2, 9, 4, 5, 6}
	SliceReplace(m, 3, 9)
	assert.Equal(t, n, m)
}

func TestReverseSlice(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	SliceReverse(m)
	expected := []int{6, 5, 4, 3, 2, 1}
	unexpected := []int{6, 4, 5, 3, 2, 1}
	assert.Equal(t, m, expected)
	assert.NotEqual(t, m, unexpected)
}

func TestUniqueSlice(t *testing.T) {
	m := []int{1, 1, 3, 3, 5, 6}
	n := SliceUnique(m)
	expected := []int{1, 3, 5, 6}
	unexpected := []int{1, 1, 3, 5, 6}
	assert.Equal(t, n, expected)
	assert.NotEqual(t, n, unexpected)
}

func TestCopySlice(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	n := SliceCopy(m)
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

func TestMinN(t *testing.T) {
	data := []int{1, 2, 11, 23, 12, 113, 11}
	result := MinNWithOrder(data, 4)
	assert.Equal(t, result[0], 1)
	assert.Equal(t, result[1], 2)
	assert.Equal(t, result[2], 11)
	assert.Equal(t, result[3], 11)
}

func TestMaxN(t *testing.T) {
	data := []int{1, 2, 11, 23, 12, 113, 11}
	result := MaxNWithOrder(data, 4)
	assert.Equal(t, result[0], 113)
	assert.Equal(t, result[1], 23)
	assert.Equal(t, result[2], 12)
	assert.Equal(t, result[3], 11)
}
