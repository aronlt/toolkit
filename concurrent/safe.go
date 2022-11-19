package concurrent

import "github.com/aronlt/toolkit/ttypes"

// RunSafe 捕获panic的协程
func RunSafe(handler func(), errHandler ...ttypes.ErrHandler) {
	defer func() {
		if r := recover(); r != nil {
			if len(errHandler) > 0 {
				errHandler[0](r)
			}
		}
	}()
	handler()
}
