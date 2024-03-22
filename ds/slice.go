package ds

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"

	"github.com/aronlt/toolkit/tsort"

	"github.com/aronlt/toolkit/ttypes"
)

// SliceGetOne 从切片中取出第一个元素
func SliceGetOne[T any](data []T) (T, error) {
	if len(data) == 0 {
		var empty T
		return empty, fmt.Errorf("empty slice size")
	}
	return data[0], nil
}

// SliceGetOnlyOne 切片只能有一个元素，并取出
func SliceGetOnlyOne[T any](data []T) (T, error) {
	if len(data) != 1 {
		var empty T
		return empty, fmt.Errorf("slice size must be one, actual is:%d", len(data))
	}
	return data[0], nil
}

// SliceOpMerge 合并两个切片
func SliceOpMerge[T any](first []T, second []T) []T {
	result := make([]T, 0, len(first)+len(second))
	result = append(result, first...)
	result = append(result, second...)
	return result
}

// SliceOpReverse 转置切片
func SliceOpReverse[T any](data []T) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// SliceOpJoinToString 切片合并成字符串
func SliceOpJoinToString[T any](data []T, seps ...string) (string, error) {
	strs, err := SliceConvertToString(data)
	if err != nil {
		return "", err
	}
	sep := ","
	if len(seps) != 0 {
		sep = seps[0]
	}
	return strings.Join(strs, sep), nil
}

// SliceOpReverseCopy 转置切片并复制
func SliceOpReverseCopy[T any](data []T) []T {
	b := make([]T, 0, len(data))
	for i := len(data) - 1; i >= 0; i-- {
		b = append(b, data[i])
	}
	return b
}

func adjustIndex[T any](data []T, i int) int {
	length := len(data)
	if i < 0 && i+length >= 0 {
		i += length
	}
	if i > len(data) {
		i = len(data)
	}
	if i < 0 {
		i = 0
	}
	return i
}

// SliceOpInsert 把元素插入到data的指定位置
func SliceOpInsert[T any](data *[]T, i int, x ...T) {
	s := *data
	i = adjustIndex(s, i)
	total := len(s) + len(x)
	if total <= cap(s) {
		s2 := s[:total]
		copy(s2[i+len(x):], s[i:])
		copy(s2[i:], x)
		*data = s2
		return
	}
	s2 := make([]T, total)
	copy(s2, s[:i])
	copy(s2[i:], x)
	copy(s2[i+len(x):], s[i:])
	*data = s2
	return
}

// SliceOpPopBack 弹出切片最后一个元素
func SliceOpPopBack[T any](data *[]T) (T, bool) {
	s := *data
	if len(s) == 0 {
		var t T
		return t, false
	}
	t := s[len(s)-1]
	*data = s[:len(s)-1]
	return t, true
}

// SliceOpShuffle shuffle 切片
func SliceOpShuffle[T any](data []T) {
	rand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
}

// SliceOpReplace 原地替换元素
func SliceOpReplace[T comparable](data []T, a T, b T) {
	for i := 0; i < len(data); i++ {
		if data[i] == a {
			data[i] = b
		}
	}
}

// SliceOpRemove 原地删除元素
func SliceOpRemove[T comparable](data *[]T, b T) {
	for i := len(*data) - 1; i >= 0; i-- {
		if (*data)[i] == b {
			if i == len(*data)-1 {
				*data = (*data)[:i]
			} else {
				*data = append((*data)[:i], (*data)[i+1:]...)
			}
		}
	}
}

// SliceOpRemoveIndex 删除某个索引位置的切片
func SliceOpRemoveIndex[T any](data *[]T, i int) {
	if i < 0 || i >= len(*data) {
		return
	}
	*data = append((*data)[:i], (*data)[i+1:]...)
}

// SliceOpRemoveRange 删除某个范围内的切片
func SliceOpRemoveRange[T any](data *[]T, i int, j int) {
	if i < 0 || i >= len(*data) {
		return
	}
	if j < 0 || j >= len(*data) {
		return
	}
	if i >= j {
		return
	}
	*data = append((*data)[:i], (*data)[j:]...)
}

