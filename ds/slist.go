package ds

import "github.com/aronlt/toolkit/ttypes"

// copy from https://github.com/chen3feng/stl4go/blob/master/slist.go

// SList is a singly linked list.
type SList[T any] struct {
	head   *sListNode[T]
	tail   *sListNode[T] // To support Back and PushBack
	length int
}

type sListNode[T any] struct {
	next  *sListNode[T]
	value T
}

func NewSList[T any]() *SList[T] {
	return &SList[T]{}
}

// SListFromUnpack return a SList that contains values.
func SListFromUnpack[T any](values ...T) *SList[T] {
	l := &SList[T]{}
	for i := range values {
		l.PushBack(values[i])
	}
	return l
}

// IsEmpty checks if the container has no elements.
func (l *SList[T]) IsEmpty() bool {
	return l.length == 0
}

// Len returns the number of elements in the container.
func (l *SList[T]) Len() int {
	return l.length
}

// Clear erases all elements from the container. After this call, Len() returns zero.
func (l *SList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

// Front returns the first element in the list.
func (l *SList[T]) Front() T {
	if l.IsEmpty() {
		panic("!IsEmpty")
	}
	return l.head.value
}

// Back returns the last element in the list.
func (l *SList[T]) Back() T {
	if l.IsEmpty() {
		panic("!IsEmpty")
	}
	return l.tail.value
}

func (l *SList[T]) RemoveValue(v T, fn ttypes.CompareFn[T]) bool {
	var pre *sListNode[T]
	found := false
	for node := l.head; node != nil; node = node.next {
		if fn(v, node.value) == 0 {
			found = true
			break
		} else {
			pre = node
		}
	}
	if !found {
		return false
	}
	if pre == nil {
		if fn(l.head.value, v) == 0 {
			l.PopFront()
		} else if fn(l.tail.value, v) == 0 {
			l.PopTail()
		} else {
			panic("should not be here")
		}
		return true
	}
	node := pre.next
	if node == nil {
		panic("node should not be nil")
	}
	pre.next = node.next
	node.next = nil
	l.length--
	return true
}

func (l *SList[T]) InsertLessBound(v T, fn ttypes.LessEqFn[T]) {
	if l.IsEmpty() {
		l.PushBack(v)
		return
	}
	var pre *sListNode[T]
	for node := l.head; node != nil; node = node.next {
		if fn(v, node.value) {
			if node == l.head {
				n := &sListNode[T]{nil, v}
				n.next = l.head
				l.head = n
				l.length++
				return
			}
			break
		} else {
			pre = node
		}
	}
	if pre == nil {
		panic("should not be here")
	}

	n := &sListNode[T]{nil, v}
	n.next = pre.next
	pre.next = n
	if pre == l.tail {
		l.tail = n
	}
	l.length++
}

// PushFront pushed an element to the front of the list.
func (l *SList[T]) PushFront(v T) {
	node := sListNode[T]{l.head, v}
	l.head = &node
	if l.tail == nil {
		l.tail = &node
	}
	l.length++
}

// PushBack pushed an element to the tail of the list.
func (l *SList[T]) PushBack(v T) {
	node := sListNode[T]{nil, v}
	if l.tail != nil {
		l.tail.next = &node
	}
	l.tail = &node
	if l.head == nil {
		l.head = &node
	}
	l.length++
}

// PopFront popups an element from the front of the list.
// The list must be non-empty!
func (l *SList[T]) PopFront() T {
	if l.IsEmpty() {
		panic("!IsEmpty")
	}

	node := l.head
	l.head = node.next
	if l.head == nil {
		l.tail = nil
	}
	l.length--
	return node.value
}

// PopTail popups an element from the tail of the list.
// The list must be non-empty!
func (l *SList[T]) PopTail() T {
	if l.IsEmpty() {
		panic("!IsEmpty")
	}

	if l.head == l.tail {
		return l.PopFront()
	}

	node := l.head
	tail := l.tail
	for node.next != l.tail {
		node = node.next
	}
	node.next = l.tail.next
	l.tail = node
	l.length--
	return tail.value
}

// Reverse reverses the order of all elements in the container.
func (l *SList[T]) Reverse() {
	var head, tail *sListNode[T]
	for node := l.head; node != nil; {
		next := node.next
		node.next = head
		head = node
		if tail == nil {
			tail = node
		}
		node = next
	}
	l.head = head
	l.tail = tail
}

// Values copies all elements in the container to a slice and return it.
func (l *SList[T]) Values() []T {
	s := make([]T, 0, l.Len())
	for node := l.head; node != nil; node = node.next {
		s = append(s, node.value)
	}
	return s
}

// InsertAfter inserts an element after the iterator into the list,
// return an iterator to the inserted element.
// func (l *SList[T]) InsertAfter(it Iterator[T], value T) MutableIterator[T] {
// 	// cause internal compiler error: panic: runtime error: invalid memory address or nil pointer dereference
// 	itp := it.(*sListIterator[T])
// 	node := itp.node
// 	newNode := sListNode[T]{node.next, value}
// 	node.next = &newNode
// 	return &sListIterator[T]{&newNode}
// }

// ForEach iterate the list, apply each element to the cb callback function.
func (l *SList[T]) ForEach(cb func(T)) {
	for node := l.head; node != nil; node = node.next {
		cb(node.value)
	}
}

// ForEachIf iterate the container, apply each element to the cb callback function,
// stop if cb returns false.
func (l *SList[T]) ForEachIf(cb func(T) bool) {
	for node := l.head; node != nil; node = node.next {
		if !cb(node.value) {
			break
		}
	}
}

// ForEachMutable iterate the container, apply pointer of each element to the cb callback function.
func (l *SList[T]) ForEachMutable(cb func(*T)) {
	for node := l.head; node != nil; node = node.next {
		cb(&node.value)
	}
}

// ForEachMutableIf iterate the container, apply pointer of each element to the cb callback function,
// stop if cb returns false.
func (l *SList[T]) ForEachMutableIf(cb func(*T) bool) {
	for node := l.head; node != nil; node = node.next {
		if !cb(&node.value) {
			break
		}
	}
}

// Iterate returns an iterator to the whole container.
func (l *SList[T]) Iterate() MutableIterator[T] {
	it := sListIterator[T]{l.head}
	return &it
}

type sListIterator[T any] struct {
	node *sListNode[T]
}

func (it *sListIterator[T]) IsNotEnd() bool {
	return it.node != nil
}

func (it *sListIterator[T]) MoveToNext() {
	it.node = it.node.next
}

func (it *sListIterator[T]) Value() T {
	return it.node.value
}

func (it *sListIterator[T]) Pointer() *T {
	return &it.node.value
}
