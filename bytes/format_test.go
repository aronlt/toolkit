package bytes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIBytesToString(t *testing.T) {
	size := 10240000
	precision := 1
	v := IBytesToString(uint64(size), precision)
	assert.Equal(t, "9.8MB", v)
}