// SliceOpRemoveMany 从Slice集合中移除另外一个Slice中的元素
func SliceOpRemoveMany[T comparable](data *[]T, values []T) {
	set := SetOf[T](values...)
	for i := len(*data) - 1; i >= 0; i-- {
		if set.Has((*data)[i]) {
			if i == len(*data)-1 {
				*data = (*data)[:i]
			} else {
				*data = append((*data)[:i], (*data)[i+1:]...)
			}
		}
	}
}

// SliceOpUnique 去重切片
func SliceOpUnique[T comparable](data []T) []T {
	record := NewSet[T](len(data))
	result := make([]T, 0, len(data))
	for i := 0; i < len(data); i++ {
		if !record.Has(data[i]) {
			record.Insert(data[i])
			result = append(result, data[i])
		}
	}
	if len(data) > 2*len(result) {
		squeeze := make([]T, len(result))
		copy(squeeze, result)
		return squeeze
	}
	return result
}

/* Slice 读取
SliceGetTail 获取最后一个元素
SliceGetCopy 浅拷贝Slice
SliceGetDeepCopy 深拷贝Slice
*/

// SliceGetFilter 过滤slice
func SliceGetFilter[T any](a []T, filter func(i int) bool) []T {
	newSlice := make([]T, 0, len(a))
	for i := 0; i < len(a); i++ {
		if filter(i) {
			newSlice = append(newSlice, a[i])
		}
	}
	// 收缩内存
	if len(a) > 2*len(newSlice) {
		newSlice2 := make([]T, len(newSlice))
		copy(newSlice2, newSlice)
		newSlice = newSlice2
	}
	return newSlice
}

// SliceSetTail 设置切片最后一个元素值
func SliceSetTail[T any](data []T, d T) {
	if len(data) == 0 {
		return
	}
	data[len(data)-1] = d
}

// SliceGetTail 获取切片最后一个元素，如果没有则用默认值
func SliceGetTail[T any](data []T, d ...T) T {
	if len(data) == 0 {
		if len(d) > 0 {
			return d[0]
		}
		var t T
		return t
	}
	return data[len(data)-1]
}

// SliceSetNthTail 设置切片倒数第N个元素
func SliceSetNthTail[T any](data []T, nth int, d T) {
	if nth < 0 {
		return
	}
	if nth >= len(data) {
		return
	}

	data[len(data)-nth-1] = d
	return
}

// SliceGetNthTail 获取切片倒数第N个元素，如果没有则用默认值
func SliceGetNthTail[T any](data []T, nth int, d ...T) T {
	var t T
	if len(d) > 0 {
		t = d[0]
	}
	if nth < 0 {
		return t
	}
	if nth >= len(data) {
		return t
	}

	return data[len(data)-nth-1]
}

// SliceGetCopy 复制切片
func SliceGetCopy[T any](data []T, ns ...int) []T {
	if len(ns) > 0 {
		n := ns[0]
		if n <= 0 || n > len(data) {
			return []T{}
		}
		dst := make([]T, n)
		copy(dst, data[:n])
		return dst
	}
	dst := make([]T, len(data))
	copy(dst, data)
	return dst
}

/* Slice 转换
SliceConvertToInt64
SliceConvertToInt
SliceConvertToString
*/

// SliceConvertToInt64 切片集合统一转换为[]int64
func SliceConvertToInt64(data interface{}) ([]int64, error) {
	switch data.(type) {
	case []int:
		oriData := data.([]int)
		result := make([]int64, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			result = append(result, int64(oriData[i]))
		}
		return result, nil
	case []int8:
		oriData := data.([]int8)
		result := make([]int64, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			result = append(result, int64(oriData[i]))
		}
		return result, nil
	case []int16:
		oriData := data.([]int16)
		result := make([]int64, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			result = append(result, int64(oriData[i]))
		}
		return result, nil
	case []int32:
		oriData := data.([]int32)
		result := make([]int64, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			result = append(result, int64(oriData[i]))
		}
		return result, nil
	case []int64:
		oriData := data.([]int64)
		return SliceGetCopy(oriData), nil
	case []uint:
		oriData := data.([]uint)
		result := make([]int64, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			result = append(result, int64(oriData[i]))
		}
		return result, nil
	case []uint8:
		oriData := data.([]uint8)
		result := make([]int64, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			result = append(result, int64(oriData[i]))
		}
		return result, nil
	case []uint16:
		oriData := data.([]uint16)
		result := make([]int64, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			result = append(result, int64(oriData[i]))
		}
		return result, nil
	case []uint32:
		oriData := data.([]uint32)
		result := make([]int64, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			result = append(result, int64(oriData[i]))
		}
		return result, nil
	case []uint64:
		oriData := data.([]uint64)
		result := make([]int64, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			if oriData[i] > math.MaxInt64 {
				return make([]int64, 0), fmt.Errorf("overflow uint64:%d", oriData[i])
			}
			result = append(result, int64(oriData[i]))
		}
		return result, nil
	case []string:
		oriData := data.([]string)
		result := make([]int64, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			v, err := strconv.ParseInt(oriData[i], 10, 64)
			if err != nil {
				return make([]int64, 0), fmt.Errorf("convert string:%s fail, error:%+v", oriData[i], err)
			}
			result = append(result, v)
		}
		return result, nil
	default:
		return make([]int64, 0), fmt.Errorf("unspport convert type")
	}
}

