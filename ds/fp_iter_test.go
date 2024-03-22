package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceIterV2(t *testing.T) {
	v := []int{1, 2, 3, 4}
	m2 := func() []int {
		m := make([]int, 0)
		SliceIterV2(v, func(i int) {
			m = append(m, v[i])
		})
		return m
	}()
	assert.Equal(t, m2, v)
}

func TestSliceIter(t *testing.T) {
	v := []int{1, 2, 3, 4}
	m2 := func() []int {
		m := make([]int, 0)
		SliceIter(v, func(v []int, i int) {
			m = append(m, v[i])
		})
		return m
	}()
	assert.Equal(t, m2, v)
}
