package concurrent

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLimit(t *testing.T) {
	limiter := NewLimit(3)
	var v int64
	wg := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			limiter.Get()
			defer limiter.Put()
			defer wg.Done()
			atomic.AddInt64(&v, 1)
		}()
	}
	wg.Wait()
	assert.Equal(t, v, int64(4))
}

func TestLimitLoop(t *testing.T) {
	limiter := NewLimit(3)
	var v int64
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 4; i++ {
		go func() {
			limiter.Get()
			atomic.AddInt64(&v, 1)
			wg.Done()
			time.Sleep(3 * time.Second)
		}()
	}
	wg.Wait()
	assert.Equal(t, v, int64(3))
}
