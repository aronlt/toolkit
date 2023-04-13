package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFpFilterSlice(t *testing.T) {
	a := []int{2, 4, 6, 7}
	v := FpFilterSlice(a, func(a []int, i int) bool {
		return a[i] == 7
	})
	assert.Equal(t, v, []int{7})
}

func TestFpFilterMap(t *testing.T) {
	a := map[int]int{2: 2, 4: 4, 6: 6}
	v := FpFilterMap(a, func(a map[int]int, k int, v int) bool {
		return k == 2
	})
	assert.Equal(t, v, map[int]int{2: 2})
}

func TestFpFilterList(t *testing.T) {
	a := DList[int]{}
	a.PushBack(2)
	a.PushBack(4)
	a.PushBack(6)

	v := FpFilterList(a, func(a DList[int], node int) bool {
		return node == 2
	})
	v.ForEach(func(val int) {
		assert.Equal(t, val, 2)
	})
}

func TestFpFilterSet(t *testing.T) {
	a := NewSet[int]()
	a.Insert(2)
	a.Insert(4)
	a.Insert(6)
	v := FpFilterSet(a, func(a BuiltinSet[int], node int) bool {
		return node == 2
	})
	v.ForEach(func(k int) {
		assert.Equal(t, k, 2)
	})
}
