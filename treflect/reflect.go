package treflect

import (
	"fmt"
	"reflect"
	"strings"
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

// ToAnyMap 把任意数据转换为json字符串形式的任意map
func ToAnyMap(item interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	r := reflect.TypeOf(item)
	if r.Kind() == reflect.Pointer {
		r = r.Elem()
	}
	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return result
	}
	n := r.NumField()
	for i := 0; i < n; i++ {
		field := r.FieldByIndex([]int{i})
		value := v.FieldByIndex([]int{i})
		tag := field.Tag.Get("json")
		if tag == "" {
			continue
		}
		if idx := strings.Index(tag, ","); idx != -1 {
			tag = tag[:idx]
		}
		tag = strings.TrimSpace(tag)
		if value.Kind() == reflect.Pointer {
			result[tag] = value.Elem().Interface()
		} else {
			result[tag] = value.Interface()
		}
	}
	return result
}

// ContainTag 判断结构体是否含有特定的json内容
func ContainTag(item interface{}, tag string) bool {
	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return false
	}
	r := reflect.TypeOf(item)
	if r.Kind() == reflect.Pointer {
		r = r.Elem()
	}
	n := r.NumField()
	for i := 0; i < n; i++ {
		field := r.FieldByIndex([]int{i})
		ftag := field.Tag.Get("json")
		if ftag == "" {
			continue
		}
		if idx := strings.Index(ftag, ","); idx != -1 {
			ftag = ftag[:idx]
		}
		ftag = strings.TrimSpace(ftag)
		if ftag == tag {
			return true
		}
	}
	return false
}
