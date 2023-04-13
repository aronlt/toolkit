package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SList_Clean(t *testing.T) {
	sl := SList[int]{}
	sl.PushFront(1)
	sl.Clear()
	assert.True(t, sl.IsEmpty())
	assert.Equal(t, sl.Len(), 0)
}

func Test_SList_Front(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		sl := SList[int]{}
		assert.Panics(t, func() { sl.Front() })
	})

	t.Run("normal", func(t *testing.T) {
		sl := SList[int]{}
		sl.PushFront(1)
		assert.Equal(t, sl.Front(), 1)

		sl.PushBack(2)
		assert.Equal(t, sl.Front(), 1)

		sl.PushFront(3)
		assert.Equal(t, sl.Front(), 3)
	})
}

func Test_SList_Back(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		sl := SList[int]{}
		assert.Panics(t, func() { sl.Back() })
	})

	t.Run("normal", func(t *testing.T) {
		sl := SList[int]{}
		sl.PushBack(1)
		assert.Equal(t, sl.Back(), 1)

		sl.PushFront(2)
		assert.Equal(t, sl.Back(), 1)

		sl.PushBack(3)
		assert.Equal(t, sl.Back(), 3)
	})
}

func Test_SList_PushFront(t *testing.T) {
	sl := SList[int]{}
	for i := 1; i < 10; i++ {
		sl.PushFront(i)
		assert.Equal(t, sl.Front(), i)
		assert.Equal(t, sl.Len(), i)
	}
}

func Test_SList_PushBack(t *testing.T) {
	sl := SList[int]{}
	for i := 1; i < 10; i++ {
		sl.PushBack(i)
		assert.Equal(t, sl.Back(), i)
		assert.Equal(t, sl.Len(), i)
		assert.False(t, sl.IsEmpty())
	}
}

func Test_SList_PopFront(t *testing.T) {
	sl := SList[int]{}
	assert.Panics(t, func() { sl.PopFront() })

	sl.PushFront(1)
	sl.PushFront(2)
	assert.Equal(t, sl.PopFront(), 2)
	assert.Equal(t, sl.PopFront(), 1)
	assert.Panics(t, func() { sl.PopFront() })
}

func Test_SList_Reverse(t *testing.T) {
	sl := SListOf(1, 2, 3, 4)
	sl.Reverse()
	assert.Equal(t, sl.Values(), []int{4, 3, 2, 1})
}

func Test_SList_ForEach(t *testing.T) {
	sl := SListOf(1, 2, 3, 4)
	i := 0
	sl.ForEach(func(v int) {
		i++
		assert.Equal(t, v, i)
	})
	assert.Equal(t, i, sl.Len())
}

func Test_SList_ForEachIf(t *testing.T) {
	sl := SListOf(1, 2, 3, 4)
	i := 0
	sl.ForEachIf(func(v int) bool {
		i++
		assert.Equal(t, v, i)
		return i < 3
	})
	assert.Equal(t, i, 3)
}

func Test_SList_ForEachMutable(t *testing.T) {
	sl := SListOf(1, 2, 3, 4)
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

func Test_SList_ForEachMutableIf(t *testing.T) {
	sl := SListOf(1, 2, 3, 4)
	i := 0
	sl.ForEachMutableIf(func(v *int) bool {
		i++
		assert.Equal(t, *v, i)
		return i < 3
	})
	assert.Equal(t, i, 3)
}

func Test_SList_Iterate(t *testing.T) {
	sl := SList[int]{}
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
