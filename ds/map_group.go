package ds

type GroupMap[K comparable, V any] map[K][]V

func NewGroupMap[K comparable, V any](data []V, getKeyHandler func(V) K) GroupMap[K, V] {
	group := make(map[K][]V)

	for _, elem := range data {
		key := getKeyHandler(elem)
		if _, ok := group[key]; !ok {
			group[key] = make([]V, 0)
		}
		group[key] = append(group[key], elem)
	}
	return group
}
