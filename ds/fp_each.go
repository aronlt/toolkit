package ds

// FpEachSlice Iterates over a slice of elements, yielding each in turn to an iterate function, Returns the slice for chaining.
func FpEachSlice[T any](a []T, iterate func(a []T, i int) T) []T {
	for i := 0; i < len(a); i++ {
		a[i] = iterate(a, i)
	}
	return a
}

// FpEachMap Iterates over map, yielding each key, value in turn to an iterate function, Returns the map for chaining.
func FpEachMap[K comparable, V any](a map[K]V, iterate func(a map[K]V, k K, v V) V) map[K]V {
	for k, v := range a {
		v2 := iterate(a, k, v)
		a[k] = v2
	}
	return a
}

// FpEachList Iterates over list, yielding each value in turn to an iterate function, Returns the list for chaining.
func FpEachList[T any](a DList[T], iterate func(a DList[T], node T) T) DList[T] {
	iterator := a.Iterate()
	for iterator.IsNotEnd() {
		p := iterator.Pointer()
		*p = iterate(a, iterator.Value())
		iterator.MoveToNext()
	}
	return a
}

// FpEachSet Iterates over set, yielding each value in turn to an iterate function, Returns the set for chaining.
func FpEachSet[T comparable](a BuiltinSet[T], iterate func(a BuiltinSet[T], node T) T) BuiltinSet[T] {
	b := NewSet[T](a.Len())
	a.ForEach(func(v T) {
		v2 := iterate(a, v)
		b.Insert(v2)
	})
	a.Clear()
	b.ForEach(func(k T) {
		a.Insert(k)
	})
	return a
}
