package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSlice(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	m2 := SliceIterMapCopy(m, func(i int) int {
		return m[i] + 1
	})

	assert.Equal(t, m2, []int{2, 3, 4, 5, 6, 7})
}

func TestMapMap(t *testing.T) {
	m := map[int]int{1: 2, 2: 3, 3: 4}

	m2 := MapIterMapKVCopy(m, func(k int, v int) (int, int) {
		return k, v + 1
	})

	assert.Equal(t, m2, map[int]int{1: 3, 2: 4, 3: 5})
}

func TestMapList(t *testing.T) {
	a := DList[int]{}
	a.PushBack(1)
	a.PushBack(2)
	a.PushBack(3)

	b := ListIterMapCopy(a, func(a DList[int], node int) int {
		return node + 1
	})

	count := 2
	b.ForEach(func(val int) {
		assert.Equal(t, val, count)
		count += 1
	})
}

func TestMapSet(t *testing.T) {
	a := NewSet[int]()
	a.Insert(1)
	a.Insert(2)
	a.Insert(3)

	b := SetIterMapCopy(a, func(node int) int {
		return node + 1
	})

	c := []int{2, 3, 4}
	SliceIterMapInPlace(c, func(i int) int {
		assert.True(t, b.Has(c[i]))
		return c[i]
	})

}
