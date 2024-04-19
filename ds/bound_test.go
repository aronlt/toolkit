package ds

import (
	"testing"

	"github.com/aronlt/toolkit/ttypes"
	"github.com/stretchr/testify/assert"
)

func TestLowerBound(t *testing.T) {
	a := []int{1, 2, 4, 5, 5, 6}
	assert.Equal(t, LowerBound(a, 1), 0)
	assert.Equal(t, LowerBound(a, 5), 3)
	assert.Equal(t, LowerBound(a, 7), len(a))
}

func TestLowerBoundFunc(t *testing.T) {
	a := []int{1, 2, 4, 5, 5, 6}
	assert.Equal(t, LowerBoundFunc(a, 1, ttypes.Less[int]), 0)
	assert.Equal(t, LowerBoundFunc(a, 5, ttypes.Less[int]), 3)
	assert.Equal(t, LowerBoundFunc(a, 7, ttypes.Less[int]), len(a))
}

func TestUpperBound(t *testing.T) {
	a := []int{1, 2, 4, 5, 5, 6}
	assert.Equal(t, UpperBound(a, 1), 1)
	assert.Equal(t, UpperBound(a, 5), 5)
	assert.Equal(t, UpperBound(a, 7), len(a))
}

func TestUpperBoundFunc(t *testing.T) {
	a := []int{1, 2, 4, 5, 5, 6}
	assert.Equal(t, UpperBoundFunc(a, 1, ttypes.Less[int]), 1)
	assert.Equal(t, UpperBoundFunc(a, 5, ttypes.Less[int]), 5)
	assert.Equal(t, UpperBoundFunc(a, 7, ttypes.Less[int]), len(a))
}
