package ds

import (
	"reflect"
	"sync"

	"github.com/aronlt/toolkit/tsort"
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

// MapMerge 合并两个map，如果key重复则以第二个元素中的key为主
func MapMerge[K comparable, V any](m1 map[K]V, m2 map[K]V) map[K]V {
	len1 := len(m1)
	len2 := len(m2)
	m3 := make(map[K]V, SliceUnpackMax(len1, len2))
	for k, v := range m1 {
		m3[k] = v
	}
	for k, v := range m2 {
		m3[k] = v
	}
	return m3
}

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

// MapValueToSlice 提取map的值
func MapValueToSlice[T comparable, V any](a map[T]V) []V {
	data := make([]V, 0, len(a))
	for _, v := range a {
		data = append(data, v)
	}
	return data
}

// MapKeyToSlice 提取map的key
func MapKeyToSlice[T comparable, V any](a map[T]V) []T {
	data := make([]T, 0, len(a))
	for k := range a {
		data = append(data, k)
	}
	return data
}

// MapZipSliceToMap 两个slice，一个key，一个value转换为map
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

// SortedMap 有序Map，底层维护了有序切片
// 如果需要对map进行修改需要执行Rebuild来维护有序行，否则会导致不一致
type SortedMap[K ttypes.Ordered, V any] struct {
	ReverseOpt bool
	Tuples     []ttypes.Tuple[K, V]
	RawMap     map[K]V
}

// Rebuild 重新构建有序Map，一般用于map修改后再次维护tuples的有序行
func (s *SortedMap[K, V]) Rebuild() {
	*s = MapNewSortedMap(s.RawMap, s.ReverseOpt)
}

func MapNewSortedMap[K ttypes.Ordered, V any](data map[K]V, reverseOpts ...bool) SortedMap[K, V] {
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
	tsort.SortSlice(keys, reverseOpt)

	tuples := make([]ttypes.Tuple[K, V], len(keys))
	for i := 0; i < len(keys); i++ {
		tuples[i] = ttypes.Tuple[K, V]{keys[i], data[keys[i]]}
	}
	return SortedMap[K, V]{
		ReverseOpt: reverseOpt,
		Tuples:     tuples,
		RawMap:     data,
	}
}

// MapLocker 带锁的map，用于简单的并发读写场景
type MapLocker[K comparable, V any] struct {
	data   map[K]V
	locker sync.Mutex
}

func MapNewMapLocker[K comparable, V any]() *MapLocker[K, V] {
	return &MapLocker[K, V]{
		data: make(map[K]V, 0),
	}
}

func (m *MapLocker[K, V]) Get(key K) (V, bool) {
	m.locker.Lock()
	defer m.locker.Unlock()
	v, ok := m.data[key]
	return v, ok
}

func (m *MapLocker[K, V]) Contain(key K) bool {
	m.locker.Lock()
	defer m.locker.Unlock()
	_, ok := m.data[key]
	return ok
}

func (m *MapLocker[K, V]) Set(key K, value V) {
	m.locker.Lock()
	defer m.locker.Unlock()
	m.data[key] = value
}

func (m *MapLocker[K, V]) Foreach(handler func(key K, value V)) {
	m.locker.Lock()
	defer m.locker.Unlock()
	for k, v := range m.data {
		handler(k, v)
	}
}

// MapRWLocker 简单的读写锁map
type MapRWLocker[K comparable, V any] struct {
	data   map[K]V
	locker sync.RWMutex
}

func MapNewMapRWLocker[K comparable, V any]() *MapRWLocker[K, V] {
	return &MapRWLocker[K, V]{
		data: make(map[K]V, 0),
	}
}

func (m *MapRWLocker[K, V]) Get(key K) (V, bool) {
	m.locker.RLock()
	defer m.locker.RUnlock()
	v, ok := m.data[key]
	return v, ok
}

func (m *MapRWLocker[K, V]) Contain(key K) bool {
	m.locker.RLock()
	defer m.locker.RUnlock()
	_, ok := m.data[key]
	return ok
}

func (m *MapRWLocker[K, V]) Set(key K, value V) {
	m.locker.Lock()
	defer m.locker.Unlock()
	m.data[key] = value
}

func (m *MapRWLocker[K, V]) Foreach(handler func(key K, value V)) {
	m.locker.RLock()
	defer m.locker.RUnlock()
	for k, v := range m.data {
		handler(k, v)
	}
}
