package treflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopy(t *testing.T) {
	type V struct {
		M []int
		K []map[string][]int
	}

	v1 := V{
		M: []int{1, 2, 3, 4},
		K: []map[string][]int{
			{"a": []int{1, 2, 3}},
			{"b": []int{1, 2, 3}},
		},
	}
	v2 := Copy(v1).(V)
	assert.Equal(t, v2.M[0], 1)

	assert.Equal(t, v2.K[0]["a"][1], 2)
	assert.Equal(t, v2.K[1]["b"][0], 1)
}
