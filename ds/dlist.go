package ds

import "github.com/aronlt/toolkit/ttypes"

// copy from https://github.com/chen3feng/stl4go/blob/master/dlist.go

// DList is a doubly linked list.
type DList[T any] struct {
	head   *dListNode[T]
	length int
}

type dListNode[T any] struct {
	prev, next *dListNode[T]
	value      T
}

func NewDList[T any]() *DList[T] {
	return &DList[T]{}
}

// DListFromUnpack make a new DList from a serial of values.
func DListFromUnpack[T any](vs ...T) *DList[T] {
	l := &DList[T]{}
	for _, v := range vs {
		l.PushBack(v)
	}
	return l
}

// Clear cleanup the list.
func (l *DList[T]) Clear() {
	if l.head != nil {
		l.head.prev = l.head
		l.head.next = l.head
	}
	l.length = 0
}

// Len return the length of the list.
func (l *DList[T]) Len() int {
	return l.length
}

// IsEmpty return whether the list is empty.
func (l *DList[T]) IsEmpty() bool {
	return l.length == 0
}

type dlistIterator[T any] struct {
	dl   *DList[T]
	node *dListNode[T]
}

func (l *DList[T]) Iterate() MutableIterator[T] {
	node := l.head
	if node != nil {
		node = node.next
	}
	return &dlistIterator[T]{l, node}
}

// Iterator is the interface for container's iterator.
type Iterator[T any] interface {
	IsNotEnd() bool // Whether it is point to the end of the range.
	MoveToNext()    // Let it point to the next element.
	Value() T       // Return the value of current element.
}

// MutableIterator is the interface for container's mutable iterator.
type MutableIterator[T any] interface {
	Iterator[T]
	Pointer() *T // Return the pointer to the value of current element.
}

func (it *dlistIterator[T]) IsNotEnd() bool {
	return it.node != it.dl.head
}

func (it *dlistIterator[T]) MoveToNext() {
	it.node = it.node.next
}

func (it *dlistIterator[T]) Value() T {
	return it.node.value
}

func (it *dlistIterator[T]) Pointer() *T {
	return &it.node.value
}

// Front returns the first element in the container.
func (l *DList[T]) Front() T {
	if l.IsEmpty() {
		panic("!IsEmpty")
	}
	return l.head.next.value
}

// Back returns the last element in the container.
func (l *DList[T]) Back() T {
	if l.IsEmpty() {
		panic("!IsEmpty")
	}
	return l.head.prev.value
}

// PushFront pushes an element at the front of the list.
func (l *DList[T]) PushFront(val T) {
	l.ensureHead()
	n := dListNode[T]{l.head, l.head.next, val}
	l.head.next.prev = &n
	l.head.next = &n
	l.length++
}

// PushBack pushes an element at the back of the list.
func (l *DList[T]) PushBack(val T) {
	l.ensureHead()
	n := dListNode[T]{l.head.prev, l.head, val}
	l.head.prev.next = &n
	l.head.prev = &n
	l.length++
}

// PopFront popups an element from the front of the list.
func (l *DList[T]) PopFront() T {
	r, ok := l.TryPopFront()
	if !ok {
		panic("DList.PopFront: empty list")
	}
	return r
}

// PopBack popups an element from the back of the list.
func (l *DList[T]) PopBack() T {
	r, ok := l.TryPopBack()
	if !ok {
		panic("DList.PopBack: empty list")
	}
	return r
}

func (l *DList[T]) Values() []T {
	v := make([]T, 0, l.Len())
	l.ForEach(func(val T) {
		v = append(v, val)
	})
	return v
}

// TryPopFront tries to pop up an element from the front of the list.
func (l *DList[T]) TryPopFront() (T, bool) {
	var val T
	if l.IsEmpty() {
		return val, false
	}
	node := l.head.next
	val = node.value
	l.head.next = node.next
	l.head.next.prev = l.head
	node.prev = nil
	node.next = nil
	l.length--
	return val, true
}

// TryPopBack tries to pop up an element from the back of the list.
func (l *DList[T]) TryPopBack() (T, bool) {
	var val T
	if l.IsEmpty() {
		return val, false
	}
	node := l.head.prev
	val = node.value
	l.head.prev = l.head.prev.prev
	l.head.prev.next = l.head
	node.prev = nil
	node.next = nil
	l.length--
	return val, true
}

// ForEach iterate the list, apply each element to the cb callback function.
func (l *DList[T]) ForEach(cb func(val T)) {
	if l.IsEmpty() {
		return
	}
	for n := l.head.next; n != l.head; n = n.next {
		cb(n.value)
	}
}

// ForEachIf iterate the list, apply each element to the cb callback function,
// stop if cb returns false.
func (l *DList[T]) ForEachIf(cb func(val T) bool) {
	if l.IsEmpty() {
		return
	}
	for n := l.head.next; n != l.head; n = n.next {
		if !cb(n.value) {
			break
		}
	}
}

// ForEachMutable iterate the list, apply pointer of each element to the cb callback function.
func (l *DList[T]) ForEachMutable(cb func(val *T)) {
	if l.IsEmpty() {
		return
	}
	for n := l.head.next; n != l.head; n = n.next {
		cb(&n.value)
	}
}

// ForEachMutableIf iterate the list, apply pointer of each element to the cb callback function,
// stop if cb returns false.
func (l *DList[T]) ForEachMutableIf(cb func(val *T) bool) {
	if l.IsEmpty() {
		return
	}
	for n := l.head.next; n != l.head; n = n.next {
		if !cb(&n.value) {
			break
		}
	}
}

// ensureHead ensure head is valid.
func (l *DList[T]) ensureHead() {
	if l.head == nil {
		l.head = &dListNode[T]{}
		l.head.prev = l.head
		l.head.next = l.head
	}
}

func (l *DList[T]) RemoveValue(v T, fn ttypes.CompareFn[T]) bool {
	if l.IsEmpty() {
		return false
	}
	for n := l.head.next; n != l.head; n = n.next {
		if fn(v, n.value) == 0 {
			n.prev.next = n.next
			n.next.prev = n.prev
			n.prev = nil
			n.next = nil
			l.length--
			return true
		}
	}
	return false
}

func (l *DList[T]) InsertLessBound(v T, fn ttypes.LessEqFn[T]) {
	if l.IsEmpty() {
		l.PushBack(v)
		return
	}
	for node := l.head.next; node != l.head; node = node.next {
		if fn(v, node.value) {
			n := dListNode[T]{node.prev, node, v}
			node.prev.next = &n
			node.prev = &n
			l.length++
			return
		}
	}
	if fn(l.head.prev.value, v) {
		l.PushBack(v)
	} else if fn(v, l.head.next.value) {
		l.PushFront(v)
	} else {
		panic("should not be here")
	}
}
