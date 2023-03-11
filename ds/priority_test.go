package ds

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPriority(t *testing.T) {
	pq := NewWaitForPriorityQueue[int]()
	pq.Push(&WaitFor[int]{
		data:    1,
		readyAt: time.Now().Add(3 * time.Second),
		index:   0,
	}, false)
	pq.Push(&WaitFor[int]{
		data:    2,
		readyAt: time.Now().Add(5 * time.Second),
		index:   0,
	}, false)
	pq.Push(&WaitFor[int]{
		data:    3,
		readyAt: time.Now().Add(1 * time.Second),
		index:   0,
	}, false)

	pk := pq.Peek()
	assert.Equal(t, pk.data, 3)
	pq.Push(&WaitFor[int]{
		data:    3,
		readyAt: time.Now().Add(10 * time.Second),
		index:   0,
	}, false)
	pk = pq.Peek()
	assert.Equal(t, pk.data, 1)
	pk1 := pq.Pop()
	pk2 := pq.Pop()
	pk3 := pq.Pop()
	assert.Equal(t, pk1.data, 1)
	assert.Equal(t, pk2.data, 2)
	assert.Equal(t, pk3.data, 3)
}
