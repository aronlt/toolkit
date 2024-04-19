package ttypes

// copy from golang.org/x/exp

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}
type Integer interface {
	Signed | Unsigned
}
type Float interface {
	~float32 | ~float64
}
type Ordered interface {
	Integer | Float | ~string
}
type IComparator interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// LessEqFn is a function that returns whether 'a' less than or equal to 'b'.
type LessEqFn[T any] func(a, b T) bool

// LessFn is a function that returns whether 'a' is less than 'b'.
type LessFn[T any] func(a, b T) bool

// GreaterFn is a function that returns whether 'a' greater than 'b'.
type GreaterFn[T any] func(a, b T) bool

// GreaterEqFn is a function that returns whether 'a' greater than or equal to 'b'.
type GreaterEqFn[T any] func(a, b T) bool

// CompareFn is a function that return a compare to b
type CompareFn[T any] func(a, b T) int

// Less wraps the '<' operator for ordered types.
func Less[T Ordered](a, b T) bool {
	return a < b
}

// LessEq wraps the '<=' operator for ordered types.
func LessEq[T Ordered](a, b T) bool {
	return a <= b
}

// Greater wraps the '>' operator for ordered types.
func Greater[T Ordered](a, b T) bool {
	return a > b
}

// GreaterEq wraps the '>=' operator for ordered types.
func GreaterEq[T Ordered](a, b T) bool {
	return a >= b
}

// OrderedCompare provide default CompareFn for ordered types.
func OrderedCompare[T Ordered](a, b T) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}
