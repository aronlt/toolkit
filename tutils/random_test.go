package tutils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aronlt/toolkit/ds"
)

func TestRandStringBytesMask(t *testing.T) {
	v := RandStringBytesMask(12)
	t.Logf("%s", v)
}

func TestRandPick(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 12}
	black := ds.SetFromUnpack(5, 6, 8, 1, 2, 1, 3, 4, 11, 12)
	for i := 0; i < 100; i++ {
		v, ok := RandPick(data, black)
		assert.True(t, ok)
		assert.False(t, black.Has(v))
	}

	black = ds.SetFromSlice(data)
	for i := 0; i < 100; i++ {
		_, ok := RandPick(data, black)
		assert.False(t, ok)
	}
}
