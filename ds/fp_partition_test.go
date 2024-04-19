package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFpPartitionSlice(t *testing.T) {
	a := []int{2, 4, 6, 7}
	pa, pb := SliceIterPartition(a, func(a []int, i int) bool {
		return a[i]%2 == 0
	})
	assert.Equal(t, pa, []int{2, 4, 6})
	assert.Equal(t, pb, []int{7})
}

func TestFpPartitionMap(t *testing.T) {
	a := map[int]int{2: 2, 4: 4, 6: 6, 7: 7}
	pa, pb := MapIterPartition(a, func(a map[int]int, k int, v int) bool {
		return k%2 == 0
	})
	assert.Equal(t, pa, map[int]int{2: 2, 4: 4, 6: 6})
	assert.Equal(t, pb, map[int]int{7: 7})
}

func TestFpPartitionDList(t *testing.T) {
	a := DList[int]{}
	a.PushBack(2)
	a.PushBack(4)
	a.PushBack(6)
	a.PushBack(7)

	pa, pb := DListPartition(a, func(a DList[int], node int) bool {
		return node%2 == 0
	})

	counter := 2
	pa.ForEach(func(val int) {
		assert.Equal(t, val, counter)
		counter += 2
	})

	pb.ForEach(func(val int) {
		assert.Equal(t, val, 7)
	})
}

func TestFpPartitionSList(t *testing.T) {
	a := SList[int]{}
	a.PushBack(2)
	a.PushBack(4)
	a.PushBack(6)
	a.PushBack(7)

	pa, pb := SListPartition(a, func(a SList[int], node int) bool {
		return node%2 == 0
	})

	counter := 2
	pa.ForEach(func(val int) {
		assert.Equal(t, val, counter)
		counter += 2
	})

	pb.ForEach(func(val int) {
		assert.Equal(t, val, 7)
	})
}

func TestFpPartitionSet(t *testing.T) {
	a := NewSet[int]()
	a.Insert(2)
	a.Insert(4)
	a.Insert(6)
	a.Insert(7)
	pa, pb := SetPartition(a, func(a BuiltinSet[int], node int) bool {
		return node%2 == 0
	})

	pa.ForEach(func(k int) {
		assert.True(t, k%2 == 0)
	})

	pb.ForEach(func(k int) {
		assert.True(t, k%2 == 1)
	})
}
