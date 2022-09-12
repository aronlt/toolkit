package concurrent

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type FanOutTest struct {
}

func (f *FanOutTest) Produce() chan int {
	cout := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			cout <- i
		}
		close(cout)
		return
	}()
	return cout
}

func (f *FanOutTest) Compute(cin chan int) chan int {
	cout := make(chan int, 10)
	go func() {
		for v := range cin {
			cout <- v + 10
		}
		close(cout)
	}()
	return cout
}

func (f *FanOutTest) Merge(cins ...chan int) chan int {
	cout := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(len(cins))
	for i := 0; i < len(cins); i++ {
		go func(i int) {
			for v := range cins[i] {
				cout <- v
			}
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(cout)
	}()
	return cout
}

func TestFanOut(t *testing.T) {
	data := &FanOutTest{}
	result := FanOut[int](data, 3)
	sum := 0
	for _, v := range result {
		sum += v
	}
	assert.Equal(t, sum, 145)
}
