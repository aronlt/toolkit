package ds

import (
	"sort"

	"github.com/aronlt/toolkit/ttypes"
)

// SliceFilter 过滤slice
func SliceFilter[T any](a []T, filter func(v T) bool) []T {
	newSlice := make([]T, 0)
	for _, v := range a {
		if filter(v) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// SliceMap 对slice中的元素执行操作
func SliceMap[T any](a []T, handler func(v *T)) {
	for i := range a {
		handler(&a[i])
	}
}

// SliceAbsoluteEqual 判断两个slice是否一样，严格按照顺序比较
func SliceAbsoluteEqual[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] != b[j] {
				return false
			}
		}
	}
	return true
}

// SliceLogicalEqual 判断两个Slice是否逻辑一样，和顺序无关
func SliceLogicalEqual[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	mapA := NewCounterMap(a)
	mapB := NewCounterMap(b)
	return mapA.Equal(mapB)
}

// SliceReverse 转置切片
func SliceReverse[T any](data []T) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// SliceUnique 去重切片
func SliceUnique[T comparable](data []T) []T {
	record := make(map[T]struct{}, len(data))
	result := make([]T, 0, len(data))
	for i := 0; i < len(data); i++ {
		if _, ok := record[data[i]]; !ok {
			record[data[i]] = struct{}{}
			result = append(result, data[i])
		}
	}
	return result
}

// SliceCopy 复制切片
func SliceCopy[T any](data []T) []T {
	dst := make([]T, len(data))
	copy(dst, data)
	return dst
}

// BinarySearch 二分搜索
func BinarySearch[T ttypes.Ordered](data []T, value T) int {
	idx := sort.Search(len(data), func(i int) bool {
		return data[i] >= value
	})

	if idx < len(data) && data[idx] == value {
		return idx
	} else {
		return -1
	}
}

// SliceMax 求数组的最大值
func SliceMax[T ttypes.Ordered](data []T) T {
	var empty T
	if len(data) == 0 {
		return empty
	}
	max := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

// Max 求最大值
func Max[T ttypes.Ordered](data ...T) T {
	return SliceMax[T](data)
}

// SliceMin 求数组的最小值
func SliceMin[T ttypes.Ordered](data []T) T {
	var empty T
	if len(data) == 0 {
		return empty
	}
	min := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] < min {
			min = data[i]
		}
	}
	return min
}

// Min 求最小值
func Min[T ttypes.Ordered](data ...T) T {
	return SliceMin[T](data)
}
