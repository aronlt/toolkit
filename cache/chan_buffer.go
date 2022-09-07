package cache

import (
	"context"
	"time"

	"github.com/aronlt/toolkit/types"
)

// ChanBuffer 基于channel实现的缓冲区
type ChanBuffer[T any] struct {
	waitingChan chan struct{}
	closeChan   chan struct{}
	fetcher     types.FetchHandler[T]
	result      chan T
	timeout     time.Duration
}

func NewChanBuffer[T any](fetcher types.FetchHandler[T], resultLen int, timeout time.Duration) *ChanBuffer[T] {
	return &ChanBuffer[T]{
		waitingChan: make(chan struct{}),
		closeChan:   make(chan struct{}),
		fetcher:     fetcher,
		result:      make(chan T, resultLen),
		timeout:     timeout,
	}

}

func (c *ChanBuffer[T]) Close() {
	select {
	case <-c.closeChan:
	default:
		close(c.closeChan)
	}
}

func (c *ChanBuffer[T]) Get() (T, error) {
	c.notify()

	// fast path
	select {
	case n := <-c.result:
		return n, nil
	default:
	}

	ticker := time.NewTicker(c.timeout)
	defer ticker.Stop()

	select {
	case n := <-c.result:
		return n, nil
	case <-ticker.C:
		var empty T
		return empty, types.ErrorTimeout
	}
}

// 通知开始获取数据
func (c *ChanBuffer[T]) notify() {
	select {
	case c.waitingChan <- struct{}{}:
	default:
	}
}

func (c *ChanBuffer[T]) start() {
	go func() {
		ctx := context.Background()
		for {
			select {
			case <-c.waitingChan:
				c.fetch(ctx)
			case <-c.closeChan:
				return
			}
		}
	}()
}

func (c *ChanBuffer[T]) fetch(ctx context.Context) {
	c.fetcher(ctx, c.result)
}
