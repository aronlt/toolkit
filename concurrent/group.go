package concurrent

import (
	"sync"

	"github.com/aronlt/toolkit/ttypes"
	"golang.org/x/sync/errgroup"
)

func ErrorGroupSame(number int, handler ttypes.ErrorGroupHandler, closeHandler ...ttypes.ErrHandler) error {
	if number <= 0 {
		return ttypes.ErrorInvalidParameter
	}

	var eg errgroup.Group
	for i := 0; i < number; i++ {
		eg.Go(handler)
	}

	err := eg.Wait()
	if err != nil && len(closeHandler) > 0 {
		closeHandler[0](err)
	}
	return err
}

func ErrorGroupDiff(handlers []ttypes.ErrorGroupHandler, closeHandler ...ttypes.ErrHandler) error {
	if len(handlers) <= 0 {
		return ttypes.ErrorInvalidParameter
	}

	var eg errgroup.Group
	for i := 0; i < len(handlers); i++ {
		eg.Go(handlers[i])
	}

	err := eg.Wait()
	if err != nil && len(closeHandler) > 0 {
		closeHandler[0](err)
	}
	return err
}

func ErrorGroupLimit(n int, handlers []ttypes.ErrorGroupHandler, closeHandler ...ttypes.ErrHandler) error {
	var eg errgroup.Group
	eg.SetLimit(n)
	for i := 0; i < len(handlers); i++ {
		eg.Go(handlers[i])
	}

	err := eg.Wait()
	if err != nil && len(closeHandler) > 0 {
		closeHandler[0](err)
	}
	return err
}

func WaitGroupLimit(n int, handlers []ttypes.WaitGroupHandler) {
	limit := NewLimit(n)
	var wg sync.WaitGroup
	for i := 0; i < len(handlers); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			limit.Get()
			defer limit.Put()
			handlers[i]()
			return
		}(i)
	}
	wg.Wait()
}

func WaitGroupSame(number int, handler ttypes.WaitGroupHandler) {
	if number <= 0 {
		return
	}

	var wg sync.WaitGroup
	for i := 0; i < number; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			handler()
		}()
	}

	wg.Wait()
}

func WaitGroupDiff(handlers []ttypes.WaitGroupHandler) {
	if len(handlers) <= 0 {
		return
	}

	var wg sync.WaitGroup
	for i := 0; i < len(handlers); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			handlers[i]()
		}(i)
	}

	wg.Wait()
}
