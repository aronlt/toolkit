package bytes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestU16ToBytes(t *testing.T) {
	var v uint16 = 30
	b := U16ToBytes(v)
	v2 := BytesToU16(b)
	assert.Equal(t, v, v2)
}

func TestU32ToBytes(t *testing.T) {
	var v uint32 = 30
	b := U32ToBytes(v)
	v2 := BytesToU32(b)
	assert.Equal(t, v, v2)
}

func TestU32SliceToBytes(t *testing.T) {
	v := []uint32{
		1,
		2,
		3,
		4,
	}
	b := U32SliceToBytes(v)
	v2 := BytesToU32Slice(b)
	assert.Equal(t, v, v2)
}

func TestU64ToBytes(t *testing.T) {
	var v uint64 = 30
	b := U64ToBytes(v)
	v2 := BytesToU64(b)
	assert.Equal(t, v, v2)
}

func TestU64SliceToBytes(t *testing.T) {
	v := []uint64{
		1,
		2,
		3,
		4,
	}
	b := U64SliceToBytes(v)
	v2 := BytesToU64Slice(b)
	assert.Equal(t, v, v2)
}