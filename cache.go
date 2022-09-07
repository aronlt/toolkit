package toolkit

import (
	"time"

	"github.com/aronlt/toolkit/cache"
	"github.com/aronlt/toolkit/types"
)

type IBufferPool interface {
	Get(n int) []byte
	Put(b []byte)
}

func NewBufferPool(baseline int) IBufferPool {
	return cache.NewBufferPool(baseline)
}

type ICacheBuffer[T any] interface {
	Close()
	Get() (T, error)
}

func NewCacheBuffer[T any](fetcher types.FetchHandler[T]) ICacheBuffer[T] {
	return cache.NewChanBuffer(fetcher, 16, 3*time.Second)
}
