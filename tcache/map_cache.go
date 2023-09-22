package tcache

// MemCache 轻量级即插即用的map缓存，只能用于局部cache
type MemCache[K comparable, V any] struct {
	data map[K]V
}

func NewMemCache[K comparable, V any]() *MemCache[K, V] {
	return &MemCache[K, V]{
		data: make(map[K]V, 1024),
	}
}

func (m *MemCache[K, V]) Get(key K) (V, bool) {
	v, ok := m.data[key]
	return v, ok
}

func (m *MemCache[K, V]) Load(key K, value V) {
	m.data[key] = value
}
