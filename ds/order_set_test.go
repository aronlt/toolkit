package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OrderBuiltinSet_IsEmpty(t *testing.T) {
	os := NewOrderSet[string]()
	assert.Equal(t, os.IsEmpty(), true)
	os.Insert("hello")
	assert.Equal(t, os.IsEmpty(), false)
}

func Test_OrderBuiltinSet_Clear(t *testing.T) {
	os := OrderSetFromUnpack("hello", "world")
	os.Clear()
	assert.True(t, os.IsEmpty())
}

func Test_OrderBuiltinSet_Has(t *testing.T) {
	os := OrderSetFromUnpack("hello", "world")

	assert.True(t, os.Has("hello"))
	assert.True(t, os.Has("world"))
	assert.False(t, os.Has("!"))
}

func Test_OrderBuiltinSet_Insert(t *testing.T) {
	s := NewOrderSet[string]()
	assert.True(t, s.Insert("hello"))
	assert.False(t, s.Insert("hello"))
	assert.Equal(t, s.Has("world"), false)
	assert.True(t, s.Insert("world"))
	assert.Equal(t, s.Has("hello"), true)
	assert.Equal(t, s.Len(), 2)
}

func Test_OrderBuiltinSet_InsertN(t *testing.T) {
	s := NewOrderSet[string]()
	assert.Equal(t, s.InsertN("hello", "world"), 2)
	assert.Equal(t, s.Len(), 2)
}

func Test_OrderBuiltinSet_Remove(t *testing.T) {
	os := OrderSetFromUnpack("hello", "world")
	assert.True(t, os.Remove("hello"))
	assert.Equal(t, os.Len(), 1)
	assert.False(t, os.Remove("hello"))
	assert.Equal(t, os.Len(), 1)
	assert.True(t, os.Remove("world"))
	assert.Equal(t, os.Len(), 0)
}

func Test_OrderBuiltinSet_Delete(t *testing.T) {
	os := OrderSetFromUnpack("hello", "world")
	os.Delete("hello")
	assert.Equal(t, os.Len(), 1)
	os.Delete("hello")
	assert.Equal(t, os.Len(), 1)
	os.Delete("world")
	assert.Equal(t, os.Len(), 0)
}

func Test_OrderBuiltinSet_RemoveN(t *testing.T) {
	os := OrderSetFromUnpack("hello", "world")
	assert.Equal(t, os.RemoveN("hello", "world"), 2)
	assert.False(t, os.Remove("world"))
	assert.True(t, os.IsEmpty())
}

func Test_OrderBuiltinSet_Keys(t *testing.T) {
	os := OrderSetFromUnpack("world", "hello")
	ks := os.Keys()
	assert.Equal(t, 2, ks.length)
	assert.Equal(t, ks.Values(), []string{"hello", "world"})
}

func Test_OrderBuiltinSet_For(t *testing.T) {
	os := OrderSetFromUnpack("hello", "world")
	os.ForEach(func(k string) {
		assert.True(t, k == "hello" || k == "world")
	})

}

func Test_OrderBuiltinSet_ForEach(t *testing.T) {
	os := OrderSetFromUnpack("world", "hello", "go")
	v := make([]string, 0)
	os.ForEach(func(k string) {
		v = append(v, k)
	})
	assert.Equal(t, v, []string{"go", "hello", "world"})
}

func Test_OrderBuiltinSet_ForEachIf(t *testing.T) {
	os := OrderSetFromUnpack("hello", "go", "world")
	v := make([]string, 0)
	os.ForEachIf(func(k string) bool {
		v = append(v, k)
		return false
	})
	assert.Equal(t, v, []string{"go"})
}

func Test_OrderBuiltinSet_Update(t *testing.T) {
	os := OrderSetFromUnpack(1, 2, 3)
	other := OrderSetFromUnpack(3, 4)
	os.Update(other)
	assert.Equal(t, os.Len(), 4)
	assert.True(t, os.Has(4))
}

