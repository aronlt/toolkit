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
