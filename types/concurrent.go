package types

type IFanOut[T any] interface {
	Produce() chan T
	Compute(chan T) chan T
	Merge(...chan T) chan T
}

type PriorityHandler func() error
type ErrHandler func(err any)
type WaitGroupHandler func() error
