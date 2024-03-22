package treflect

import (
	"fmt"
	"reflect"
	"strings"
)

func VerifyField(item interface{}, fieldNames []string) error {
	v := reflect.ValueOf(item)
	v, err := checkStruct(v)
	if err != nil {
		return err
	}

	for _, fieldName := range fieldNames {
		f := v.FieldByName(fieldName)
		if !f.IsValid() {
			return fmt.Errorf("can't find field name")
		}

		if f.Kind() == reflect.Pointer {
			if f.IsNil() {
				return fmt.Errorf("field is nil")
			}
			f = f.Elem()
		}
		switch f.Kind() {
		case reflect.Int:
			if result := f.Interface().(int); result <= 0 {
				return fmt.Errorf("invalid field:%s, value:%v", fieldName, result)
			}
		case reflect.Int8:
			if result := f.Interface().(int8); result <= 0 {
				return fmt.Errorf("invalid field:%s, value:%v", fieldName, result)
			}
		case reflect.Int16:
			if result := f.Interface().(int16); result <= 0 {
				return fmt.Errorf("invalid field:%s, value:%v", fieldName, result)
			}
		case reflect.Int32:
			if result := f.Interface().(int32); result <= 0 {
				return fmt.Errorf("invalid field:%s, value:%v", fieldName, result)
			}
		case reflect.Int64:
			if result := f.Interface().(int64); result <= 0 {
				return fmt.Errorf("invalid field:%s, value:%v", fieldName, result)
			}
		case reflect.Uint:
			if result := f.Interface().(uint); result <= 0 {
				return fmt.Errorf("invalid field:%s, value:%v", fieldName, result)
			}
		case reflect.Uint8:
			if result := f.Interface().(uint); result <= 0 {
				return fmt.Errorf("invalid field:%s, value:%v", fieldName, result)
			}
		case reflect.Uint16:
			if result := f.Interface().(uint); result <= 0 {
				return fmt.Errorf("invalid field:%s, value:%v", fieldName, result)
			}

		case reflect.Uint32:
			if result := f.Interface().(uint32); result <= 0 {
				return fmt.Errorf("invalid field:%s, value:%v", fieldName, result)
			}

		case reflect.Uint64:
			if result := f.Interface().(uint64); result <= 0 {
				return fmt.Errorf("invalid field:%s, value:%v", fieldName, result)
			}
		case reflect.String:
			if result := f.Interface().(string); result == "" {
				return fmt.Errorf("invalid field:%s, value:%v", fieldName, result)
			}
		default:
			return fmt.Errorf("not support check this kind")
		}
	}
	return nil
}

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

func checkStruct(value reflect.Value) (reflect.Value, error) {
	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return value, fmt.Errorf("param must be struct")
	}
	return value, nil
}

func GetFieldValue(item interface{}, fieldName string) (reflect.Value, error) {
	r := reflect.ValueOf(item)
	r, err := checkStruct(r)
	if err != nil {
		return r, err
	}
	f := reflect.Indirect(r).FieldByName(fieldName)
	if !f.IsValid() {
		return f, fmt.Errorf("can't find field name")
	}
	return f, nil
}

// GetAllFields 取出全部Field字段
func GetAllFields(item interface{}) ([]*reflect.StructField, error) {
	v := reflect.ValueOf(item)
	v, err := checkStruct(v)
	if err != nil {
		return nil, err
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

// ContainTag 判断结构体是否含有特定的json内容
func ContainTag(item interface{}, tag string) bool {
	v := reflect.ValueOf(item)
	v, err := checkStruct(v)
	if err != nil {
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
