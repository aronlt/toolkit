package treflect

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/aronlt/toolkit/ds"
)

// SetField 修改结构体字段的值
func SetField(item interface{}, fieldName string, value interface{}) error {
	if reflect.TypeOf(item).Kind() != reflect.Pointer {
		return fmt.Errorf("expected pointer type, but accept:%v", reflect.TypeOf(item).Kind())
	}
	data := reflect.ValueOf(item).Elem()
	if data.Kind() != reflect.Struct {
		return fmt.Errorf("invalid elem type")
	}
	field := data.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("can't find field")
	}
	if !field.CanSet() {
		return fmt.Errorf("field name %v is not exported in struct %v", fieldName, data.Type().String())
	}
	if field.Kind() == reflect.Pointer {
		field = field.Elem()
	}
	fType := field.Type()
	vValue := reflect.ValueOf(value)
	if vValue.Type().AssignableTo(fType) {
		field.Set(vValue)
	} else if vValue.CanConvert(fType) {
		field.Set(vValue.Convert(fType))
	}

	return nil
}

func GetFieldValueToFloat(item interface{}, fieldName string) (float64, error) {
	value, err := GetFieldValue(item, fieldName)
	if err != nil {
		return 0.0, fmt.Errorf("call GetFieldValue fail, err:%+v", err)
	}
	switch value.Kind() {
	case reflect.Float32:
		var result float32
		result = value.Interface().(float32)
		return float64(result), nil
	case reflect.Float64:
		var result float64
		result = value.Interface().(float64)
		return result, nil
	default:
		return 0.0, fmt.Errorf("expect type be int, get:%s", value.Kind())
	}
}

func GetFieldValueToInt(item interface{}, fieldName string) (int64, error) {
	value, err := GetFieldValue(item, fieldName)
	if err != nil {
		return -1, fmt.Errorf("call GetFieldValue fail, err:%+v", err)
	}
	switch value.Kind() {
	case reflect.Int:
		var result int
		result = value.Interface().(int)
		return int64(result), nil
	case reflect.Int8:
		var result int8
		result = value.Interface().(int8)
		return int64(result), nil
	case reflect.Int16:
		var result int16
		result = value.Interface().(int16)
		return int64(result), nil
	case reflect.Int32:
		var result int32
		result = value.Interface().(int32)
		return int64(result), nil
	case reflect.Int64:
		var result int64
		result = value.Interface().(int64)
		return result, nil
	case reflect.Uint:
		var result uint
		result = value.Interface().(uint)
		return int64(result), nil
	case reflect.Uint8:
		var result uint8
		result = value.Interface().(uint8)
		return int64(result), nil
	case reflect.Uint16:
		var result uint16
		result = value.Interface().(uint16)
		return int64(result), nil
	case reflect.Uint32:
		var result uint32
		result = value.Interface().(uint32)
		return int64(result), nil
	case reflect.Uint64:
		var result uint64
		result = value.Interface().(uint64)
		return int64(result), nil
	default:
		return -1, fmt.Errorf("expect type be int, get:%s", value.Kind())
	}
}

// GetFieldSpecificValue 取出Struct的Field值, 传入精确的类型
func GetFieldSpecificValue[T any](item interface{}, fieldName string) (T, reflect.Kind, error) {
	value, err := GetFieldValue(item, fieldName)
	if err != nil {
		var empty T
		return empty, reflect.Invalid, fmt.Errorf("call GetFieldValue fail, err:%+v", err)
	}

	var result T
	result = value.Interface().(T)
	return result, value.Kind(), nil
}

func GetFieldValue(item interface{}, fieldName string) (reflect.Value, error) {
	r := reflect.ValueOf(item)
	f := reflect.Indirect(r).FieldByName(fieldName)
	if !f.IsValid() {
		return f, fmt.Errorf("can't find field name")
	}
	return f, nil
}

// GetAllFields 取出全部Field字段
func GetAllFields(item interface{}) ([]*reflect.StructField, error) {
	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("param must be struct")
	}
	r := reflect.TypeOf(item)
	if r.Kind() == reflect.Pointer {
		r = r.Elem()
	}
	n := r.NumField()

	var result []*reflect.StructField
	for i := 0; i < n; i++ {
		field := r.FieldByIndex([]int{i})
		result = append(result, &field)
	}
	return result, nil
}

