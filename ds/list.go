// Package ds Dlist copy from https://github.com/chen3feng/stl4go/blob/master/dlist.go
package ds

// List is a doubly linked list.
type List[T any] struct {
	head   *listNode[T]
	length int
}

type listNode[T any] struct {
	prev, next *listNode[T]
	value      T
}

func NewList[T any]() List[T] {
	return List[T]{}
}

// ListOf make a new List from a serial of values.
func ListOf[T any](vs ...T) List[T] {
	l := List[T]{}
	for _, v := range vs {
		l.PushBack(v)
	}
	return l
}

// Clear cleanup the list.
func (l *List[T]) Clear() {
	if l.head != nil {
		l.head.prev = l.head
		l.head.next = l.head
	}
	l.length = 0
}

// Len return the length of the list.
func (l *List[T]) Len() int {
	return l.length
}

// IsEmpty return whether the list is empty.
func (l *List[T]) IsEmpty() bool {
	return l.length == 0
}

type dlistIterator[T any] struct {
	dl   *List[T]
	node *listNode[T]
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

// Iterate returns an iterator to the first element in the list.
func (l *List[T]) Iterate() MutableIterator[T] {
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

// Front returns the first element in the container.
func (l *List[T]) Front() T {
	if l.IsEmpty() {
		panic("!IsEmpty")
	}
	return l.head.next.value
}

// Back returns the last element in the container.
func (l *List[T]) Back() T {
	if l.IsEmpty() {
		panic("!IsEmpty")
	}
	return l.head.prev.value
}

// PushFront pushes an element at the front of the list.
func (l *List[T]) PushFront(val T) {
	l.ensureHead()
	n := listNode[T]{l.head, l.head.next, val}
	l.head.next.prev = &n
	l.head.next = &n
	l.length++
}

// PushBack pushes an element at the back of the list.
func (l *List[T]) PushBack(val T) {
	l.ensureHead()
	n := listNode[T]{l.head.prev, l.head, val}
	l.head.prev.next = &n
	l.head.prev = &n
	l.length++
}

// PopFront popups a element from the front of the list.
func (l *List[T]) PopFront() T {
	r, ok := l.TryPopFront()
	if !ok {
		panic("List.PopFront: empty list")
	}
	return r
}

// PopBack popups a element from the back of the list.
func (l *List[T]) PopBack() T {
	r, ok := l.TryPopBack()
	if !ok {
		panic("List.PopBack: empty list")
	}
	return r
}

// TryPopFront tries to popup a element from the front of the list.
func (l *List[T]) TryPopFront() (T, bool) {
	var val T
	if l.length == 0 {
		return val, false
	}
	node := l.head.next
	val = node.value
	l.head.next = node.next
	l.head.prev = l.head
	node.prev = nil
	node.next = nil
	l.length--
	return val, true
}

// TryPopBack tries to popup a element from the back of the list.
func (l *List[T]) TryPopBack() (T, bool) {
	var val T
	if l.length == 0 {
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
func (l *List[T]) ForEach(cb func(val T)) {
	if l.head == nil {
		return
	}
	for n := l.head.next; n != l.head; n = n.next {
		cb(n.value)
	}
}

// ForEachIf iterate the list, apply each element to the cb callback function,
// stop if cb returns false.
func (l *List[T]) ForEachIf(cb func(val T) bool) {
	if l.head == nil {
		return
	}
	for n := l.head.next; n != l.head; n = n.next {
		if !cb(n.value) {
			break
		}
	}
}

// ForEachMutable iterate the list, apply pointer of each element to the cb callback function.
func (l *List[T]) ForEachMutable(cb func(val *T)) {
	if l.head == nil {
		return
	}
	for n := l.head.next; n != l.head; n = n.next {
		cb(&n.value)
	}
}

// ForEachMutableIf iterate the list, apply pointer of each element to the cb callback function,
// stop if cb returns false.
func (l *List[T]) ForEachMutableIf(cb func(val *T) bool) {
	if l.head == nil {
		return
	}
	for n := l.head.next; n != l.head; n = n.next {
		if !cb(&n.value) {
			break
		}
	}
}

// ensureHead ensure head is valid.
func (l *List[T]) ensureHead() {
	if l.head == nil {
		l.head = &listNode[T]{}
		l.head.prev = l.head
		l.head.next = l.head
	}
}
