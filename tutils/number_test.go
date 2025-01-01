package tutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundUp(t *testing.T) {
	assert.Equal(t, uint64(1), RoundUp(1))
	assert.Equal(t, uint64(4), RoundUp(4))
	assert.Equal(t, uint64(8), RoundUp(7))
	assert.Equal(t, uint64(16), RoundUp(15))
	assert.Equal(t, uint64(32), RoundUp(19))
}
