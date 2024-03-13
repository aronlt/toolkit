package concurrent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuture(t *testing.T) {
	ch := make(chan struct{})
	future := Run(func() *int {
		var i = 10
		<-ch
		return &i
	})

	n := 10
	for i := 0; i < n; i++ {
		if i == n-1 {
			ch <- struct{}{}
		}
		v, ok := future.TryGet()
		if i != n-1 {
			assert.False(t, ok)
			assert.Nil(t, v)
		} else {
			assert.True(t, ok)
			assert.Equal(t, *v, 10)
		}
	}
}
