package toolkit

type CounterMap[K comparable] map[K]int

func NewCounterMap[K comparable](data []K) CounterMap[K] {
	counter := make(map[K]int, len(data))
	for _, elem := range data {
		if _, ok := counter[elem]; ok {
			counter[elem] += 1
		} else {
			counter[elem] = 1
		}
	}
	return counter
}

func (c CounterMap[K]) Equal(other CounterMap[K]) bool {
	for v, counter1 := range c {
		if counter2, ok := other[v]; !ok || counter1 != counter2 {
			return false
		}
	}

	for v, counter1 := range other {
		if counter2, ok := c[v]; !ok || counter1 != counter2 {
			return false
		}
	}
	return true
}
