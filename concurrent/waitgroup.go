package concurrent

import (
	"github.com/aronlt/toolkit/types"
	"golang.org/x/sync/errgroup"
)

// WaitGroup wait group封装
func WaitGroup(number int, handler types.WaitGroupHandler, closeHandler ...func()) error {
	var wg errgroup.Group
	var err error

	for i := 0; i < number; i++ {
		wg.Go(handler)
	}

	defer func(wg *errgroup.Group) {
		if len(closeHandler) > 0 {
			closeHandler[0]()
		}
		err = wg.Wait()
	}(&wg)
	return err
}
