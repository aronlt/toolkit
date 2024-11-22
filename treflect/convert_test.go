package treflect

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFloat64(t *testing.T) {
	m, ok := ToFloat64(1)
	assert.True(t, ok)
	assert.Equal(t, float64(1), m)
	m, ok = ToFloat64(1.1)
	assert.True(t, ok)
	assert.Equal(t, 1.1, m)

	m, ok = ToFloat64("1.1")
	assert.True(t, ok)
	assert.Equal(t, 1.1, m)

	m, ok = ToFloat64([]byte("1.1"))
	assert.True(t, ok)
	assert.Equal(t, 1.1, m)

	k := "1.11"
	m, ok = ToFloat64(&k)
	assert.True(t, ok)
	assert.Equal(t, 1.11, m)

	_, ok = ToFloat64(false)
	assert.False(t, ok)

	_, ok = ToFloat64(func() {})
	assert.False(t, ok)

	_, ok = ToFloat64(func() {})
	assert.False(t, ok)

	_, ok = ToFloat64(struct {
		v float64
	}{})
	assert.False(t, ok)

	_, ok = ToFloat64("a")
	assert.False(t, ok)

	_, ok = ToFloat64(nil)
	assert.False(t, ok)

	_, ok = ToFloat64([]int64{1, 2, 3})
	assert.False(t, ok)

	m, ok = ToFloat64(reflect.ValueOf(1.101111))
	assert.True(t, ok)
	assert.Equal(t, 1.101111, m)

	m, ok = ToFloat64(reflect.ValueOf(&k))
	assert.True(t, ok)
	assert.Equal(t, 1.11, m)

	k2 := 1.11
	m, ok = ToFloat64(reflect.ValueOf(&k2))
	assert.True(t, ok)
	assert.Equal(t, 1.11, m)

	m, ok = ToFloat64(reflect.ValueOf(10))
	assert.True(t, ok)
	assert.Equal(t, float64(10), m)
}

func TestToInt64(t *testing.T) {
	m, ok := ToInt64(int64(1))
	assert.True(t, ok)
	assert.Equal(t, int64(1), m)

	m, ok = ToInt64(int(1))
	assert.True(t, ok)
	assert.Equal(t, int64(1), m)

	m, ok = ToInt64("1")
	assert.True(t, ok)
	assert.Equal(t, int64(1), m)

	_, ok = ToInt64("1.1")
	assert.False(t, ok)

	_, ok = ToInt64(1.1)
	assert.False(t, ok)

	m, ok = ToInt64([]byte("1"))
	assert.True(t, ok)
	assert.Equal(t, int64(1), m)

	k := "1"
	m, ok = ToInt64(&k)
	assert.True(t, ok)
	assert.Equal(t, int64(1), m)

	_, ok = ToInt64(false)
	assert.False(t, ok)

	_, ok = ToInt64(func() {})
	assert.False(t, ok)

	_, ok = ToInt64(struct {
		v int64
	}{})
	assert.False(t, ok)

	_, ok = ToInt64("a")
	assert.False(t, ok)

	_, ok = ToInt64(nil)
	assert.False(t, ok)

	_, ok = ToInt64([]int64{1, 2, 3})
	assert.False(t, ok)

	m, ok = ToInt64(reflect.ValueOf(12))
	assert.True(t, ok)
	assert.Equal(t, int64(12), m)

	k2 := 1
	m, ok = ToInt64(reflect.ValueOf(&k2))
	assert.True(t, ok)
	assert.Equal(t, int64(1), m)
}

func TestToString(t *testing.T) {
	m, ok := ToString(1)
	assert.True(t, ok)
	assert.Equal(t, "1", m)

	m, ok = ToString("1")
	assert.True(t, ok)
	assert.Equal(t, "1", m)

	m, ok = ToString([]byte("1"))
	assert.True(t, ok)
	assert.Equal(t, "1", m)

	m, ok = ToString(1.1)
	assert.True(t, ok)
	assert.Equal(t, "1.1", m)

	k := "1"
	m, ok = ToString(&k)
	assert.True(t, ok)
	assert.Equal(t, "1", m)

	_, ok = ToString(false)
	assert.False(t, ok)

	_, ok = ToString(func() {})
	assert.False(t, ok)

	_, ok = ToString(struct {
		v int64
	}{})
	assert.False(t, ok)

	_, ok = ToString(nil)
	assert.False(t, ok)

	_, ok = ToString([]int64{1, 2, 3})
	assert.False(t, ok)

	m, ok = ToString(reflect.ValueOf(12))
	assert.True(t, ok)
	assert.Equal(t, "12", m)

	k2 := 1
	m, ok = ToString(reflect.ValueOf(&k2))
	assert.True(t, ok)
	assert.Equal(t, "1", m)
}
