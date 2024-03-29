package ds

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func IntComparator(a int, b int) int { return a - b }

func TestInsert(t *testing.T) {
	list := NewSkipList[int, int](IntComparator, WithMaxLevel(5))

	m := make(map[int]int)
	for i := 0; i < 100; i++ {
		key := rand.Int() % 100
		list.Insert(key, i)
		m[key] = i
	}
	for key, v := range m {
		ret, _ := list.Get(key)
		assert.Equal(t, v, ret)
	}
	assert.Equal(t, len(m), list.Len())
}

func TestRemove(t *testing.T) {
	list := NewSkipList[int, int](IntComparator, WithMaxLevel(5))

	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		key := rand.Int() % 1000
		list.Insert(key, i)
		m[key] = i
	}
	assert.Equal(t, len(m), list.Len())

	for i := 0; i < 300; i++ {
		key := rand.Int() % 1000
		list.Remove(key)
		delete(m, key)
		key2 := rand.Int() % 10440
		list.Insert(key2, key)
		m[key2] = key
	}

	for key, v := range m {
		ret, _ := list.Get(key)
		assert.Equal(t, v, ret)
	}
	assert.Equal(t, len(m), list.Len())
}

func TestSkiplist_Traversal(t *testing.T) {
	list := NewSkipList[int, int](IntComparator, WithMaxLevel(5))
	for i := 0; i < 10; i++ {
		list.Insert(i, i*10)
	}
	keys := list.Keys()
	for i := 0; i < 10; i++ {
		assert.Equal(t, i, keys[i])
	}
	i := 0
	list.Traversal(func(key, value int) bool {
		assert.Equal(t, i, key)
		assert.Equal(t, i*10, value)
		i++
		return true
	})
}
