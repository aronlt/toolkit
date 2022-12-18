package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewStack(t *testing.T) {
	si := NewStack[int]()
	ss := NewStack[string]()
	assert.Equal(t, si.Len(), 0)
	assert.Equal(t, ss.Len(), 0)
}

func Test_NewStackCap(t *testing.T) {
	s := NewStackCap[int](10)
	assert.Equal(t, s.Len(), 0)
	assert.Equal(t, s.Cap(), 10)
}

func Test_StackCap(t *testing.T) {
	s := NewStackCap[int](10)
	s.Push(1)
	assert.Equal(t, s.Len(), 1)
	assert.Equal(t, s.Cap(), 10)
}

func Test_Stack_Clear(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Clear()
	assert.Equal(t, s.Len(), 0)
	assert.True(t, s.IsEmpty())
}

func Test_Stack_Push(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	assert.Equal(t, s.Len(), 1)
	s.Push(2)
	assert.Equal(t, s.Len(), 2)
}

func Test_Stack_TryPop(t *testing.T) {
	s := NewStack[int]()
	_, ok := s.TryPop()
	assert.False(t, ok)
	s.Push(1)
	v, ok := s.TryPop()
	assert.True(t, ok)
	assert.Equal(t, v, 1)
}

func Test_Stack_Pop(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	v := s.Pop()
	assert.Equal(t, v, 1)
	assert.Panics(t, func() { s.Pop() })
}

func Test_Stack_Top(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	v := s.Top()
	assert.Equal(t, v, 1)
	s.Pop()
	assert.Panics(t, func() { s.Top() })
}
