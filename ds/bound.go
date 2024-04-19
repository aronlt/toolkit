package ds

import "github.com/aronlt/toolkit/ttypes"

// copy from https://github.com/chen3feng/stl4go/blob/master/binary_search.go

// LowerBound returns an index to the first element in the ascending ordered slice a that
// does not satisfy element < value (i.e. greater or equal to),
// or len(a) if no such element is found.
//
// Complexity: O(log(len(a))).
func LowerBound[T ttypes.Ordered](a []T, value T) int {
	loc := 0
	count := len(a)
	for count > 0 {
		i := loc
		step := count / 2
		i += step
		if a[i] < value {
			i++
			loc = i
			count -= step + 1
		} else {
			count = step
		}
	}
	return loc
}

// LowerBoundFunc returns an index to the first element in the ordered slice a that
// does not satisfy less(element, value)), or len(a) if no such element is found.
//
// The elements in the slice a should sorted according with compare func less.
//
// Complexity: O(log(len(a))).
func LowerBoundFunc[T any](a []T, value T, less ttypes.LessFn[T]) int {
	loc := 0
	count := len(a)
	for count > 0 {
		i := loc
		step := count / 2
		i += step
		if less(a[i], value) {
			i++
			loc = i
			count -= step + 1
		} else {
			count = step
		}
	}
	return loc
}

// UpperBound returns an index to the first element in the ascending ordered slice a such that
// value < element (i.e. strictly greater), or len(a) if no such element is found.
//
// Complexity: O(log(len(a))).
func UpperBound[T ttypes.Ordered](a []T, value T) int {
	loc := 0
	count := len(a)
	for count > 0 {
		i := loc
		step := count / 2
		i += step
		if !(value < a[i]) {
			i++
			loc = i
			count -= step + 1
		} else {
			count = step
		}
	}
	return loc
}

// UpperBoundFunc returns an index to the first element in the ordered slice a such that
// less(value, element)) is true (i.e. strictly greater), or len(a) if no such element is found.
//
// The elements in the slice a should sorted according with compare func less.
//
// Complexity: O(log(len(a))).
func UpperBoundFunc[T any](a []T, value T, less ttypes.LessFn[T]) int {
	loc := 0
	count := len(a)
	for count > 0 {
		i := loc
		step := count / 2
		i += step
		if !less(value, a[i]) {
			i++
			loc = i
			count -= step + 1
		} else {
			count = step
		}
	}
	return loc
}
