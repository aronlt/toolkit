package concurrent

type Future[T any] struct {
	Chan chan *T
}

func Run[T any](fn func() *T) *Future[T] {
	future := &Future[T]{
		Chan: make(chan *T),
	}
	go func() {
		future.Chan <- fn()
	}()
	return future
}

func (f *Future[T]) TryGet() (*T, bool) {
	select {
	case v := <-f.Chan:
		return v, true
	default:
		return nil, false
	}
}
