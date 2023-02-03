package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapNativeKeyCompare(t *testing.T) {
	m1 := make(map[string]string, 0)
	m2 := make(map[string]string, 0)
	m1["a"] = "a"
	m1["b"] = "b"
	m1["d"] = "d"

	m2["c"] = "a"
	m2["a"] = "a2"
	m2["b"] = "b"

	assert.Equal(t, MapNativeCompareWithKey(m1, m2, "a"), LeftLessThanRight)
	assert.Equal(t, MapNativeCompareWithKey(m1, m2, "b"), Equal)
	assert.Equal(t, MapNativeCompareWithKey(m1, m2, "c"), LeftKeyMiss)
	assert.Equal(t, MapNativeCompareWithKey(m1, m2, "d"), RightKeyMiss)
	assert.Equal(t, MapNativeCompareWithKey(m1, m2, "e"), AllKeyMiss)
}

func TestMapNativeKeyFullCompare(t *testing.T) {
	m1 := make(map[string]string, 0)
	m2 := make(map[string]string, 0)
	m1["a"] = "a"
	m1["b"] = "b"
	m1["d"] = "d"

	m2["c"] = "a"
	m2["a"] = "a2"
	m2["b"] = "b"

	assert.Equal(t, MapNativeFullCompare(m1, m2), NotEqual)

	m3 := map[string]string{"a": "a", "b": "b"}
	m4 := map[string]string{"a": "a", "b": "b"}
	assert.Equal(t, MapNativeFullCompare(m3, m4), Equal)
}

func TestMapComplexKeyCompare(t *testing.T) {
	type T struct {
		A int
	}
	m1 := map[string]T{
		"a": {
			A: 1,
		},
	}
	m2 := map[string]T{
		"a": {
			A: 3,
		},
	}
	assert.Equal(t, MapComplexCompareWithKey(m1, m2, "a"), NotEqual)

	m3 := map[string]T{
		"a": {
			A: 1,
		},
	}
	m4 := map[string]T{
		"a": {
			A: 1,
		},
	}
	assert.Equal(t, MapComplexCompareWithKey(m3, m4, "a"), Equal)
}

func TestMapComplexKeyFullCompare(t *testing.T) {
	type T struct {
		A int
	}
	m1 := map[string]T{
		"a": {
			A: 1,
		},
	}
	m2 := map[string]T{
		"a": {
			A: 3,
		},
	}
	assert.Equal(t, MapComplexFullCompare(m1, m2), NotEqual)

	m3 := map[string]T{
		"a": {
			A: 1,
		},
	}
	m4 := map[string]T{
		"a": {
			A: 1,
		},
	}
	assert.Equal(t, MapComplexFullCompare(m3, m4), Equal)
}
