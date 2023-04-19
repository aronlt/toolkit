package treflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFieldValue(t *testing.T) {
	type V struct {
		Name    string
		address string
	}
	v := V{
		Name:    "name",
		address: "address",
	}
	rv, err := GetFieldValue(v, "Name")
	assert.Nil(t, err)
	assert.Equal(t, rv.String(), "name")

	rv, err = GetFieldValue(v, "address")
	assert.Nil(t, err)
	assert.Equal(t, rv.String(), "address")
}

func TestSetField(t *testing.T) {

	type V struct {
		Name    string
		Address *string
	}
	address := "address"
	v := V{
		Name:    "name",
		Address: &address,
	}
	err := SetField(&v, "Name", "name_two")
	assert.Nil(t, err)
	assert.Equal(t, "name_two", v.Name)
	err = SetField(v, "Name", "name_two")
	assert.NotNil(t, err)

	m := 10
	err = SetField(&m, "Name", "name_two")
	assert.NotNil(t, err)

	err = SetField(&v, "Address", "address_two")
	assert.Nil(t, err)
	assert.Equal(t, *v.Address, "address_two")
}

func TestToAnyMapWithJson(t *testing.T) {
	type M struct {
		Age int `json:"age"`
	}
	type V struct {
		M       `json:"m"`
		Name    string  `json:"name,omitempty"`
		Address *string `json:"address"`
	}
	address := "address"
	v := V{
		M: M{
			Age: 10,
		},
		Name:    "name",
		Address: &address,
	}
	m := ToAnyMapWithJson(v, "name")
	anyMap := map[string]interface{}{
		"address": "address",
		"age":     10,
		"m":       M{Age: 10},
	}
	assert.Equal(t, m, anyMap)
}

func TestToAnyMap(t *testing.T) {
	type M struct {
		Age int
	}
	type V struct {
		M
		Name    string
		Address *string
	}
	address := "address"
	v := V{
		M: M{
			Age: 10,
		},
		Name:    "name",
		Address: &address,
	}
	m := ToAnyMap(v, "Name")
	anyMap := map[string]interface{}{
		"Address": "address",
		"M":       M{Age: 10},
	}
	assert.Equal(t, m, anyMap)

	type V2 struct {
		Age    int
		Detail []struct {
			Name    string
			Address string
		}
	}

	v2 := V2{
		Age: 10,
		Detail: []struct {
			Name    string
			Address string
		}{
			{
				Name:    "a",
				Address: "b",
			},
		},
	}

	m2 := ToAnyMap(v2)
	t.Logf("%+v", m2)

	type V3 struct {
		M       *M
		Name    string
		Address *string
	}

	v3 := V3{
		M:       nil,
		Name:    "a",
		Address: nil,
	}
	m3 := ToAnyMap(v3)
	t.Logf("%+v", m3)

	v4 := V3{
		M:       &M{Age: 10},
		Name:    "name",
		Address: &address,
	}
	m4 := ToAnyMapDeep(v4)
	t.Logf("%+v", m4)

	v5 := V3{
		M:       nil,
		Name:    "name",
		Address: &address,
	}
	m5 := ToAnyMapDeep(v5)
	t.Logf("%+v", m5)
}
