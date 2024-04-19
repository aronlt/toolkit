package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFpEverySlice(t *testing.T) {
	a := []int{2, 4, 6}
	v := SliceIterAllOkV2(a, func(i int) bool {
		return a[i]%2 == 0
	})
	assert.True(t, v)

	a = []int{1, 2, 4, 6}
	v = SliceIterAllOkV2(a, func(i int) bool {
		return a[i]%2 == 0
	})
	assert.False(t, v)
}

func TestFpEveryMap(t *testing.T) {
	a := map[int]int{2: 2, 4: 4, 6: 6}
	v := MapIterAllOk(a, func(k int, v int) bool {
		return v%2 == 0 && k%2 == 0
	})
	assert.True(t, v)

	a = map[int]int{2: 2, 4: 4, 6: 7}
	v = MapIterAllOk(a, func(k int, v int) bool {
		return v%2 == 0 && k%2 == 0
	})
	assert.False(t, v)
}

func TestFpEverySList(t *testing.T) {
	a := SList[int]{}
	a.PushBack(2)
	a.PushBack(4)
	a.PushBack(6)

	v := SListIterAllOk(a, func(a SList[int], node int) bool {
		return node%2 == 0
	})

	assert.True(t, v)

	a = SList[int]{}
	a.PushBack(2)
	a.PushBack(4)
	a.PushBack(6)
	a.PushBack(7)

	v = SListIterAllOk(a, func(a SList[int], node int) bool {
		return node%2 == 0
	})

	assert.False(t, v)
}

func TestFpEveryDList(t *testing.T) {
	a := DList[int]{}
	a.PushBack(2)
	a.PushBack(4)
	a.PushBack(6)

	v := DListIterAllOk(a, func(a DList[int], node int) bool {
		return node%2 == 0
	})

	assert.True(t, v)

	a = DList[int]{}
	a.PushBack(2)
	a.PushBack(4)
	a.PushBack(6)
	a.PushBack(7)

	v = DListIterAllOk(a, func(a DList[int], node int) bool {
		return node%2 == 0
	})

	assert.False(t, v)
}

func TestFpEverySet(t *testing.T) {
	a := NewSet[int]()
	a.Insert(2)
	a.Insert(4)
	a.Insert(6)
	v := SetIterAllOk(a, func(node int) bool {
		return node%2 == 0
	})
	assert.True(t, v)

	a = NewSet[int]()
	a.Insert(2)
	a.Insert(4)
	a.Insert(6)
	a.Insert(7)
	v = SetIterAllOk(a, func(node int) bool {
		return node%2 == 0
	})
	assert.False(t, v)
}
