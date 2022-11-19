package treflect

import (
	"fmt"
	"reflect"
)

func SetField(item interface{}, fieldName string, value interface{}) error {
	if reflect.TypeOf(item).Kind() != reflect.Pointer {
		return fmt.Errorf("expected pointer type, but accept:%v", reflect.TypeOf(item).Kind())
	}
	elem := reflect.ValueOf(item).Elem()
	if elem.Kind() != reflect.Struct {
		return fmt.Errorf("invalid elem type")
	}
	v := elem.FieldByName(fieldName)
	if !v.IsValid() {
		return fmt.Errorf("can't find field")
	}
	if !v.CanSet() {
		return fmt.Errorf("field name %v is not exported in struct %v", fieldName, elem.Type().String())
	}
	if v.Kind() == reflect.Pointer {
		m := v.Elem()
		m.Set(reflect.ValueOf(value))
	} else {
		v.Set(reflect.ValueOf(value))
	}
	return nil
}

func GetFieldValue(item interface{}, fieldName string) (reflect.Value, error) {
	r := reflect.ValueOf(item)
	f := reflect.Indirect(r).FieldByName(fieldName)
	if !f.IsValid() {
		return f, fmt.Errorf("can't find field name")
	}
	return f, nil
}
