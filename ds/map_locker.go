package ds

import "sync"

type MapLocker[K comparable, V any] struct {
	data   map[K]V
	locker sync.Mutex
}

func NewMapLocker[K comparable, V any]() *MapLocker[K, V] {
	return &MapLocker[K, V]{
		data: make(map[K]V, 0),
	}
}

func (m *MapLocker[K, V]) Get(key K) (V, bool) {
	m.locker.Lock()
	defer m.locker.Unlock()
	v, ok := m.data[key]
	return v, ok
}

func (m *MapLocker[K, V]) Contain(key K) bool {
	m.locker.Lock()
	defer m.locker.Unlock()
	_, ok := m.data[key]
	return ok
}

func (m *MapLocker[K, V]) Set(key K, value V) {
	m.locker.Lock()
	defer m.locker.Unlock()
	m.data[key] = value
}

func (m *MapLocker[K, V]) Foreach(handler func(key K, value V)) {
	m.locker.Lock()
	defer m.locker.Unlock()
	for k, v := range m.data {
		handler(k, v)
	}
}
