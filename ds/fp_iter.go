package ds

// SliceIter 只是迭代元素
func SliceIter[T any](a []T, iterate func(a []T, i int)) {
	for i := 0; i < len(a); i++ {
		iterate(a, i)
	}
}

// SliceIterV2 只是迭代元素
func SliceIterV2[T any](a []T, iterate func(i int)) {
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

// DListIter 迭代单向链表元素
func DListIter[T any](a DList[T], iterate func(a DList[T], node T)) {
	iterator := a.Iterate()
	for iterator.IsNotEnd() {
		iterate(a, iterator.Value())
		iterator.MoveToNext()
	}
}

// SListIter 迭代单向链表元素
func SListIter[T any](a SList[T], iterate func(a SList[T], node T)) {
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
