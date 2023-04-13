package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_List_Clean(t *testing.T) {
	sl := DList[int]{}
	sl.PushFront(1)
	sl.Clear()
	assert.True(t, sl.IsEmpty())
	assert.Equal(t, sl.Len(), 0)
}

func Test_List_PushFront(t *testing.T) {
	sl := DList[int]{}
	for i := 1; i < 10; i++ {
		sl.PushFront(i)
		assert.Equal(t, sl.Front(), i)
		assert.Equal(t, sl.Len(), i)
	}
}

func Test_List_PushBack(t *testing.T) {
	sl := DList[int]{}
	for i := 1; i < 10; i++ {
		sl.PushBack(i)
		assert.Equal(t, sl.Back(), i)
		assert.Equal(t, sl.Len(), i)
		assert.False(t, sl.IsEmpty())
	}
}

func Test_List_PopFront(t *testing.T) {
	sl := DList[int]{}
	sl.PushFront(1)
	sl.PushFront(2)
	assert.Equal(t, sl.PopFront(), 2)
	assert.Equal(t, sl.PopFront(), 1)
	assert.Panics(t, func() { sl.PopFront() })
}
func Test_List_ForEach(t *testing.T) {
	sl := ListOf(1, 2, 3, 4)
	i := 0
	sl.ForEach(func(v int) {
		i++
		assert.Equal(t, v, i)
	})
	assert.Equal(t, i, sl.Len())
}

func Test_List_ForEachIf(t *testing.T) {
	sl := ListOf(1, 2, 3, 4)
	i := 0
	sl.ForEachIf(func(v int) bool {
		i++
		assert.Equal(t, v, i)
		return i < 3
	})
	assert.Equal(t, i, 3)
}

func Test_List_ForEachMutable(t *testing.T) {
	sl := ListOf(1, 2, 3, 4)
	i := 0
	sl.ForEachMutable(func(v *int) {
		i++
		assert.Equal(t, *v, i)
		*v = -*v
	})
	assert.Equal(t, i, sl.Len())
	sl.ForEachMutable(func(v *int) {
		assert.Less(t, *v, 0)
	})
}

func Test_List_ForEachMutableIf(t *testing.T) {
	sl := ListOf(1, 2, 3, 4)
	i := 0
	sl.ForEachMutableIf(func(v *int) bool {
		i++
		assert.Equal(t, *v, i)
		return i < 3
	})
	assert.Equal(t, i, 3)
}

func Test_List_Iterate(t *testing.T) {
	sl := DList[int]{}
	sl.PushBack(1)
	sl.PushBack(2)
	sl.PushBack(3)
	i := 0
	for it := sl.Iterate(); it.IsNotEnd(); it.MoveToNext() {
		i++
		assert.Equal(t, it.Value(), i)
		assert.Equal(t, *it.Pointer(), i)
	}
	assert.Equal(t, i, 3)
}