// SliceConvertToInt 切片集合统一转换为[]int
func SliceConvertToInt(data interface{}) ([]int, error) {
	switch data.(type) {
	case []int:
		oriData := data.([]int)
		return SliceGetCopy(oriData), nil
	case []int8:
		oriData := data.([]int8)
		result := make([]int, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			result = append(result, int(oriData[i]))
		}
		return result, nil
	case []int16:
		oriData := data.([]int16)
		result := make([]int, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			result = append(result, int(oriData[i]))
		}
		return result, nil
	case []int32:
		oriData := data.([]int32)
		result := make([]int, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			result = append(result, int(oriData[i]))
		}
		return result, nil
	case []int64:
		oriData := data.([]int64)
		result := make([]int, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			if oriData[i] > math.MaxInt {
				return make([]int, 0), fmt.Errorf("overflow int64:%d", oriData[i])
			}
			result = append(result, int(oriData[i]))
		}
		return result, nil
	case []uint:
		oriData := data.([]uint)
		result := make([]int, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			if oriData[i] > math.MaxInt {
				return make([]int, 0), fmt.Errorf("overflow uint:%d", oriData[i])
			}
			result = append(result, int(oriData[i]))
		}
		return result, nil
	case []uint8:
		oriData := data.([]uint8)
		result := make([]int, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			result = append(result, int(oriData[i]))
		}
		return result, nil
	case []uint16:
		oriData := data.([]uint16)
		result := make([]int, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			result = append(result, int(oriData[i]))
		}
		return result, nil
	case []uint32:
		oriData := data.([]uint32)
		result := make([]int, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			result = append(result, int(oriData[i]))
		}
		return result, nil
	case []uint64:
		oriData := data.([]uint64)
		result := make([]int, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			if oriData[i] > math.MaxInt {
				return make([]int, 0), fmt.Errorf("overflow uint64:%d", oriData[i])
			}
			result = append(result, int(oriData[i]))
		}
		return result, nil
	case []string:
		oriData := data.([]string)
		result := make([]int, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			v, err := strconv.Atoi(oriData[i])
			if err != nil {
				return make([]int, 0), fmt.Errorf("convert string:%s fail, error:%+v", oriData[i], err)
			}
			result = append(result, v)
		}
		return result, nil
	default:
		return make([]int, 0), fmt.Errorf("unspport convert type")
	}
}

// SliceConvertToString 切片集合统一转换为[]string
func SliceConvertToString(data interface{}) ([]string, error) {
	switch data.(type) {
	case []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64:
		ints, err := SliceConvertToInt64(data)
		if err != nil {
			return make([]string, 0), err
		}
		result := make([]string, 0, len(ints))
		for i := 0; i < len(ints); i++ {
			result = append(result, strconv.FormatInt(ints[i], 10))
		}
		return result, nil
	case []string:
		oriData := data.([]string)
		return SliceGetCopy(oriData), nil
	case [][]byte:
		oriData := data.([][]byte)
		result := make([]string, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			result = append(result, string(oriData[i]))
		}
		return result, nil
	case []error:
		oriData := data.([]error)
		result := make([]string, 0, len(oriData))
		for i := 0; i < len(oriData); i++ {
			result = append(result, oriData[i].Error())
		}
		return result, nil
	default:
		return make([]string, 0), fmt.Errorf("unspport convert type")
	}
}

