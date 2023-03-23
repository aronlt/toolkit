package ds

// FpMapSlice Produces a new slice by mapping each value in list through a transformation function (iterate).
func FpMapSlice[T any](a []T, iterate func(a []T, i int) T) []T {
	b := make([]T, 0, len(a))
	for i := 0; i < len(a); i++ {
		v := iterate(a, i)
		b = append(b, v)
	}
	return b
}

// FpMapMap Produces a new map by mapping each key, value in map through a transformation function (iterate).
func FpMapMap[K comparable, V any](a map[K]V, iterate func(a map[K]V, k K, v V) (K, V)) map[K]V {
	b := make(map[K]V, len(a))
	for k, v := range a {
		k2, v2 := iterate(a, k, v)
		b[k2] = v2
	}
	return b
}

// FpMapList Produces a new list by mapping each value in list through a transformation function (iterate).
func FpMapList[T any](a List[T], iterate func(a List[T], node T) T) List[T] {
	b := List[T]{}
	iterator := a.Iterate()
	for iterator.IsNotEnd() {
		b.PushBack(iterate(a, iterator.Value()))
		iterator.MoveToNext()
	}
	return b
}

// FpMapSet Produces a new list by mapping each value in list through a transformation function (iterate).
func FpMapSet[T comparable](a BuiltinSet[T], iterate func(a BuiltinSet[T], node T) T) BuiltinSet[T] {
	b := NewSet[T]()
	a.ForEach(func(v T) {
		v2 := iterate(a, v)
		b.Insert(v2)
	})
	return b
}