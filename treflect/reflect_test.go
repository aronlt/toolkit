package treflect

import (
	"reflect"
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

func TestDeepCopySlice(t *testing.T) {
	type V struct {
		Name    string
		Address *string
	}
	addr1 := "addr1"
	addr2 := "addr2"
	v := []V{
		{Name: "v1", Address: &addr1},
		{Name: "v2", Address: &addr2},
	}

	v2 := DeepCopySlice(v)
	assert.True(t, reflect.DeepEqual(v, v2))
	if v[0].Address != v2[0].Address {
		assert.True(t, true)
	} else {
		assert.True(t, false)
	}
}

func TestContainTag(t *testing.T) {
	type V struct {
		Name    string  `json:"name"`
		Address *string `json:"address"`
	}
	addr := "addr"
	v := V{
		Name:    "name",
		Address: &addr,
	}
	assert.True(t, ContainTag(v, "name"))
	assert.False(t, ContainTag(v, "addr"))
	assert.True(t, ContainTag(v, "address"))
}

func TestGetFieldValueToFloat(t *testing.T) {
	type M struct {
		A float32
		B float64
	}

	m := &M{
		A: 10,
		B: 11.1,
	}
	v1, err1 := GetFieldValueToFloat(m, "A")
	assert.Nil(t, err1)
	assert.Equal(t, v1, float64(10))

	v2, err2 := GetFieldValueToFloat(m, "B")
	assert.Nil(t, err2)
	assert.Equal(t, v2, float64(11.1))
}

func TestGetFieldValueToInt(t *testing.T) {
	type M struct {
		A int
		B int16
	}

	m := &M{
		A: 10,
		B: 11,
	}
	v1, err1 := GetFieldValueToInt(m, "A")
	assert.Nil(t, err1)
	assert.Equal(t, v1, int64(10))

	v2, err2 := GetFieldValueToInt(m, "B")
	assert.Nil(t, err2)
	assert.Equal(t, v2, int64(11))
}

func TestGetFieldSpecificValue(t *testing.T) {
	type M struct {
		A float32 `json:"a"`
	}
	type V struct {
		M
		Age     int
		Name    string  `json:"name"`
		Address *string `json:"address"`
		M2      M
	}
	address := "address"
	v := &V{
		Age:     10,
		Name:    "name",
		Address: &address,
		M: M{
			A: 0.1,
		},
		M2: M{
			A: 1.0,
		},
	}
	value1, type1, err1 := GetFieldSpecificValue[string](v, "Name")
	assert.Nil(t, err1)
	assert.Equal(t, type1, reflect.String)
	assert.Equal(t, value1, "name")

	value2, type2, err2 := GetFieldSpecificValue[*string](v, "Address")
	assert.Nil(t, err2)
	assert.Equal(t, type2, reflect.Pointer)
	assert.Equal(t, *value2, "address")

	value3, type3, err3 := GetFieldSpecificValue[float32](v, "A")
	assert.Nil(t, err3)
	assert.Equal(t, type3, reflect.Float32)
	assert.Equal(t, value3, float32(0.1))

	value4, type4, err4 := GetFieldSpecificValue[M](v, "M2")
	assert.Nil(t, err4)
	assert.Equal(t, type4, reflect.Struct)
	assert.Equal(t, value4, M{A: 1.0})

	func() {
		defer func() {
			recover()
		}()
		_, _, err5 := GetFieldSpecificValue[int64](v, "Age")
		assert.NotNil(t, err5)
	}()
}

func TestGetAllFields(t *testing.T) {
	type M struct {
		A float32 `json:"a"`
	}
	type V struct {
		M
		Age     int
		Name    string  `json:"name"`
		Address *string `json:"address"`
		M2      M
	}
	address := "address"
	v := &V{
		Age:     10,
		Name:    "name",
		Address: &address,
		M: M{
			A: 0.1,
		},
		M2: M{
			A: 1.0,
		},
	}

	allFields, err := GetAllFields(v)
	assert.Nil(t, err)
	assert.Equal(t, len(allFields), 5)
}