/* Slice 包含判断
	SliceIncludeWithFn 根据用户自定义函数判断是否包含
	SliceInclude 判断元素是否在切片中
	SliceIncludeUnpack 判断元素是否在切片中, 参数解耦
	SliceIncludeIndex 判断元素是否在切片中，返回下标
	SliceIncludeIndexUnpack 判断元素是否在切片中，返回下标, 参数解耦
    SliceIncludeBinarySearch 以二分搜索方式判断判断元素是否在切片中
	SliceExclude 判断元素是否不在切片中
*/

func SliceIncludeWithFn[T comparable](a []T, fn func(a []T, i int) bool) bool {
	for i := 0; i < len(a); i++ {
		if fn(a, i) {
			return true
		}
	}
	return false
}

func SliceIncludeWithFnV2[T comparable](a []T, fn func(int) bool) bool {
	for i := 0; i < len(a); i++ {
		if fn(i) {
			return true
		}
	}
	return false
}

// SliceInclude 判断元素是否在slice中
func SliceInclude[T comparable](a []T, b T) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == b {
			return true
		}
	}
	return false
}

// SliceIncludeUnpack 判断第一个元素是否在后续元素集合中
func SliceIncludeUnpack[T comparable](a T, others ...T) bool {
	if len(others) == 0 {
		return false
	}
	for i := 0; i < len(others); i++ {
		if a == others[i] {
			return true
		}
	}
	return false
}

// SliceIncludeBinarySearch 二分搜索
func SliceIncludeBinarySearch[T ttypes.Ordered](data []T, value T) int {
	idx := sort.Search(len(data), func(i int) bool {
		return data[i] >= value
	})

	if idx < len(data) && data[idx] == value {
		return idx
	} else {
		return -1
	}
}

// SliceIncludeIndex 获取元素在切片中的下标，如果不存在返回-1
func SliceIncludeIndex[T comparable](a []T, b T) int {
	for i := 0; i < len(a); i++ {
		if a[i] == b {
			return i
		}
	}
	return -1
}

// SliceIncludeIndexUnpack 获取元素在切片中的下标，如果不存在返回-1
func SliceIncludeIndexUnpack[T comparable](a T, others ...T) int {
	if len(others) == 0 {
		return -1
	}
	for i := 0; i < len(others); i++ {
		if a == others[i] {
			return i
		}
	}
	return -1
}

// SliceExclude 判断元素是否不在slice中
func SliceExclude[T comparable](a []T, b T) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == b {
			return false
		}
	}
	return true
}

/* Slice比较判断
SliceCmpAbsEqual 判断两个slice是否一样，严格按照顺序比较
SliceCmpLogicEqual 判断两个slice是否逻辑一样，和顺序无关
SliceCmpLogicSub 判断b切片是否是a切片的子集
SliceCmpAbsSub 判断b切片是否是a切片的子集
SliceCmpTwoDiff 获取两个切片的Diff
*/

