package ds

// FpIterSlice Iterate slice by iterate func, do not modify slice
func FpIterSlice[T any](a []T, iterate func(a []T, i int)) {
	for i := 0; i < len(a); i++ {
		iterate(a, i)
	}
}

// FpIterMap Iterate map by iterate func, do not modify map
func FpIterMap[K comparable, V any](a map[K]V, iterate func(a map[K]V, k K, v V)) {
	for k, v := range a {
		iterate(a, k, v)
	}
}

// FpIterList Iterate list by iterate func, do not modify list
func FpIterList[T any](a DList[T], iterate func(a DList[T], node T)) {
	iterator := a.Iterate()
	for iterator.IsNotEnd() {
		iterate(a, iterator.Value())
		iterator.MoveToNext()
	}
}

// FpIterSet Iterate set by iterate func, do not modify set
func FpIterSet[T comparable](a BuiltinSet[T], iterate func(a BuiltinSet[T], node T)) {
	a.ForEach(func(v T) {
		iterate(a, v)
	})
}