func Test_OrderBuiltinSet_Union(t *testing.T) {
	s := OrderSetFromUnpack(1, 2, 3)
	other := OrderSetFromUnpack(3, 4)
	s2 := s.Union(other)
	assert.Equal(t, s2.Len(), 4)
	assert.True(t, s2.Has(4))
}

func Test_OrderBuiltinSet_Intersection(t *testing.T) {
	s := OrderSetFromUnpack(1, 2, 3).Intersection(OrderSetFromUnpack(3, 4))
	assert.Equal(t, s.Len(), 1)
	assert.True(t, s.Has(3))
	s = OrderSetFromUnpack(3, 4).Intersection(OrderSetFromUnpack(1, 2, 3))
	assert.Equal(t, s.Len(), 1)
	assert.True(t, s.Has(3))
}

func Test_OrderBuiltinSet_Difference(t *testing.T) {
	s := OrderSetFromUnpack(1, 2, 3).Difference(OrderSetFromUnpack(3, 4))
	assert.Equal(t, s.Len(), 2)
	assert.True(t, s.Has(1))
	assert.True(t, s.Has(2))
	s = OrderSetFromUnpack(1, 2).Difference(OrderSetFromUnpack(3, 4))
	assert.Equal(t, s.Len(), 2)
	assert.True(t, s.Has(1))
	assert.True(t, s.Has(2))
}

func Test_OrderBuiltinSet_IsDisjointOf(t *testing.T) {
	s1 := OrderSetFromUnpack(1, 2, 3)
	s2 := OrderSetFromUnpack(3, 4)
	assert.False(t, s1.IsDisjointOf(s2))
	assert.True(t, s1.IsDisjointOf(OrderSetFromUnpack(4, 5)))
}

func Test_OrderBuiltinSet_IsSubsetOf(t *testing.T) {
	assert.True(t, OrderSetFromUnpack[int]().IsSubsetOf(OrderSetFromUnpack[int]()))
	assert.True(t, OrderSetFromUnpack[int]().IsSubsetOf(OrderSetFromUnpack(1)))
	assert.True(t, OrderSetFromUnpack(1, 2, 3).IsSubsetOf(OrderSetFromUnpack(1, 2, 3)))
	assert.True(t, OrderSetFromUnpack(1, 2).IsSubsetOf(OrderSetFromUnpack(1, 2, 3)))
	assert.False(t, OrderSetFromUnpack(1, 2, 3).IsSubsetOf(OrderSetFromUnpack(1, 2)))
	assert.False(t, OrderSetFromUnpack(1, 2).IsSubsetOf(OrderSetFromUnpack(2, 3)))
}

func Test_OrderBuiltinSet_IsSupersetOf(t *testing.T) {
	assert.True(t, OrderSetFromUnpack[int]().IsSupersetOf(OrderSetFromUnpack[int]()))
	assert.True(t, OrderSetFromUnpack(1).IsSupersetOf(OrderSetFromUnpack[int]()))
	assert.True(t, OrderSetFromUnpack(1, 2, 3).IsSupersetOf(OrderSetFromUnpack(1, 2, 3)))
	assert.True(t, OrderSetFromUnpack(1, 2, 3).IsSupersetOf(OrderSetFromUnpack(1, 2)))
	assert.False(t, OrderSetFromUnpack(1, 2).IsSupersetOf(OrderSetFromUnpack(1, 2, 3)))
	assert.False(t, OrderSetFromUnpack(1, 2).IsSupersetOf(OrderSetFromUnpack(2, 3)))
}

func Test_OrderBuiltinSet_Equal(t *testing.T) {
	v1 := OrderSetFromUnpack[int](1, 2, 3, 4, 5)
	v2 := OrderSetFromUnpack[int](1, 2, 3, 4, 6)
	assert.False(t, v1.Equal(v2))

	v3 := OrderSetFromUnpack[int](1, 2, 3, 4, 5)
	assert.True(t, v1.Equal(v3))

	v4 := OrderSetFromUnpack[int](1, 2, 3, 4)
	assert.False(t, v1.Equal(v4))
	assert.False(t, v4.Equal(v1))

}
