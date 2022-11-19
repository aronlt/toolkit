package ttypes

import "time"

type PriorityType int

const HighPriorityType PriorityType = 0
const LowPriorityType PriorityType = 1

type ILimit interface {
	Put()
	Get()
}

type IPriorityChan[T any] interface {
	Put(event T, t PriorityType)
	TryPut(event T, t PriorityType) error
	PutWithTimeout(event T, t PriorityType, timeout time.Duration) error
	Get() T
	TryGet() T
	GetWithTimeout(timeout time.Duration) (T, error)
	HandleSignal(highHandler PriorityHandler[T], lowHandler PriorityHandler[T]) error
}

type IFanOut[T any] interface {
	Produce() chan T
	Compute(chan T) chan T
	Merge(...chan T) chan T
}

type PriorityHandler[T any] func(event T) error
type ErrHandler func(err any)
type WaitGroupHandler func() error
