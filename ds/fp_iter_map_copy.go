package ds

// SliceIterMapCopy Produces a new slice by mapping each value in list through a transformation function (iterate).
func SliceIterMapCopy[T any](a []T, iterate func(i int) T) []T {
	b := make([]T, 0, len(a))
	for i := 0; i < len(a); i++ {
		v := iterate(i)
		b = append(b, v)
	}
	return b
}

// MapIterMapKVCopy Produces a new map by mapping each key, value in map through a transformation function (iterate).
func MapIterMapKVCopy[K comparable, V any](a map[K]V, iterate func(k K, v V) (K, V)) map[K]V {
	b := make(map[K]V, len(a))
	for k, v := range a {
		k2, v2 := iterate(k, v)
		b[k2] = v2
	}
	return b
}

// ListIterMapCopy Produces a new list by mapping each value in list through a transformation function (iterate).
func ListIterMapCopy[T any](a DList[T], iterate func(a DList[T], node T) T) DList[T] {
	b := DList[T]{}
	iterator := a.Iterate()
	for iterator.IsNotEnd() {
		b.PushBack(iterate(a, iterator.Value()))
		iterator.MoveToNext()
	}
	return b
}

// SetIterMapCopy Produces a new list by mapping each value in list through a transformation function (iterate).
func SetIterMapCopy[T comparable](a BuiltinSet[T], iterate func(node T) T) BuiltinSet[T] {
	b := NewSet[T]()
	a.ForEach(func(v T) {
		v2 := iterate(v)
		b.Insert(v2)
	})
	return b
}
