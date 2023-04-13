package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFpIterSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	counter := 1
	FpIterSlice(a, func(a []int, i int) {
		assert.Equal(t, a[i], counter)
		counter += 1
	})
}

func TestFpIterMap(t *testing.T) {
	a := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	b := make(map[int]int, 0)
	FpIterMap(a, func(a map[int]int, k int, v int) {
		b[k] = v
	})

	assert.Equal(t, b, a)
}

func TestFpIterList(t *testing.T) {
	a := DList[int]{}
	a.PushBack(1)
	a.PushBack(2)
	a.PushBack(3)
	a.PushBack(4)

	counter := 1
	FpIterList(a, func(a DList[int], node int) {
		assert.Equal(t, node, counter)
		counter += 1
	})
}

func TestFpIterSet(t *testing.T) {
	a := NewSet[int]()
	a.Insert(1)
	a.Insert(2)
	a.Insert(3)
	a.Insert(4)
	a.Insert(5)

	b := NewSet[int]()
	FpIterSet(a, func(a BuiltinSet[int], node int) {
		b.Insert(node)
	})

	sa := SetToSlice(a)
	sb := SetToSlice(b)
	assert.True(t, SliceLogicalEqual(sa, sb))
}
