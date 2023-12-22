package treflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAnyMap(t *testing.T) {
	m := make(map[string]interface{})
	m["1"] = 1
	m["2"] = "2"
	v := GetAnyMapValue[int](m, "1", -1)
	assert.Equal(t, 1, v)
	v = GetAnyMapValue[int](m, "3", -1)
	assert.Equal(t, -1, v)

	v2 := GetAnyMapValue[string](m, "2", "")
	assert.Equal(t, "2", v2)

	v3 := GetAnyMapValue[bool](m, "2", false)
	assert.Equal(t, false, v3)
}

func TestConvertAnyMapToStruct(t *testing.T) {
	type M struct {
		A string `json:"a"`
		B int    `json:"b"`
		C bool   `json:"c"`
	}
	d := make(map[string]interface{})
	d["a"] = "1"
	d["b"] = 1
	d["c"] = true

	m, err := ConvertAnyMapToStruct[M](d)
	assert.Nil(t, err)
	assert.Equal(t, &M{
		A: "1",
		B: 1,
		C: true,
	}, m)

}
