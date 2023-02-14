package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTuple(t *testing.T) {
	t2 := NewTuple2E("a", "b")
	assert.Equal(t, t2.E1, "a")
	assert.Equal(t, t2.E2, "b")

	t6 := NewTuple6E("a", "b", "c", 1, 2, 3)
	assert.Equal(t, t6.E1, "a")
	assert.Equal(t, t6.E6, 3)

	a, b, c, d, e, f := t6.Unpack()

	assert.Equal(t, a, "a")
	assert.Equal(t, b, "b")
	assert.Equal(t, c, "c")
	assert.Equal(t, d, 1)
	assert.Equal(t, e, 2)
	assert.Equal(t, f, 3)
}
