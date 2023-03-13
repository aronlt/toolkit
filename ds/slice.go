package ds

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"

	"github.com/aronlt/toolkit/tsort"

	"github.com/aronlt/toolkit/ttypes"
)

// SliceIndex 获取元素在切片中的下标，如果不存在返回-1
func SliceIndex[T comparable](a []T, b T) int {
	for i := 0; i < len(a); i++ {
		if a[i] == b {
			return i
		}
	}
	return -1
}

func SliceIndexOrder[T ttypes.Ordered](a []T, b T) int {
	return BinarySearch(a, b)
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

// SliceExclude 判断元素是否不在slice中
func SliceExclude[T comparable](a []T, b T) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == b {
			return false
		}
	}
	return true
}

// SliceFilter 过滤slice
func SliceFilter[T any](a []T, filter func(i int) bool) []T {
	newSlice := make([]T, 0, len(a))
	for i := 0; i < len(a); i++ {
		if filter(i) {
			newSlice = append(newSlice, a[i])
		}
	}
	// 收缩内存
	if len(a) > 2*len(newSlice) {
		newSlice = newSlice[:]
	}
	return newSlice
}

// SliceMap 对slice中的元素执行操作
func SliceMap[T any](a []T, handler func(i int)) {
	for i := 0; i < len(a); i++ {
		handler(i)
	}
}

// SliceAbsoluteEqual 判断两个slice是否一样，严格按照顺序比较
func SliceAbsoluteEqual[T comparable](a []T, b []T) bool {
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

// SliceReverseCopy 转置切片并复制
func SliceReverseCopy[T any](data []T) []T {
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

// SliceInsert 把元素插入到data的指定位置
func SliceInsert[T any](data *[]T, i int, x ...T) {
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

func SliceTail[T any](data []T, d ...T) T {
	if len(data) == 0 {
		if len(d) > 0 {
			return d[0]
		}
		var t T
		return t
	}
	return data[len(data)-1]
}

func SlicePopBack[T any](data *[]T) (T, bool) {
	s := *data
	if len(s) == 0 {
		var t T
		return t, false
	}
	t := s[len(s)-1]
	*data = s[:len(s)-1]
	return t, true
}

// SliceShuffle shuffle 切片
func SliceShuffle[T any](data []T) {
	rand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
}

// SliceReplace 原地替换元素
func SliceReplace[T comparable](data []T, a T, b T) {
	for i := 0; i < len(data); i++ {
		if data[i] == a {
			data[i] = b
		}
	}
}

// SliceRemove 原地删除元素
func SliceRemove[T comparable](data *[]T, b T) {
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

func SliceRemoveIndex[T any](data *[]T, i int) {
	if i < 0 || i >= len(*data) {
		return
	}
	*data = append((*data)[:i], (*data)[i+1:]...)
}

func SliceRemoveRange[T any](data *[]T, i int, j int) {
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
	if len(data) > 2*len(result) {
		result = result[:]
	}
	return result
}

// SliceCopy 复制切片
func SliceCopy[T any](data []T, ns ...int) []T {
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

func MaxNWithOrder[T ttypes.Ordered](data []T, n int) []T {
	result := MaxN(data, n)
	tsort.SortSlice(result, true)
	return result
}

func MaxN[T ttypes.Ordered](data []T, n int) []T {
	if len(data) < n || n <= 0 {
		return []T{}
	}
	if n == 1 {
		return []T{SliceMax(data)}
	}
	tmpData := SliceCopy(data)
	if n == len(data) {
		tsort.SortSlice(tmpData, true)
		return tmpData
	}

	var fastSort func(left, right, k int)
	fastSort = func(left, right, k int) {
		l, r, tmp := left, right, tmpData[left]
		for l < r {
			for l < r && tmp >= tmpData[r] {
				r--
			}
			if l < r {
				tmpData[l] = tmpData[r]
				l++
			}
			for l < r && tmp <= tmpData[l] {
				l++
			}
			if l < r {
				tmpData[r] = tmpData[l]
				r--
			}
		}
		tmpData[l] = tmp
		if k == l-left+1 || k == l-left {
			return
		}
		if k < l-left {
			fastSort(left, l-1, k)
			return
		}
		if k > l-left+1 {
			fastSort(l+1, right, k-(l-left+1))
			return
		}
		return
	}
	fastSort(0, len(tmpData)-1, n)
	return SliceCopy(tmpData, n)
}

func MinNWithOrder[T ttypes.Ordered](data []T, n int) []T {
	result := MinN(data, n)
	tsort.SortSlice(result)
	return result
}

func MinN[T ttypes.Ordered](data []T, n int) []T {
	if len(data) < n || n <= 0 {
		return []T{}
	}
	if n == 1 {
		return []T{SliceMin(data)}
	}
	tmpData := SliceCopy(data)
	if n == len(data) {
		tsort.SortSlice(tmpData)
		return tmpData
	}

	var fastSort func(left, right, k int)
	fastSort = func(left, right, k int) {
		l, r, tmp := left, right, tmpData[left]
		for l < r {
			for l < r && tmp <= tmpData[r] {
				r--
			}
			if l < r {
				tmpData[l] = tmpData[r]
				l++
			}
			for l < r && tmp >= tmpData[l] {
				l++
			}
			if l < r {
				tmpData[r] = tmpData[l]
				r--
			}
		}
		tmpData[l] = tmp
		if k == l-left+1 || k == l-left {
			return
		}
		if k < l-left {
			fastSort(left, l-1, k)
			return
		}
		if k > l-left+1 {
			fastSort(l+1, right, k-(l-left+1))
			return
		}
		return
	}
	fastSort(0, len(tmpData)-1, n)
	return SliceCopy(tmpData, n)
}

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
		return SliceCopy(oriData), nil
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

func SliceConvertToInt(data interface{}) ([]int, error) {
	switch data.(type) {
	case []int:
		oriData := data.([]int)
		return SliceCopy(oriData), nil
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
		return SliceCopy(oriData), nil
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

func Include[T comparable](a T, others ...T) bool {
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

func SliceTwoDiff[T comparable](a []T, b []T) ([]T, []T) {
	sa := SetFromSlice(a)
	sb := SetFromSlice(b)
	return SetToSlice(sa.Difference(sb)), SetToSlice(sb.Difference(sa))
}
