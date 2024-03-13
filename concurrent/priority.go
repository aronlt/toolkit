package concurrent

import (
	"time"

	"github.com/aronlt/toolkit/ttypes"
)

type PriorityChan[T any] struct {
	highPriority chan T
	lowPriority  chan T
}

func NewPriorityChan[T any](hsize int, lsize int) *PriorityChan[T] {
	return &PriorityChan[T]{
		highPriority: make(chan T, hsize),
		lowPriority:  make(chan T, lsize),
	}
}

func (p *PriorityChan[T]) Put(event T, t ttypes.PriorityType) {
	switch t {
	case ttypes.HighPriorityType:
		p.highPriority <- event
	case ttypes.LowPriorityType:
		p.lowPriority <- event
	default:
		return
	}
}

func (p *PriorityChan[T]) TryPut(event T, t ttypes.PriorityType) error {
	switch t {
	case ttypes.HighPriorityType:
		select {
		case p.highPriority <- event:
			return nil
		default:
			return ttypes.ErrorFullChan
		}
	case ttypes.LowPriorityType:
		select {
		case p.lowPriority <- event:
			return nil
		default:
			return ttypes.ErrorFullChan
		}
	default:
		return ttypes.ErrorInvalidParameter
	}
}

func (p *PriorityChan[T]) PutWithTimeout(event T, t ttypes.PriorityType, timeout time.Duration) error {
	switch t {
	case ttypes.HighPriorityType:
		select {
		case p.highPriority <- event:
			return nil
		default:
		}
	case ttypes.LowPriorityType:
		select {
		case p.lowPriority <- event:
			return nil
		default:
		}
	}
	if timeout <= 0 {
		timeout = 1 * time.Second
	}
	ticker := time.NewTicker(timeout)
	defer ticker.Stop()
	switch t {
	case ttypes.HighPriorityType:
		select {
		case p.highPriority <- event:
			return nil
		case <-ticker.C:
			return ttypes.ErrorTimeout
		}
	case ttypes.LowPriorityType:
		select {
		case p.lowPriority <- event:
			return nil
		case <-ticker.C:
			return ttypes.ErrorTimeout
		}
	default:
		return ttypes.ErrorInvalidParameter
	}
}

// HandleSignal 通过传入处理函数，处理队列信号
func (p *PriorityChan[T]) HandleSignal(highHandler ttypes.PriorityHandler[T], lowHandler ttypes.PriorityHandler[T]) error {
	select {
	case t := <-p.highPriority:
		return highHandler(t)
	default:
	}
	select {
	case t := <-p.highPriority:
		return highHandler(t)
	case t := <-p.lowPriority:
		return lowHandler(t)
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
		return empty, ttypes.ErrorTimeout
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
