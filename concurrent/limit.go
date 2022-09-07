package concurrent

// Limit 限制并发器，限制最大同时运行的数量
type Limit struct {
	ch chan struct{}
}

func NewLimit(max int) *Limit {
	if max <= 0 {
		max = 1
	}
	rlimit := &Limit{ch: make(chan struct{}, max)}
	for i := 0; i < max; i++ {
		rlimit.ch <- struct{}{}
	}
	return rlimit
}
func (r *Limit) Put() {
	r.ch <- struct{}{}
}
func (r *Limit) Get() {
	<-r.ch
}
