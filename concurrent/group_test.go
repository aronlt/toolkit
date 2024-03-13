package concurrent

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"

	"github.com/aronlt/toolkit/ttypes"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Int()
}

func TestErrorGroupSame(t *testing.T) {
	var counter int64
	handler := func() error {
		atomic.AddInt64(&counter, 1)
		return nil
	}
	err := ErrorGroupSame(10, handler)
	assert.Nil(t, err)
	assert.Equal(t, counter, int64(10))
	var rerr = fmt.Errorf("random error")
	err = ErrorGroupSame(10, func() error {
		atomic.AddInt64(&counter, 1)
		if rand.Int()%2 == 0 {
			return rerr
		}
		return nil
	})
	assert.True(t, errors.Is(err, rerr))
}

func TestErrorGroupDiff(t *testing.T) {
	var counter int64
	handlers := []ttypes.ErrorGroupHandler{
		func() error {
			atomic.AddInt64(&counter, 1)
			return nil
		},
		func() error {
			atomic.AddInt64(&counter, 2)
			return nil
		},
		func() error {
			atomic.AddInt64(&counter, 3)
			return nil
		},
		func() error {
			atomic.AddInt64(&counter, 4)
			return nil
		},
		func() error {
			atomic.AddInt64(&counter, 5)
			return nil
		},
		func() error {
			atomic.AddInt64(&counter, 6)
			return nil
		}, func() error {
			atomic.AddInt64(&counter, 7)
			return nil
		}, func() error {
			atomic.AddInt64(&counter, 8)
			return nil
		}}
	err := ErrorGroupDiff(handlers)
	assert.Nil(t, err)
	assert.Equal(t, counter, int64(36))
}

func TestErrorGroupLimit(t *testing.T) {
	var counter int64
	handlers := []ttypes.ErrorGroupHandler{
		func() error {
			atomic.AddInt64(&counter, 1)
			return nil
		},
		func() error {
			atomic.AddInt64(&counter, 2)
			return nil
		},
		func() error {
			atomic.AddInt64(&counter, 3)
			return nil
		},
		func() error {
			atomic.AddInt64(&counter, 4)
			return nil
		},
		func() error {
			atomic.AddInt64(&counter, 5)
			return nil
		},
		func() error {
			atomic.AddInt64(&counter, 6)
			return nil
		}, func() error {
			atomic.AddInt64(&counter, 7)
			return nil
		}, func() error {
			atomic.AddInt64(&counter, 8)
			return nil
		}}
	err := ErrorGroupLimit(2, handlers)
	assert.Nil(t, err)
	assert.Equal(t, counter, int64(36))
}

func TestWaitGroupLimit(t *testing.T) {
	var counter int64
	handler := func() {
		atomic.AddInt64(&counter, 1)
	}
	WaitGroupLimit(2, []ttypes.WaitGroupHandler{
		handler,
		handler,
		handler,
		handler,
		handler,
		handler,
		handler,
		handler,
		handler,
		handler,
	})
	assert.Equal(t, counter, int64(10))
}

func TestWaitGroupSame(t *testing.T) {
	var counter int64
	handler := func() {
		atomic.AddInt64(&counter, 1)
	}
	WaitGroupSame(10, handler)
	assert.Equal(t, counter, int64(10))
}

func TestWaitGroupDiff(t *testing.T) {
	var counter int64
	handlers := []ttypes.WaitGroupHandler{
		func() {
			atomic.AddInt64(&counter, 1)
		},
		func() {
			atomic.AddInt64(&counter, 2)
		},
		func() {
			atomic.AddInt64(&counter, 3)
		},
		func() {
			atomic.AddInt64(&counter, 4)
		},
		func() {
			atomic.AddInt64(&counter, 5)
		},
		func() {
			atomic.AddInt64(&counter, 6)
		},
		func() {
			atomic.AddInt64(&counter, 7)
		},
		func() {
			atomic.AddInt64(&counter, 8)
		}}
	WaitGroupDiff(handlers)
	assert.Equal(t, counter, int64(36))
}
