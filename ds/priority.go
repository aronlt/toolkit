// copy from k8s delay_queue, src:vendor/k8s.io/client-go/util/workqueue/delaying_queue.go

package ds

import (
	"container/heap"
	"time"
)

// WaitFor holds the data to add and the time it should be added
type WaitFor[T comparable] struct {
	data    T
	readyAt time.Time
	// index in the priority queue (heap)
	index int
}

// WaitForPriorityQueue implements a priority queue for WaitFor items.
//
// WaitForPriorityQueue implements heap.Interface. The item occurring next in
// time (i.e., the item with the smallest readyAt) is at the root (index 0).
// Peek returns this minimum item at index 0. Pop returns the minimum item after
// it has been removed from the queue and placed at index Len()-1 by
// container/heap. Push adds an item at index Len(), and container/heap
// percolates it into the correct location.
type WaitForPriorityQueue[T comparable] struct {
	queue        priorityQueue[T]
	knownEntries map[T]*WaitFor[T]
}

func newPriorityQueue[T comparable]() priorityQueue[T] {
	return make([]*WaitFor[T], 0)
}

type priorityQueue[T comparable] []*WaitFor[T]

func (pq *priorityQueue[T]) Len() int {
	return len(*pq)
}
func (pq *priorityQueue[T]) Less(i, j int) bool {
	return (*pq)[i].readyAt.Before((*pq)[j].readyAt)
}
func (pq *priorityQueue[T]) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

// Push adds an item to the queue. Push should not be called directly; instead,
// use `heap.Push`.
func (pq *priorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*WaitFor[T])
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes an item from the queue. Pop should not be called directly;
// instead, use `heap.Pop`.
func (pq *priorityQueue[T]) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	item.index = -1
	*pq = (*pq)[0:(n - 1)]
	return item
}

// Push adds the entry to the priority queue, or updates the readyAt if it already exists in the queue
func (pq *WaitForPriorityQueue[T]) Push(entry *WaitFor[T], ignore bool) {
	// if the entry already exists, update the time only if it would cause the item to be queued sooner
	existing, exists := pq.knownEntries[entry.data]
	if exists {
		if !ignore {
			existing.readyAt = entry.readyAt
			heap.Fix(&pq.queue, existing.index)
		}
		return
	}

	heap.Push(&pq.queue, entry)
	pq.knownEntries[entry.data] = entry
}

func (pq *WaitForPriorityQueue[T]) Len() int {
	return len(pq.queue)
}

func (pq *WaitForPriorityQueue[T]) Pop() *WaitFor[T] {
	k := heap.Pop(&pq.queue)
	w := k.(*WaitFor[T])
	delete(pq.knownEntries, w.data)
	return w
}

// Peek returns the item at the beginning of the queue, without removing the
// item or otherwise mutating the queue. It is safe to call directly.
func (pq *WaitForPriorityQueue[T]) Peek() *WaitFor[T] {
	return pq.queue[0]
}

func NewWaitForPriorityQueue[T comparable]() *WaitForPriorityQueue[T] {
	return &WaitForPriorityQueue[T]{
		queue:        newPriorityQueue[T](),
		knownEntries: make(map[T]*WaitFor[T], 0),
	}
}
