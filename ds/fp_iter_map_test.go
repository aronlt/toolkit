package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceIterMap(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	SliceIterMapInPlace(a, func(i int) int {
		return a[i] + 1
	})
	assert.Equal(t, a, []int{2, 3, 4, 5, 6})
}

func TestMapIterKV(t *testing.T) {
	a := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	MapIterMapKVInPlace(a, func(k int, v int) int {
		return v + 1
	})
	assert.Equal(t, a, map[int]int{1: 2, 2: 3, 3: 4, 4: 5})
}

func TestFpEachList(t *testing.T) {
	a := DList[int]{}
	a.PushBack(1)
	a.PushBack(2)
	a.PushBack(3)
	a.PushBack(4)

	ListIterMapInPlace(a, func(a DList[int], node int) int {
		return node + 1
	})

	count := 2
	a.ForEach(func(val int) {
		assert.Equal(t, count, val)
		count += 1
	})
}

func TestFpEachSet(t *testing.T) {
	a := NewSet[int]()
	a.Insert(1)
	a.Insert(2)
	a.Insert(3)
	a.Insert(4)
	a.Insert(5)
	SetIterMapInPlace(a, func(node int) int {
		return node + 1
	})

	b := []int{2, 3, 4, 5, 6}
	SliceIterMapInPlace(b, func(i int) int {
		assert.True(t, a.Has(b[i]))
		return b[i]
	})

}
