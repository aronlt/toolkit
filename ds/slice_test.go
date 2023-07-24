package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceIndex(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	i := SliceIncludeIndex(m, 4)
	assert.Equal(t, i, 3)
	i = SliceIncludeIndex(m, 8)
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
	v := SliceGetFilter(m, func(i int) bool {
		return m[i] > 4
	})
	assert.Equal(t, v, []int{5, 6})
}

func TestSliceAbsoluteEqual(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	n := []int{1, 3, 4, 2, 5, 6}
	h := []int{1, 2, 3, 4, 5, 6}
	ok := SliceCmpAbsEqual(m, n)
	assert.Equal(t, ok, false)
	ok = SliceCmpAbsEqual(m, h)
	assert.Equal(t, ok, true)
}

func TestSliceLogicalEqual(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	n := []int{1, 3, 4, 2, 5, 6}
	h := []int{1, 2, 3, 4, 5, 7}
	ok := SliceCmpLogicEqual(m, n)
	assert.Equal(t, ok, true)
	ok = SliceCmpLogicEqual(m, h)
	assert.Equal(t, ok, false)
}

func TestSliceReverseCopy(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	k := []int{6, 5, 4, 3, 2, 1}
	n := SliceOpReverseCopy(m)
	assert.Equal(t, n, k)
	n[0] = 9
	assert.Equal(t, m[0], 1)
}

func TestSliceRemove(t *testing.T) {
	m := []int{1, 2, 3, 3, 2, 1, 2, 3, 4, 2}
	k := []int{1, 3, 3, 1, 3, 4}
	SliceOpRemove(&m, 2)
	assert.Equal(t, k, m)
}

func TestSliceRemoveMany(t *testing.T) {
	m := []int{1, 2, 3, 3, 5, 6, 2, 1, 2, 3, 4, 2}
	k := []int{1, 3, 2, 3, 1, 3, 4}
	SliceOpRemoveMany(&m, k)
	assert.Equal(t, m, []int{5, 6})
}

func TestSliceReplace(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	n := []int{1, 2, 9, 4, 5, 6}
	SliceOpReplace(m, 3, 9)
	assert.Equal(t, n, m)
}

func TestReverseSlice(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	SliceOpReverse(m)
	expected := []int{6, 5, 4, 3, 2, 1}
	unexpected := []int{6, 4, 5, 3, 2, 1}
	assert.Equal(t, m, expected)
	assert.NotEqual(t, m, unexpected)
}

func TestUniqueSlice(t *testing.T) {
	m := []int{1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 6}
	n := SliceOpUnique(m)
	expected := []int{1, 3, 5, 6}
	unexpected := []int{1, 1, 3, 5, 6}
	assert.Equal(t, n, expected)
	assert.NotEqual(t, n, unexpected)
}

func TestCopySlice(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	n := SliceGetCopy(m)
	assert.Equal(t, m, n)
}

func TestBinarySearch(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	idx := SliceIncludeBinarySearch(m, 4)
	assert.Equal(t, idx, 3)

	m = []int{1, 2, 2, 2, 4, 4, 4, 5, 9}
	idx = SliceIncludeBinarySearch(m, 4)
	assert.Equal(t, idx, 4)
	idx = SliceIncludeBinarySearch(m, 3)
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
	m := SliceMaxUnpack(1, 2, 10, 18, 99, 10, 12)
	assert.Equal(t, m, 99)
}

func TestMin(t *testing.T) {
	m := SliceMinUnpack(1, 2, 10, 18, 99, 10, 12)
	assert.Equal(t, m, 1)
}

func TestMinN(t *testing.T) {
	data := []int{1, 2, 11, 23, 12, 113, 11}
	result := SliceMinNWithOrder(data, 4)
	assert.Equal(t, result[0], 1)
	assert.Equal(t, result[1], 2)
	assert.Equal(t, result[2], 11)
	assert.Equal(t, result[3], 11)
}

func TestMaxN(t *testing.T) {
	data := []int{1, 2, 11, 23, 12, 113, 11}
	result := SliceMaxNWithOrder(data, 4)
	assert.Equal(t, result[0], 113)
	assert.Equal(t, result[1], 23)
	assert.Equal(t, result[2], 12)
	assert.Equal(t, result[3], 11)
}

func TestSliceConvertToInt64(t *testing.T) {
	data := []uint{1, 2, 3, 4, 5}
	ints, err := SliceConvertToInt64(data)
	assert.Nil(t, err)
	assert.Equal(t, ints[0], int64(1))
	assert.Equal(t, len(data), len(ints))

	strs := []string{"1", "2", "3"}
	ints, err = SliceConvertToInt64(strs)
	assert.Nil(t, err)
	assert.Equal(t, ints[0], int64(1))
	assert.Equal(t, len(strs), len(ints))
}

