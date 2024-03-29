package ds

// SliceIterMapInPlace Slice每个元素进行map映射
func SliceIterMapInPlace[T any](a []T, iterate func(i int) T) []T {
	for i := 0; i < len(a); i++ {
		a[i] = iterate(i)
	}
	return a
}

// MapIterMapKVInPlace Map每个kv进行map映射
func MapIterMapKVInPlace[K comparable, V any](a map[K]V, iterate func(k K, v V) V) map[K]V {
	for k, v := range a {
		v2 := iterate(k, v)
		a[k] = v2
	}
	return a
}

// ListIterMapInPlace Iterates over list, yielding each value in turn to an iterate function, Returns the list for chaining.
func ListIterMapInPlace[T any](a DList[T], iterate func(a DList[T], node T) T) DList[T] {
	iterator := a.Iterate()
	for iterator.IsNotEnd() {
		p := iterator.Pointer()
		*p = iterate(a, iterator.Value())
		iterator.MoveToNext()
	}
	return a
}

// SetIterMapInPlace Iterates over set, yielding each value in turn to an iterate function, Returns the set for chaining.
func SetIterMapInPlace[T comparable](a BuiltinSet[T], iterate func(node T) T) BuiltinSet[T] {
	b := NewSet[T](a.Len())
	a.ForEach(func(v T) {
		v2 := iterate(v)
		b.Insert(v2)
	})
	a.Clear()
	b.ForEach(func(k T) {
		a.Insert(k)
	})
	return a
}
