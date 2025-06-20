package treflect

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/aronlt/toolkit/ds"
	"github.com/aronlt/toolkit/terror"
)

// GetAnyMapValue 从Map中取出对应值
func GetAnyMapValue[T any](anyMap map[string]any, key string, defaultValue T) T {
	v, ok := anyMap[key]
	if !ok {
		return defaultValue
	}
	v2, ok2 := v.(T)
	if !ok2 {
		return defaultValue
	}
	return v2
}

// ConvertAnyMapToStruct 把任意map转换为目标结构体
func ConvertAnyMapToStruct[T any](anyMap map[string]any) (*T, error) {
	content, err := json.Marshal(anyMap)
	if err != nil {
		return nil, terror.Wrap(err, "call Marshal fail")
	}
	var data T
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, terror.Wrap(err, "call Unmarshal fail")
	}
	return &data, nil
}

func ToAnyMap(item interface{}, skip ...string) map[string]interface{} {
	return StructToAnyMap(item, skip...)
}

// StructToAnyMap 把任意结构体转换为字符串形式的任意map
func StructToAnyMap(item interface{}, skip ...string) map[string]interface{} {
	skipSet := ds.SetFromUnpack(skip...)
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
		if field.Anonymous { // 如果字段是内嵌字段并没有 json 标签，递归提升内嵌字段
			subResult := StructToAnyMap(value.Interface())
			for k, v := range subResult {
				result[k] = v
			}
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

// StructToAnyMapWithJson 把结构体转换为json字符串形式的任意map
func StructToAnyMapWithJson(item interface{}, skip ...string) map[string]interface{} {
	skipSet := ds.SetFromUnpack(skip...)
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
		if tag == "" && field.Anonymous { // 如果字段是内嵌字段并没有 json 标签，递归提升内嵌字段
			subResult := StructToAnyMapWithJson(value.Interface())
			for k, v := range subResult {
				result[k] = v
			}
			continue
		}
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

// StructToAnyMapDeep 把任意结构体数据递归转换为字符串形式的任意map
func StructToAnyMapDeep(item interface{}, skip ...string) map[string]interface{} {
	skipSet := ds.SetFromUnpack(skip...)
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
		if field.Anonymous {
			subResult := StructToAnyMapDeep(value.Interface())
			for k, v := range subResult {
				result[k] = v
			}
			continue
		}
		if value.Kind() == reflect.Struct {
			subResult := StructToAnyMapDeep(value.Interface(), skip...)
			result[field.Name] = subResult
			continue
		}
		if value.Kind() == reflect.Pointer {
			if value.IsNil() {
				result[field.Name] = nil
			} else {
				if value.Elem().Kind() == reflect.Struct {
					subResult := StructToAnyMapDeep(value.Interface(), skip...)
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

// StructToAnyMapWithJsonDeep 把任意结构体数据递归转换为json字符串形式的任意map
func StructToAnyMapWithJsonDeep(item interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	r := reflect.TypeOf(item) // 获取类型
	if r == nil {
		return make(map[string]interface{})
	}
	if r.Kind() == reflect.Pointer {
		r = r.Elem() // 解引用指针获取真实类型
	}

	v := reflect.ValueOf(item)
	if reflect.DeepEqual(v, reflect.Value{}) { // 检查是否为空值
		return make(map[string]interface{})
	}

	if v.Kind() == reflect.Pointer {
		v = v.Elem() // 解引用指针获取真实值
	}

	if v.Kind() != reflect.Struct { // 如果不是结构体，直接返回空结果
		return result
	}

	// 遍历结构体的字段
	n := r.NumField()
	for i := 0; i < n; i++ {
		field := r.Field(i)
		value := v.Field(i)

		// 获取 JSON 标签
		tag := field.Tag.Get("json")
		if tag == "" && field.Anonymous { // 如果字段是内嵌字段并没有 json 标签，递归提升内嵌字段
			subResult := StructToAnyMapWithJsonDeep(value.Interface())
			for k, v := range subResult {
				result[k] = v
			}
			continue
		}
		if tag == "" {
			continue
		}
		if idx := strings.Index(tag, ","); idx != -1 {
			tag = tag[:idx]
		}
		tag = strings.TrimSpace(tag)

		// 处理字段值
		if value.Kind() == reflect.Struct { // 如果是嵌套的结构体，递归处理
			subResult := StructToAnyMapWithJsonDeep(value.Interface())
			result[tag] = subResult
			continue
		}
		if value.Kind() == reflect.Pointer { // 处理指针值
			if value.IsNil() {
				result[tag] = nil
			} else {
				if value.Elem().Kind() == reflect.Struct {
					subResult := StructToAnyMapWithJsonDeep(value.Interface())
					result[tag] = subResult
				} else {
					result[tag] = value.Elem().Interface()
				}
			}
			continue
		}

		// 普通字段直接赋值
		result[tag] = value.Interface()
	}

	return result
}
