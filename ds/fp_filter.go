package ds

// FpFilterSlice Looks through each value in the slice, returning a slice of all the values that pass a truth test (predicate).
func FpFilterSlice[T any](a []T, iterate func(a []T, i int) bool) []T {
	return SliceFilter(a, func(i int) bool {
		return iterate(a, i)
	})
}

// FpFilterMap Looks through each value in the map, returning a map of all the values that pass a truth test (predicate).
func FpFilterMap[K comparable, V any](a map[K]V, iterate func(a map[K]V, k K, v V) bool) map[K]V {
	b := make(map[K]V, len(a))
	count := 0
	for k, v := range a {
		if iterate(a, k, v) {
			b[k] = v
			count += 1
		}
	}
	if 2*count < len(a) {
		c := make(map[K]V, count)
		for k, v := range b {
			c[k] = v
		}
		return c
	}
	return b
}

// FpFilterList Looks through each value in the list, returning a list of all the values that pass a truth test (predicate).
func FpFilterList[T any](a List[T], iterate func(a List[T], node T) bool) List[T] {
	iterator := a.Iterate()
	b := List[T]{}
	for iterator.IsNotEnd() {
		v := iterator.Value()
		if iterate(a, iterator.Value()) {
			b.PushBack(v)
		}
		iterator.MoveToNext()
	}
	return a
}

// FpFilterSet Looks through each value in the set, returning a set of all the values that pass a truth test (predicate).
func FpFilterSet[T comparable](a BuiltinSet[T], iterate func(a BuiltinSet[T], node T) bool) BuiltinSet[T] {
	b := NewSet[T](a.Len())
	count := 0
	a.ForEach(func(v T) {
		if iterate(a, v) {
			b.Insert(v)
			count += 1
		}
	})
	if 2*count < len(a) {
		c := NewSet[T](count)
		b.ForEach(func(k T) {
			c.Insert(k)
		})
		return c
	}
	return b
}
