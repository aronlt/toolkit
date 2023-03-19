package ds

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartial(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int) string {
		return strconv.Itoa(int(x) + y)
	}
	f := FpPartial(add, 5)
	is.Equal("15", f(10))
	is.Equal("0", f(-5))
}

func TestPartial1(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int) string {
		return strconv.Itoa(int(x) + y)
	}
	f := FpPartial1(add, 5)
	is.Equal("15", f(10))
	is.Equal("0", f(-5))
}

func TestPartial2(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int, z int) string {
		return strconv.Itoa(int(x) + y + z)
	}
	f := FpPartial2(add, 5)
	is.Equal("24", f(10, 9))
	is.Equal("8", f(-5, 8))
}

func TestPartial3(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int, z int, a float32) string {
		return strconv.Itoa(int(x) + y + z + int(a))
	}
	f := FpPartial3(add, 5)
	is.Equal("21", f(10, 9, -3))
	is.Equal("15", f(-5, 8, 7))
}

func TestPartial4(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int, z int, a float32, b int32) string {
		return strconv.Itoa(int(x) + y + z + int(a) + int(b))
	}
	f := FpPartial4(add, 5)
	is.Equal("21", f(10, 9, -3, 0))
	is.Equal("14", f(-5, 8, 7, -1))
}

func TestPartial5(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	add := func(x float64, y int, z int, a float32, b int32, c int) string {
		return strconv.Itoa(int(x) + y + z + int(a) + int(b) + c)
	}
	f := FpPartial5(add, 5)
	is.Equal("26", f(10, 9, -3, 0, 5))
	is.Equal("21", f(-5, 8, 7, -1, 7))
}
