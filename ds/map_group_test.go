package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGroupMap(t *testing.T) {
	type V struct {
		key   string
		value string
	}
	data := make([]*V, 0)
	data = append(data, &V{
		key:   "a",
		value: "b",
	}, &V{
		key:   "a",
		value: "b",
	}, &V{
		key:   "c",
		value: "b",
	})
	group := NewGroupMap(data, func(v *V) string {
		return v.key
	})
	assert.Equal(t, len(group["a"]), 2)
	assert.Equal(t, group["c"][0].value, "b")
}
