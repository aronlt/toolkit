package concurrent

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	var counter int64
	handler := func() error {
		atomic.AddInt64(&counter, 1)
		return nil
	}
	number := 10
	errHandler := func(err any) {
		fmt.Printf("error handler:%v", err)
	}
	err := WaitGroup(number, handler, errHandler)
	assert.Nil(t, err)
	assert.Equal(t, counter, int64(10))
}
