package toolkit

import (
	"sort"

	"github.com/aronlt/toolkit/types"
)

// ReverseSlice 转置切片
func ReverseSlice[T any](data []T) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// UniqueSlice 去重切片
func UniqueSlice[T comparable](data []T) []T {
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

// CopySlice 复制切片
func CopySlice[T any](data []T) []T {
	dst := make([]T, len(data))
	copy(dst, data)
	return dst
}

// BinarySearch 二分搜索
func BinarySearch[T types.Ordered](data []T, value T) int {
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
func SliceMax[T types.Ordered](data []T) T {
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
func Max[T types.Ordered](data ...T) T {
	return SliceMax[T](data)
}

// SliceMin 求数组的最小值
func SliceMin[T types.Ordered](data []T) T {
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
func Min[T types.Ordered](data ...T) T {
	return SliceMin[T](data)
}
