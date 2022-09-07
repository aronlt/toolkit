package concurrent

import (
	"time"

	"github.com/aronlt/toolkit/types"
)

type PriorityChan[T any] struct {
	highPriority chan T
	lowPriority  chan T
}

func NewPriorityChan[T any](size int) *PriorityChan[T] {
	return &PriorityChan[T]{
		highPriority: make(chan T, size),
		lowPriority:  make(chan T, size),
	}
}

// HandleSignal 通过传入处理函数，处理队列信号
func (p *PriorityChan[T]) HandleSignal(highHandler types.PriorityHandler, lowHandler types.PriorityHandler) error {
	select {
	case <-p.highPriority:
		return highHandler()
	default:
	}
	select {
	case <-p.highPriority:
		return highHandler()
	case <-p.lowPriority:
		return lowHandler()
	}
}

// GetWithTimeout 获取数据，支持超时返回
func (p *PriorityChan[T]) GetWithTimeout(timeout time.Duration) (T, error) {
	select {
	case v := <-p.highPriority:
		return v, nil
	default:
	}
	if timeout <= 0 {
		timeout = 1 * time.Second
	}
	ticker := time.NewTicker(timeout)
	defer ticker.Stop()

	select {
	case v := <-p.highPriority:
		return v, nil
	case v := <-p.lowPriority:
		return v, nil
	case <-ticker.C:
		var empty T
		return empty, types.ErrorTimeout
	}
}

// TryGet 尝试获取数据，如果没准备好直接返回
func (p *PriorityChan[T]) TryGet() T {
	select {
	case v := <-p.highPriority:
		return v
	default:
	}
	select {
	case v := <-p.highPriority:
		return v
	case v := <-p.lowPriority:
		return v
	default:
		var empty T
		return empty
	}
}

// Get 获取数据，阻塞等待
func (p *PriorityChan[T]) Get() T {
	select {
	case v := <-p.highPriority:
		return v
	default:
	}
	select {
	case v := <-p.highPriority:
		return v
	case v := <-p.lowPriority:
		return v
	}
}
