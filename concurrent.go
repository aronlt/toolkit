package toolkit

import (
	"time"

	"github.com/aronlt/toolkit/concurrent"
	"github.com/aronlt/toolkit/types"
)

func FanOut[T any](data types.IFanOut[T], number int) []T {
	return concurrent.FanOut(data, number)
}

type ILimit interface {
	Put()
	Get()
}

func NewLimit(max int) ILimit {
	return concurrent.NewLimit(max)
}

type IPriorityChan[T any] interface {
	Get() T
	TryGet() T
	GetWithTimeout(timeout time.Duration) (T, error)
	HandleSignal(highHandler types.PriorityHandler, lowHandler types.PriorityHandler) error
}

func NewPriorityChan[T any](size int) IPriorityChan[T] {
	return concurrent.NewPriorityChan[T](size)
}

func RunSafe(handler func(), errHandler ...types.ErrHandler) {
	concurrent.RunSafe(handler, errHandler...)
}
func WaitGroup(number int, handler types.WaitGroupHandler, closeHandler ...func()) error {
	return concurrent.WaitGroup(number, handler, closeHandler...)
}
