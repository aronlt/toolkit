package ds

// SliceIter 只是迭代元素
func SliceIter[T any](a []T, iterate func(a []T, i int)) {
	for i := 0; i < len(a); i++ {
		iterate(a, i)
	}
}

// MapIter 只是迭代元素
func MapIter[K comparable, V any](a map[K]V, iterate func(a map[K]V, k K, v V)) {
	for k, v := range a {
		iterate(a, k, v)
	}
}

// ListIter 迭代元素
func ListIter[T any](a DList[T], iterate func(a DList[T], node T)) {
	iterator := a.Iterate()
	for iterator.IsNotEnd() {
		iterate(a, iterator.Value())
		iterator.MoveToNext()
	}
}

// SetIter 迭代元素
func SetIter[T comparable](a BuiltinSet[T], iterate func(a BuiltinSet[T], node T)) {
	a.ForEach(func(v T) {
		iterate(a, v)
	})
}