// SliceCmpAbsEqual 判断两个slice是否一样，严格按照顺序比较
func SliceCmpAbsEqual[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// SliceCmpLogicEqual 判断两个Slice是否逻辑一样，和顺序无关
func SliceCmpLogicEqual[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	mapA := SliceGroupToCounter(a)
	mapB := SliceGroupToCounter(b)
	return MapEqualCounter(mapA, mapB)
}

// SliceCmpTwoDiff 显示两个切片的区别
func SliceCmpTwoDiff[T comparable](a []T, b []T) ([]T, []T) {
	sa := SetFromSlice(a)
	sb := SetFromSlice(b)
	return SetToSlice(sa.Difference(sb)), SetToSlice(sb.Difference(sa))
}

// SliceCmpLogicSub 判断b切片是否是a切片的子集
func SliceCmpLogicSub[T comparable](a []T, b []T) bool {
	if len(a) < len(b) {
		return false
	}
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	ga := SliceGroupToCounter(a)
	gb := SliceGroupToCounter(b)
	for k, v := range gb {
		if v2, ok := ga[k]; !ok || v2 != v {
			return false
		}
	}
	return true
}

// SliceCmpAbsSub 判断b切片是否是a切片的子集，严格比较
func SliceCmpAbsSub[T comparable](a []T, b []T) int {
	if len(a) < len(b) {
		return -1
	}
	if len(a) == 0 || len(b) == 0 {
		return -1
	}
	for i := 0; i < len(a); i++ {
		k := 0
		for j := i; j < len(a) && k < len(b); {
			if a[j] == b[k] {
				j++
				k++
			} else {
				break
			}
		}
		if k == len(b) {
			return i
		}
	}
	return -1
}

/* Slice 最值
	SliceMax 最大值
    SliceMaxUnpack 最大值，参数解耦
	SliceMaxN 返回最大的N个值
	SliceMaxNWithOrder 返回最大的N个值，并按序返回
	SliceMin 最小值
    SliceMinUnpack 最小值，参数解耦
	SliceMinN 返回最小的N个值
	SliceMinNWIthOrder 返回最小的N个值，并按序返回
*/

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

// SliceMaxUnpack 求最大值
func SliceMaxUnpack[T ttypes.Ordered](data ...T) T {
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

// SliceMinUnpack 求最小值
func SliceMinUnpack[T ttypes.Ordered](data ...T) T {
	return SliceMin[T](data)
}

// SliceMaxNWithOrder 按序返回切片中最大的N个元素
func SliceMaxNWithOrder[T ttypes.Ordered](data []T, n int) []T {
	result := SliceMaxN(data, n)
	tsort.SortSlice(result, true)
	return result
}

func PartQuickSort[T ttypes.Ordered](data []T, k int, max bool) []T {
	if len(data) == 0 || k <= 0 {
		return []T{}
	}
	n := len(data)
	if n <= k {
		return data
	}
	if tsort.IsSorted(data) {
		if max {
			return data[n-k:]
		}
		return data[:k]
	}
	pivot := data[0]
	i := 0
	j := n - 1
	for i < j {
		for i < j && data[j] >= pivot {
			j--
		}
		for i < j && data[i] <= pivot {
			i++
		}
		if i != j {
			data[i], data[j] = data[j], data[i]
		}
	}
	if i != 0 {
		data[i], data[0] = data[0], data[i]
	}
	if i == 0 {
		if max {
			return PartQuickSort(data[1:], k, max)
		}
		result := make([]T, 0, k)
		result = append(result, data[0])
		result = append(result, PartQuickSort(data[1:], k-1, max)...)
		return result
	} else if i == n-1 {
		if max {
			result := make([]T, 0, k)
			result = append(result, data[i])
			result = append(result, PartQuickSort(data[:i], k-1, max)...)
			return result
		}
		return PartQuickSort(data[:i], k, max)
	} else {
		if max {
			if n-i >= k {
				return PartQuickSort(data[i:], k, max)
			}
			result := make([]T, 0, k)
			result = append(result, data[i:]...)
			result = append(result, PartQuickSort(data[:i], k-n+i, max)...)
			return result
		}
		if i >= k {
			return PartQuickSort(data[:i], k, max)
		}
		result := make([]T, 0, k)
		result = append(result, data[:i]...)
		result = append(result, PartQuickSort(data[i:], k-i, max)...)
		return result
	}
}

// SliceMaxN 返回切片中最大N个元素
func SliceMaxN[T ttypes.Ordered](data []T, n int) []T {
	if len(data) < n || n <= 0 {
		return []T{}
	}
	if n == 1 {
		return []T{SliceMax(data)}
	}
	tmpData := SliceGetCopy(data)
	if n == len(data) {
		return tmpData
	}
	return PartQuickSort(tmpData, n, true)
}

// SliceMinNWithOrder 有序返回切片中最小的N个元素
func SliceMinNWithOrder[T ttypes.Ordered](data []T, n int) []T {
	result := SliceMinN(data, n)
	tsort.SortSlice(result)
	return result
}

// SliceMinN 获取切片中最小的N个元素
func SliceMinN[T ttypes.Ordered](data []T, n int) []T {
	if len(data) < n || n <= 0 {
		return []T{}
	}
	if n == 1 {
		return []T{SliceMin(data)}
	}
	tmpData := SliceGetCopy(data)
	if n == len(data) {
		return tmpData
	}

	return PartQuickSort(tmpData, n, false)
}

/* Slice分类函数划分
	SliceGroupToMap 结果以map形式返回
	SliceGroupToSlices 结果以二维数组形式返回
	SliceGroupToSet 结果以Set形式放回
    SliceGroupToCounter 结果以计数方式返回
	SliceGroupByHandler 根据V生成K，并按照K的值进行分类
*/

// SliceGroupByHandler 对切片进行分类
func SliceGroupByHandler[K comparable, V any](data []V, getKeyHandler func(int) K) map[K][]V {
	group := make(map[K][]V, len(data))
	for i := 0; i < len(data); i++ {
		key := getKeyHandler(i)
		group[key] = append(group[key], data[i])
	}
	return group
}

// SliceGroupUniqueByHandler 对切片进行分类
func SliceGroupUniqueByHandler[K comparable, V any](data []V, getKeyHandler func(int) K) map[K]V {
	group := make(map[K]V, len(data))
	for i := 0; i < len(data); i++ {
		key := getKeyHandler(i)
		group[key] = data[i]
	}
	return group
}

func SliceGroupToTwoMap[T comparable, K comparable, V any](data []V, groupT func(*V) T, groupK func(*V) K) map[T]map[K][]V {
	result := make(map[T]map[K][]V, len(data))
	for i := 0; i < len(data); i++ {
		t := groupT(&data[i])
		tv, ok := result[t]
		if !ok {
			tv = make(map[K][]V, len(data))
		}
		k := groupK(&data[i])
		tv[k] = append(tv[k], data[i])
		result[t] = tv
	}
	return result
}

// SliceGroupToMap 对切片按照元素进行分类, 结果以map形式返回
func SliceGroupToMap[V comparable](data []V) map[V][]V {
	group := make(map[V][]V, len(data))
	for i := 0; i < len(data); i++ {
		group[data[i]] = append(group[data[i]], data[i])
	}
	return group
}

// SliceGroupByHandlerUnique 对切片按照元素进行分类, 结果以map形式返回
func SliceGroupByHandlerUnique[K comparable, V any](data []V, getKeyHandler func(int) K) map[K]V {
	group := make(map[K]V, len(data))
	for i := 0; i < len(data); i++ {
		key := getKeyHandler(i)
		group[key] = data[i]
	}
	return group
}

// SliceGroupToPartitions 按照指定步长切割Slice
func SliceGroupToPartitions[T any](data []T, step int) [][]T {
	if step <= 0 || step >= len(data) {
		return [][]T{data}
	}
	length := len(data)
	partitions := make([][]T, 0, length/step+1)
	for i := 0; i < length; i += step {
		partitions = append(partitions, data[i:SliceMinUnpack(i+step, length)])
	}
	return partitions
}

// SliceGroupToSlice 切片转换为新的切片，一般用于提取其内部元素
func SliceGroupToSlice[T any, V any](data []T, handler func(int) (V, bool)) []V {
	group := make([]V, 0, len(data))
	for i := 0; i < len(data); i++ {
		if v, ok := handler(i); ok {
			group = append(group, v)
		}
	}
	return group
}

// SliceGroupToSlices 对切片按照元素进行分类，结果以二维数组形式返回
func SliceGroupToSlices[V ttypes.Ordered](data []V) [][]V {
	if !tsort.IsSorted(data) {
		tsort.SortSlice(data)
	}
	result := make([][]V, 0, len(data))
	j := 0
	i := 0
	for i < len(data) {
		j = i + 1
		for j < len(data) {
			if data[j] != data[i] {
				break
			} else {
				j++
			}
		}
		result = append(result, data[i:j])
		i = j
	}
	return result
}

// SliceGroupToCounter 对切片按照元素进行计数
func SliceGroupToCounter[V comparable](data []V) map[V]int {
	counter := make(map[V]int, len(data))
	for i := 0; i < len(data); i++ {
		counter[data[i]] += 1
	}
	return counter
}

// SliceGroupToSet 对切片按照元素进行分类，结果以Set形式返回
func SliceGroupToSet[V comparable](data []V) BuiltinSet[V] {
	set := NewSet[V]()
	for i := 0; i < len(data); i++ {
		set.Insert(data[i])
	}
	return set
}

func MapEqualCounter[V comparable](c map[V]int, other map[V]int) bool {
	for v, counter1 := range c {
		if counter2, ok := other[v]; !ok || counter1 != counter2 {
			return false
		}
	}

	for v, counter1 := range other {
		if counter2, ok := c[v]; !ok || counter1 != counter2 {
			return false
		}
	}
	return true
}