func TestSliceConvertToInt(t *testing.T) {
	data := []uint{1, 2, 3, 4, 5}
	ints, err := SliceConvertToInt(data)
	assert.Nil(t, err)
	assert.Equal(t, ints[0], 1)
	assert.Equal(t, len(data), len(ints))

	strs := []string{"1", "2", "3"}
	ints, err = SliceConvertToInt(strs)
	assert.Nil(t, err)
	assert.Equal(t, ints[0], 1)
	assert.Equal(t, len(strs), len(ints))

	strs2 := []string{"1", "2", "3FFXX"}
	ints, err = SliceConvertToInt(strs2)
	assert.NotNil(t, err)
}

func TestSliceConvertToString(t *testing.T) {
	data := []uint{1, 2, 3, 4, 5}
	ints, err := SliceConvertToString(data)
	assert.Nil(t, err)
	assert.Equal(t, ints[0], "1")
	assert.Equal(t, len(data), len(ints))

	strs := []string{"1", "2", "3FFFFx"}
	ints, err = SliceConvertToString(strs)
	assert.Nil(t, err)
}

func TestSliceInsert(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	SliceOpInsert(&data, 1, 8, 9)
	assert.Equal(t, []int{1, 8, 9, 2, 3, 4, 5}, data)

	SliceOpInsert(&data, 17, 10)
	assert.Equal(t, []int{1, 8, 9, 2, 3, 4, 5, 10}, data)

	SliceOpInsert(&data, -2, 11)
	assert.Equal(t, []int{1, 8, 9, 2, 3, 4, 11, 5, 10}, data)
}

func TestSliceTail(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	v := SliceGetTail(data)
	assert.Equal(t, v, 6)

	var ok bool
	v, ok = SliceOpPopBack(&data)

	assert.True(t, ok, true)
	assert.Equal(t, v, 6)
	assert.Equal(t, data, []int{1, 2, 3, 4, 5})
}

func TestSliceRemoveIndex(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	SliceOpRemoveIndex(&data, 3)
	assert.Equal(t, data, []int{1, 2, 3, 5, 6})

}

func TestSliceRemoveRange(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	SliceOpRemoveRange(&data, 3, 5)
	assert.Equal(t, data, []int{1, 2, 3, 6})
}

func TestSliceIncludeWithFn(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6}
	ok := SliceIncludeWithFn(data, func(ints []int, i int) bool {
		return ints[i] == 3
	})
	assert.True(t, ok)
	ok = SliceIncludeWithFn(data, func(ints []int, i int) bool {
		return ints[i] == 9
	})
	assert.False(t, ok)
}

func TestInclude(t *testing.T) {
	ok := SliceIncludeUnpack(1, 2, 3, 1)
	assert.True(t, ok)
	ok = SliceIncludeUnpack(4, 2, 3, 1)
	assert.False(t, ok)
}

func TestDiffTwoSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{6, 7, 8, 9, 10}
	sa, sb := SliceCmpTwoDiff(a, b)
	assert.True(t, SliceCmpLogicEqual(sa, []int{1, 2, 3, 4, 5}))
	assert.True(t, SliceCmpLogicEqual(sb, []int{8, 9, 10}))
}

func TestSliceGroupByHandler(t *testing.T) {
	a := []int{1, 2, 3, 4, 4, 3, 2, 1, 33}
	b := SliceGroupByHandler(a, func(i int) int {
		return a[i]
	})

	assert.Equal(t, b[1], []int{1, 1})
	assert.Equal(t, b[33], []int{33})
}

func TestSliceGroupByHandlerUnique(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := SliceGroupByHandlerUnique(a, func(i int) int {
		return a[i]
	})

	assert.Equal(t, b[1], 1)
	assert.Equal(t, b[3], 3)
}

func TestSliceGroupIntoSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 4, 3, 2, 1, 33}
	b := SliceGroupToSlices(a)

	assert.Equal(t, b[0], []int{1, 1})
	assert.Equal(t, b[1], []int{2, 2})
	assert.Equal(t, len(b), 5)
}

func TestSliceGroupByValue(t *testing.T) {
	a := []int{1, 2, 3, 4, 4, 3, 2, 1, 33}
	b := SliceGroupToMap(a)

	assert.Equal(t, b[1], []int{1, 1})
	assert.Equal(t, b[2], []int{2, 2})
}

func TestSliceCmpLogicSub(t *testing.T) {
	a := []int{1, 2, 3, 3, 4, 5}
	b := []int{1, 3, 2, 3}
	assert.True(t, SliceCmpLogicSub(a, b))
}

func TestSliceCmpAbsSub(t *testing.T) {
	a := []int{1, 2, 3, 3, 4, 5}
	b := []int{2, 3, 3}
	assert.Equal(t, SliceCmpAbsSub(a, b), 1)
	a = []int{1, 2, 3, 3, 4, 5}
	b = []int{4, 5, 6}
	assert.Equal(t, SliceCmpAbsSub(a, b), -1)
}

func TestSliceGroupToSet(t *testing.T) {
	a := []int{1, 1, 2, 3, 3, 4, 5}
	set := SliceGroupToSet(a)
	set2 := SetFromSlice(a)
	assert.Equal(t, set, set2)
}

func TestSliceOpMerge(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 10}
	c := SliceOpMerge(a, b)
	assert.Equal(t, c, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}
