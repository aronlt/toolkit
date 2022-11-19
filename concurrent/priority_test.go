package concurrent

import (
	"github.com/aronlt/toolkit/ttypes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriority(t *testing.T) {
	priority := NewPriorityChan[string](10)
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
