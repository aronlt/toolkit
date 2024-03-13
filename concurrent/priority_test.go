package concurrent

import (
	"testing"

	"github.com/aronlt/toolkit/ttypes"
	"github.com/stretchr/testify/assert"
)

func TestPriority(t *testing.T) {
	priority := NewPriorityChan[string](2, 10)
	highHandler := func(e string) error {
		t.Logf("in high handler receive:%s", e)
		return nil
	}
	lowHandler := func(e string) error {
		t.Logf("in low handler receive:%s", e)
		return nil
	}
	go func() {
		priority.Put("high", ttypes.HighPriorityType)
		priority.Put("low", ttypes.LowPriorityType)
		priority.Put("low", ttypes.LowPriorityType)
		priority.Put("low", ttypes.LowPriorityType)
		priority.Put("low", ttypes.LowPriorityType)
	}()
	// 优先处理high priority，然后才处理low priority
	err := priority.HandleSignal(highHandler, lowHandler)
	err = priority.HandleSignal(highHandler, lowHandler)
	err = priority.HandleSignal(highHandler, lowHandler)
	err = priority.HandleSignal(highHandler, lowHandler)
	assert.Nil(t, err)
}
