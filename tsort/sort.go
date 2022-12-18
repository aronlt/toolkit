package tsort

import (
	"sort"

	"github.com/aronlt/toolkit/ttypes"
)

// SortSlice 对切片排序, 切片必须是可以比较的类型
func SortSlice[T ttypes.Ordered](data []T, reverseOpts ...bool) {
	if len(reverseOpts) == 0 || !reverseOpts[0] {
		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})
	} else {
		sort.Slice(data, func(i, j int) bool {
			return data[i] > data[j]
		})
	}
}

// SortSliceWithComparator 针对切片的自定义排序
func SortSliceWithComparator[T any](data []T, comparator func(i, j int) bool) {
	sort.Slice(data, comparator)
}

// SortComparator 对实现了比较接口的类型排序
func SortComparator[T ttypes.IComparator](data T) {
	sort.Sort(data)
}

func IsSorted[T ttypes.Ordered](data []T, reverseOpts ...bool) {
	if len(reverseOpts) == 0 || !reverseOpts[0] {
		sort.SliceIsSorted(data, func(i int, j int) bool {
			return data[i] < data[j]
		})
	} else {
		sort.SliceIsSorted(data, func(i int, j int) bool {
			return data[i] > data[j]
		})
	}
}
