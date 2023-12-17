package ds

// SliceIter 只是迭代元素
func SliceIter[T any](a []T, iterate func(i int)) {
	for i := 0; i < len(a); i++ {
		iterate(i)
	}
}

// MapIter 只是迭代元素
func MapIter[K comparable, V any](a map[K]V, iterate func(k K, v V)) {
	for k, v := range a {
		iterate(k, v)
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
func SetIter[T comparable](a BuiltinSet[T], iterate func(node T)) {
	a.ForEach(func(v T) {
		iterate(v)
	})
}
