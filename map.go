package toolkit

import (
	"github.com/aronlt/toolkit/types"
)

// Tuple 有序元组，用于map的平铺，方便排序
type Tuple[K types.Ordered, V any] struct {
	Key   K
	Value V
}

// SortedMap 有序Map，底层维护了有序切片
// 如果需要对map进行修改需要执行Rebuild来维护有序行，否则会导致不一致
type SortedMap[K types.Ordered, V any] struct {
	ReverseOpt bool
	Tuples     []Tuple[K, V]
	RawMap     map[K]V
}

// Rebuild 重新构建有序Map，一般用于map修改后再次维护tuples的有序行
func (s *SortedMap[K, V]) Rebuild() {
	*s = NewSortedMap(s.RawMap, s.ReverseOpt)
}

func NewSortedMap[K types.Ordered, V any](data map[K]V, reverseOpts ...bool) SortedMap[K, V] {
	keys := make([]K, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}

	var reverseOpt bool
	if len(reverseOpts) == 0 || !reverseOpts[0] {
		reverseOpt = false
	} else {
		reverseOpt = true
	}
	SortSlice(keys, reverseOpt)

	tuples := make([]Tuple[K, V], len(keys))
	for i := 0; i < len(keys); i++ {
		tuples[i] = Tuple[K, V]{keys[i], data[keys[i]]}
	}
	return SortedMap[K, V]{
		ReverseOpt: reverseOpt,
		Tuples:     tuples,
		RawMap:     data,
	}
}
