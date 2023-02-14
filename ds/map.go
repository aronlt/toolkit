package ds

import (
	"reflect"

	"github.com/aronlt/toolkit/ttypes"
)

type MapCompareResult int

const (
	LeftKeyMiss MapCompareResult = iota + 1
	RightKeyMiss
	AllKeyMiss
	NotEqual
	Equal
	LeftLargeThanRight
	LeftLessThanRight
)

// MapNativeCompareWithKey 简单值的key比较
func MapNativeCompareWithKey[T comparable, V ttypes.Ordered](a map[T]V, b map[T]V, key T) MapCompareResult {
	va, ok1 := a[key]
	vb, ok2 := b[key]
	if !ok1 && !ok2 {
		return AllKeyMiss
	}
	if !ok1 {
		return LeftKeyMiss
	}
	if !ok2 {
		return RightKeyMiss
	}
	if va < vb {
		return LeftLessThanRight
	}
	if va > vb {
		return LeftLargeThanRight
	}
	if va == vb {
		return Equal
	}
	return NotEqual
}

// MapComplexCompareWithKey 复杂值的key比较
func MapComplexCompareWithKey[T comparable, V any](a map[T]V, b map[T]V, key T) MapCompareResult {
	va, ok1 := a[key]
	vb, ok2 := b[key]
	if !ok1 && !ok2 {
		return AllKeyMiss
	}
	if !ok1 {
		return LeftKeyMiss
	}
	if !ok2 {
		return RightKeyMiss
	}
	ok3 := reflect.DeepEqual(va, vb)
	if !ok3 {
		return NotEqual
	}
	return Equal
}

// MapNativeFullCompare 简单值的全部比较
func MapNativeFullCompare[T comparable, V ttypes.Ordered](a map[T]V, b map[T]V) MapCompareResult {
	for ka, va := range a {
		if vb, ok := b[ka]; !ok || vb != va {
			return NotEqual
		}
	}
	for kb, vb := range b {
		if va, ok := a[kb]; !ok || vb != va {
			return NotEqual
		}
	}
	return Equal
}

// MapComplexFullCompare 复杂值的全部比较
func MapComplexFullCompare[T comparable, V any](a map[T]V, b map[T]V) MapCompareResult {
	for ka := range a {
		if _, ok := b[ka]; !ok {
			return NotEqual
		}
	}
	for kb := range b {
		if _, ok := a[kb]; !ok {
			return NotEqual
		}
	}

	ok := reflect.DeepEqual(a, b)
	if ok {
		return Equal
	}
	return NotEqual
}

func MapValueToSlice[T comparable, V any](a map[T]V) []V {
	data := make([]V, 0, len(a))
	for _, v := range a {
		data = append(data, v)
	}
	return data
}

func MapKeyToSlice[T comparable, V any](a map[T]V) []T {
	data := make([]T, 0, len(a))
	for k := range a {
		data = append(data, k)
	}
	return data
}

func MapZipSliceToMap[T comparable, V any](a []T, b []V) (map[T]V, error) {
	result := make(map[T]V, len(a))
	if len(a) != len(b) {
		return result, ttypes.ErrorSliceNotEqualLength
	}
	for i := 0; i < len(a); i++ {
		result[a[i]] = b[i]
	}
	return result, nil
}
