package ds

// SliceIterAllOk Returns true if all of the values pass the predicate truth test
func SliceIterAllOk[T any](a []T, iterate func(a []T, i int) bool) bool {
	for i := 0; i < len(a); i++ {
		if !iterate(a, i) {
			return false
		}
	}
	return true
}

// MapIterAllOk Returns true if all of the values pass the predicate truth test
func MapIterAllOk[K comparable, V any](a map[K]V, iterate func(a map[K]V, k K, v V) bool) bool {
	for k, v := range a {
		if !iterate(a, k, v) {
			return false
		}
	}
	return true
}

// ListIterAllOk Returns true if all of the values pass the predicate truth test
func ListIterAllOk[T any](a DList[T], iterate func(a DList[T], node T) bool) bool {
	iterator := a.Iterate()
	for iterator.IsNotEnd() {
		if !iterate(a, iterator.Value()) {
			return false
		}
		iterator.MoveToNext()
	}
	return true
}

// SetIterAllOk Returns true if all of the values pass the predicate truth test
func SetIterAllOk[T comparable](a BuiltinSet[T], iterate func(a BuiltinSet[T], node T) bool) bool {
	result := true
	a.ForEachIf(func(k T) bool {
		if !iterate(a, k) {
			result = false
			return false
		}
		return true
	})
	return result
}
