package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFpEverySlice(t *testing.T) {
	a := []int{2, 4, 6}
	v := FpEverySlice(a, func(a []int, i int) bool {
		return a[i]%2 == 0
	})
	assert.True(t, v)

	a = []int{1, 2, 4, 6}
	v = FpEverySlice(a, func(a []int, i int) bool {
		return a[i]%2 == 0
	})
	assert.False(t, v)
}

func TestFpEveryMap(t *testing.T) {
	a := map[int]int{2: 2, 4: 4, 6: 6}
	v := FpEveryMap(a, func(a map[int]int, k int, v int) bool {
		return v%2 == 0 && k%2 == 0
	})
	assert.True(t, v)

	a = map[int]int{2: 2, 4: 4, 6: 7}
	v = FpEveryMap(a, func(a map[int]int, k int, v int) bool {
		return v%2 == 0 && k%2 == 0
	})
	assert.False(t, v)
}

func TestFpEveryList(t *testing.T) {
	a := DList[int]{}
	a.PushBack(2)
	a.PushBack(4)
	a.PushBack(6)

	v := FpEveryList(a, func(a DList[int], node int) bool {
		return node%2 == 0
	})

	assert.True(t, v)

	a = DList[int]{}
	a.PushBack(2)
	a.PushBack(4)
	a.PushBack(6)
	a.PushBack(7)

	v = FpEveryList(a, func(a DList[int], node int) bool {
		return node%2 == 0
	})

	assert.False(t, v)
}

func TestFpEverySet(t *testing.T) {
	a := NewSet[int]()
	a.Insert(2)
	a.Insert(4)
	a.Insert(6)
	v := FpEverySet(a, func(a BuiltinSet[int], node int) bool {
		return node%2 == 0
	})
	assert.True(t, v)

	a = NewSet[int]()
	a.Insert(2)
	a.Insert(4)
	a.Insert(6)
	a.Insert(7)
	v = FpEverySet(a, func(a BuiltinSet[int], node int) bool {
		return node%2 == 0
	})
	assert.False(t, v)
}
