package treflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
