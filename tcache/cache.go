package tcache

import (
	"time"

	"github.com/aronlt/toolkit/ttypes"
)

type IBufferPool interface {
	Get(n int) []byte
	Put(b []byte)
}

func NewBufferPool(baseline int) IBufferPool {
	return NewBufferPool(baseline)
}

type ICacheBuffer[T any] interface {
	Close()
	Get() (T, error)
}

func NewCacheBuffer[T any](fetcher ttypes.FetchHandler[T]) ICacheBuffer[T] {
	return NewChanBuffer(fetcher, 16, 3*time.Second)
}
