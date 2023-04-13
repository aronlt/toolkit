package ds

// FpPartitionSlice Split slice into two slices: one whose elements all satisfy predicate and one whose elements all do not satisfy predicate
func FpPartitionSlice[T any](a []T, iterate func(a []T, i int) bool) ([]T, []T) {
	pa := make([]T, 0, len(a))
	pb := make([]T, 0, len(a))

	FpIterSlice(a, func(a []T, i int) {
		if iterate(a, i) {
			pa = append(pa, a[i])
		} else {
			pb = append(pb, a[i])
		}
	})
	return pa, pb
}

// FpPartitionMap Split map into two maps: one whose elements all satisfy predicate and one whose elements all do not satisfy predicate
func FpPartitionMap[K comparable, V any](a map[K]V, iterate func(a map[K]V, k K, v V) bool) (map[K]V, map[K]V) {
	pa := make(map[K]V, len(a))
	pb := make(map[K]V, len(a))
	FpIterMap(a, func(a map[K]V, k K, v V) {
		if iterate(a, k, v) {
			pa[k] = v
		} else {
			pb[k] = v
		}
	})
	return pa, pb
}

// FpPartitionList Split list into two lists: one whose elements all satisfy predicate and one whose elements all do not satisfy predicate
func FpPartitionList[T any](a DList[T], iterate func(a DList[T], node T) bool) (DList[T], DList[T]) {
	pa := DList[T]{}
	pb := DList[T]{}
	FpIterList(a, func(a DList[T], node T) {
		if iterate(a, node) {
			pa.PushBack(node)
		} else {
			pb.PushBack(node)
		}
	})

	return pa, pb
}

// FpPartitionSet Split set into two sets: one whose elements all satisfy predicate and one whose elements all do not satisfy predicate
func FpPartitionSet[T comparable](a BuiltinSet[T], iterate func(a BuiltinSet[T], node T) bool) (BuiltinSet[T], BuiltinSet[T]) {
	pa := NewSet[T](a.Len())
	pb := NewSet[T](a.Len())
	FpIterSet(a, func(a BuiltinSet[T], node T) {
		if iterate(a, node) {
			pa.Insert(node)
		} else {
			pb.Insert(node)
		}
	})

	return pa, pb
}
