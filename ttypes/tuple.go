package ttypes

// OrderTuple 有序元组，用于map的平铺，方便排序
type OrderTuple[K Ordered, V any] struct {
	Key   K
	Value V
}