// ToAnyMapDeep 把任意数据递归转换为字符串形式的任意map
func ToAnyMapDeep(item interface{}, skip ...string) map[string]interface{} {
	skipSet := ds.SetOf(skip...)
	result := make(map[string]interface{})
	r := reflect.TypeOf(item)
	// item is nil
	if r == nil {
		return make(map[string]interface{})
	}
	if r.Kind() == reflect.Pointer {
		r = r.Elem()
	}
	v := reflect.ValueOf(item)
	if reflect.DeepEqual(v, reflect.Value{}) {
		return make(map[string]interface{})
	}
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
		if skipSet.Has(field.Name) {
			continue
		}
		if value.Kind() == reflect.Struct {
			subResult := ToAnyMapDeep(value.Interface(), skip...)
			result[field.Name] = subResult
			continue
		}
		if value.Kind() == reflect.Pointer {
			if value.IsNil() {
				result[field.Name] = nil
			} else {
				if value.Elem().Kind() == reflect.Struct {
					subResult := ToAnyMapDeep(value.Interface(), skip...)
					result[field.Name] = subResult
				} else {
					result[field.Name] = value.Elem().Interface()
				}
			}
		} else {
			result[field.Name] = value.Interface()
		}
	}
	return result
}

// ToAnyMapWithJsonDeep 把任意数据递归转换为json字符串形式的任意map
func ToAnyMapWithJsonDeep(item interface{}, skip ...string) map[string]interface{} {
	skipSet := ds.SetOf(skip...)
	result := make(map[string]interface{})
	r := reflect.TypeOf(item)
	if r == nil {
		return make(map[string]interface{})
	}
	if r.Kind() == reflect.Pointer {
		r = r.Elem()
	}
	v := reflect.ValueOf(item)
	if reflect.DeepEqual(v, reflect.Value{}) {
		return make(map[string]interface{})
	}
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
		if skipSet.Has(tag) {
			continue
		}
		if value.Kind() == reflect.Struct {
			subResult := ToAnyMapWithJsonDeep(value.Interface(), skip...)
			result[tag] = subResult
			continue
		}
		if value.Kind() == reflect.Pointer {
			if value.IsNil() {
				result[tag] = nil
			} else {
				if value.Elem().Kind() == reflect.Struct {
					subResult := ToAnyMapWithJsonDeep(value.Interface(), skip...)
					result[tag] = subResult
				} else {
					result[tag] = value.Elem().Interface()
				}
			}
		} else {
			result[tag] = value.Interface()
		}
	}
	return result
}

// ToAnyMap 把任意数据转换为字符串形式的任意map
func ToAnyMap(item interface{}, skip ...string) map[string]interface{} {
	skipSet := ds.SetOf(skip...)
	result := make(map[string]interface{})
	r := reflect.TypeOf(item)
	// item is nil
	if r == nil {
		return make(map[string]interface{})
	}
	if r.Kind() == reflect.Pointer {
		r = r.Elem()
	}
	v := reflect.ValueOf(item)
	if reflect.DeepEqual(v, reflect.Value{}) {
		return make(map[string]interface{})
	}
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
		if skipSet.Has(field.Name) {
			continue
		}
		if value.Kind() == reflect.Pointer {
			if value.IsNil() {
				result[field.Name] = nil
			} else {
				result[field.Name] = value.Elem().Interface()
			}
		} else {
			result[field.Name] = value.Interface()
		}
	}
	return result
}

// ToAnyMapWithJson 把任意数据转换为json字符串形式的任意map
func ToAnyMapWithJson(item interface{}, skip ...string) map[string]interface{} {
	skipSet := ds.SetOf(skip...)
	result := make(map[string]interface{})
	r := reflect.TypeOf(item)
	if r == nil {
		return make(map[string]interface{})
	}
	if r.Kind() == reflect.Pointer {
		r = r.Elem()
	}
	v := reflect.ValueOf(item)
	if reflect.DeepEqual(v, reflect.Value{}) {
		return make(map[string]interface{})
	}
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
		if skipSet.Has(tag) {
			continue
		}
		if value.Kind() == reflect.Pointer {
			if value.IsNil() {
				result[tag] = nil
			} else {
				result[tag] = value.Elem().Interface()
			}
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

// GetFieldTag 获取StructField对应的tag值
func GetFieldTag(field *reflect.StructField, tagType string) (string, bool) {
	ftag := field.Tag.Get(tagType)
	if ftag == "" {
		return "", false
	}
	if idx := strings.Index(ftag, ","); idx != -1 {
		ftag = ftag[:idx]
	}
	ftag = strings.TrimSpace(ftag)
	return ftag, true
}

// DeepCopySlice 深度复制切片
func DeepCopySlice[T any](data []T, ns ...int) []T {
	if len(ns) > 0 {
		n := ns[0]
		if n <= 0 || n > len(data) {
			return []T{}
		}
		return Copy(data[:n]).([]T)
	}
	return Copy(data).([]T)
}
