package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MakeBuiltinSet(t *testing.T) {
	s := make(BuiltinSet[string])
	assert.Equal(t, s.Len(), 0)
	assert.Equal(t, s.IsEmpty(), true)
}

func Test_MakeBuiltinSet2(t *testing.T) {
	s := BuiltinSet[string]{}
	assert.Equal(t, s.Len(), 0)
	assert.Equal(t, s.IsEmpty(), true)
}

func Test_SetOf(t *testing.T) {
	s := SetOf("hello", "world")
	assert.Equal(t, s.Len(), 2)
}

func Test_BuiltinSet_IsEmpty(t *testing.T) {
	s := make(BuiltinSet[string])
	assert.Equal(t, s.IsEmpty(), true)
	s.Insert("hello")
	assert.Equal(t, s.IsEmpty(), false)
}

func Test_BuiltinSet_Clear(t *testing.T) {
	s := SetOf("hello", "world")
	s.Clear()
	assert.True(t, s.IsEmpty())
}

func Test_BuiltinSet_Has(t *testing.T) {
	s := SetOf("hello", "world")
	assert.True(t, s.Has("hello"))
	assert.True(t, s.Has("world"))
	assert.False(t, s.Has("!"))
}

func Test_BuiltinSet_Insert(t *testing.T) {
	s := make(BuiltinSet[string])
	assert.True(t, s.Insert("hello"))
	assert.False(t, s.Insert("hello"))
	assert.Equal(t, s.Has("world"), false)
	assert.True(t, s.Insert("world"))
	assert.Equal(t, s.Has("hello"), true)
	assert.Equal(t, s.Len(), 2)
}

func Test_BuiltinSet_InsertN(t *testing.T) {
	s := make(BuiltinSet[string])
	assert.Equal(t, s.InsertN("hello", "world"), 2)
	assert.Equal(t, s.Len(), 2)
}

func Test_BuiltinSet_Remove(t *testing.T) {
	s := SetOf("hello", "world")
	assert.True(t, s.Remove("hello"))
	assert.Equal(t, s.Len(), 1)
	assert.False(t, s.Remove("hello"))
	assert.Equal(t, s.Len(), 1)
	assert.True(t, s.Remove("world"))
	assert.Equal(t, s.Len(), 0)
}

func Test_BuiltinSet_Delete(t *testing.T) {
	s := SetOf("hello", "world")
	s.Delete("hello")
	assert.Equal(t, s.Len(), 1)
	s.Delete("hello")
	assert.Equal(t, s.Len(), 1)
	s.Delete("world")
	assert.Equal(t, s.Len(), 0)
}

func Test_BuiltinSet_RemoveN(t *testing.T) {
	s := SetOf("hello", "world")
	assert.Equal(t, s.RemoveN("hello", "world"), 2)
	assert.False(t, s.Remove("world"))
	assert.True(t, s.IsEmpty())
}

func Test_BuiltinSet_Keys(t *testing.T) {
	s := SetOf("hello", "world")
	ks := s.Keys()
	assert.Equal(t, 2, len(ks))
}

func Test_BuiltinSet_For(t *testing.T) {
	s := SetOf("hello", "world")
	for v := range s {
		assert.True(t, v == "hello" || v == "world")
	}
}

func Test_BuiltinSet_ForEach(t *testing.T) {
	s := SetOf("hello", "world")
	c := 0
	s.ForEach(func(string) {
		c++
	})
	assert.Equal(t, c, 2)
}

func Test_BuiltinSet_ForEachIf(t *testing.T) {
	s := SetOf("hello", "world")
	c := 0
	s.ForEachIf(func(string) bool {
		c++
		return false
	})
	assert.Less(t, c, 2)
}

func Test_BuiltinSet_Update(t *testing.T) {
	s := SetOf(1, 2, 3)
	s.Update(SetOf(3, 4))
	assert.Equal(t, s.Len(), 4)
	assert.True(t, s.Has(4))
}

func Test_BuiltinSet_Union(t *testing.T) {
	s := SetOf(1, 2, 3)
	s2 := s.Union(SetOf(3, 4))
	assert.Equal(t, s2.Len(), 4)
	assert.True(t, s2.Has(4))
}

func Test_BuiltinSet_Intersection(t *testing.T) {
	s := SetOf(1, 2, 3).Intersection(SetOf(3, 4))
	assert.Equal(t, s.Len(), 1)
	assert.True(t, s.Has(3))
	s = SetOf(3, 4).Intersection(SetOf(1, 2, 3))
	assert.Equal(t, s.Len(), 1)
	assert.True(t, s.Has(3))
}

func Test_BuiltinSet_Difference(t *testing.T) {
	s := SetOf(1, 2, 3).Difference(SetOf(3, 4))
	assert.Equal(t, s.Len(), 2)
	assert.True(t, s.Has(1))
	assert.True(t, s.Has(2))
	s = SetOf(1, 2).Difference(SetOf(3, 4))
	assert.Equal(t, s.Len(), 2)
	assert.True(t, s.Has(1))
	assert.True(t, s.Has(2))
}

func Test_BuiltinSet_IsDisjointOf(t *testing.T) {
	s1 := SetOf(1, 2, 3)
	s2 := SetOf(3, 4)
	assert.False(t, s1.IsDisjointOf(s2))
	assert.True(t, s1.IsDisjointOf(SetOf(4, 5)))
}

func Test_BuiltinSet_IsSubsetOf(t *testing.T) {
	assert.True(t, SetOf[int]().IsSubsetOf(SetOf[int]()))
	assert.True(t, SetOf[int]().IsSubsetOf(SetOf(1)))
	assert.True(t, SetOf(1, 2, 3).IsSubsetOf(SetOf(1, 2, 3)))
	assert.True(t, SetOf(1, 2).IsSubsetOf(SetOf(1, 2, 3)))
	assert.False(t, SetOf(1, 2, 3).IsSubsetOf(SetOf(1, 2)))
	assert.False(t, SetOf(1, 2).IsSubsetOf(SetOf(2, 3)))
}

func Test_BuiltinSet_IsSupersetOf(t *testing.T) {
	assert.True(t, SetOf[int]().IsSupersetOf(SetOf[int]()))
	assert.True(t, SetOf(1).IsSupersetOf(SetOf[int]()))
	assert.True(t, SetOf(1, 2, 3).IsSupersetOf(SetOf(1, 2, 3)))
	assert.True(t, SetOf(1, 2, 3).IsSupersetOf(SetOf(1, 2)))
	assert.False(t, SetOf(1, 2).IsSupersetOf(SetOf(1, 2, 3)))
	assert.False(t, SetOf(1, 2).IsSupersetOf(SetOf(2, 3)))
}

func Test_BuiltinSet_Equal(t *testing.T) {
	v1 := SetOf[int](1, 2, 3, 4, 5)
	v2 := SetOf[int](1, 2, 3, 4, 6)
	assert.False(t, v1.Equal(v2))

	v3 := SetOf[int](1, 2, 3, 4, 5)
	assert.True(t, v1.Equal(v3))

	v4 := SetOf[int](1, 2, 3, 4)
	assert.False(t, v1.Equal(v4))
	assert.False(t, v4.Equal(v1))

}
