package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceIter(t *testing.T) {
	v := []int{1, 2, 3, 4}
	m2 := func() []int {
		m := make([]int, 0)
		SliceIter(v, func(a []int, i int) {
			m = append(m, a[i])
		})
		return m
	}()
	assert.Equal(t, m2, v)
}
