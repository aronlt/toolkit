package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNamedTuple2E(t *testing.T) {
	tuple := NewNamedTuple2E[string, int]("a", "a", "b", 1)
	assert.Equal(t, tuple.Name(0), "a")
	assert.Equal(t, tuple.Name(1), "b")
	assert.Equal(t, tuple.Index("a"), 0)
	assert.Equal(t, tuple.Index("b"), 1)
	assert.Equal(t, tuple.Get("a").(string), "a")
	assert.Equal(t, tuple.Get("b").(int), 1)
	a, b := tuple.Unpack()
	assert.Equal(t, a, "a")
	assert.Equal(t, b, 1)
}

func TestNewNamedTuple3E(t *testing.T) {
	tuple := NewNamedTuple3E[string, int, int]("a", "a", "b", 1, "c", 2)
	assert.Equal(t, tuple.Name(0), "a")
	assert.Equal(t, tuple.Name(1), "b")
	assert.Equal(t, tuple.Name(2), "c")
	assert.Equal(t, tuple.Index("a"), 0)
	assert.Equal(t, tuple.Index("b"), 1)
	assert.Equal(t, tuple.Index("c"), 2)
	assert.Equal(t, tuple.Get("a").(string), "a")
	assert.Equal(t, tuple.Get("b").(int), 1)
	assert.Equal(t, tuple.Get("c").(int), 2)
	a, b, c := tuple.Unpack()
	assert.Equal(t, a, "a")
	assert.Equal(t, b, 1)
	assert.Equal(t, c, 2)
}

func TestNewNamedTuple4E(t *testing.T) {
	tuple := NewNamedTuple4E[string, int, int, int]("a", "a", "b", 1, "c", 2, "d", 3)
	assert.Equal(t, tuple.Name(0), "a")
	assert.Equal(t, tuple.Name(1), "b")
	assert.Equal(t, tuple.Name(2), "c")
	assert.Equal(t, tuple.Name(3), "d")

	assert.Equal(t, tuple.Index("a"), 0)
	assert.Equal(t, tuple.Index("b"), 1)
	assert.Equal(t, tuple.Index("c"), 2)
	assert.Equal(t, tuple.Index("d"), 3)

	assert.Equal(t, tuple.Get("a").(string), "a")
	assert.Equal(t, tuple.Get("b").(int), 1)
	assert.Equal(t, tuple.Get("c").(int), 2)
	assert.Equal(t, tuple.Get("d").(int), 3)

	a, b, c, d := tuple.Unpack()
	assert.Equal(t, a, "a")
	assert.Equal(t, b, 1)
	assert.Equal(t, c, 2)
	assert.Equal(t, d, 3)
}

func TestNewNamedTuple5E(t *testing.T) {
	tuple := NewNamedTuple5E[string, int, int, int, int]("a", "a", "b", 1, "c", 2, "d", 3, "e", 4)
	assert.Equal(t, tuple.Name(0), "a")
	assert.Equal(t, tuple.Name(1), "b")
	assert.Equal(t, tuple.Name(2), "c")
	assert.Equal(t, tuple.Name(3), "d")
	assert.Equal(t, tuple.Name(4), "e")

	assert.Equal(t, tuple.Index("a"), 0)
	assert.Equal(t, tuple.Index("b"), 1)
	assert.Equal(t, tuple.Index("c"), 2)
	assert.Equal(t, tuple.Index("d"), 3)
	assert.Equal(t, tuple.Index("e"), 4)

	assert.Equal(t, tuple.Get("a").(string), "a")
	assert.Equal(t, tuple.Get("b").(int), 1)
	assert.Equal(t, tuple.Get("c").(int), 2)
	assert.Equal(t, tuple.Get("d").(int), 3)
	assert.Equal(t, tuple.Get("e").(int), 4)

	a, b, c, d, e := tuple.Unpack()
	assert.Equal(t, a, "a")
	assert.Equal(t, b, 1)
	assert.Equal(t, c, 2)
	assert.Equal(t, d, 3)
	assert.Equal(t, e, 4)
}
