package concurrent

import (
	"github.com/aronlt/toolkit/ttypes"
	"golang.org/x/sync/errgroup"
)

// WaitGroup wait group封装
func WaitGroup(number int, handler ttypes.WaitGroupHandler, closeHandler ...ttypes.ErrHandler) error {
	var wg errgroup.Group
	var err error
	if number <= 0 {
		return ttypes.ErrorInvalidParameter
	}

	for i := 0; i < number; i++ {
		wg.Go(handler)
	}

	defer func(wg *errgroup.Group) {
		err = wg.Wait()
		if len(closeHandler) > 0 {
			closeHandler[0](err)
		}
	}(&wg)

	return err
}
