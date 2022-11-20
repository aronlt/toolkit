package ds

import "sync"

type MapRWLocker[K comparable, V any] struct {
	data   map[K]V
	locker sync.RWMutex
}

func NewMapRWLocker[K comparable, V any]() *MapRWLocker[K, V] {
	return &MapRWLocker[K, V]{
		data: make(map[K]V, 0),
	}
}

func (m *MapRWLocker[K, V]) Get(key K) (V, bool) {
	m.locker.RLock()
	defer m.locker.RUnlock()
	v, ok := m.data[key]
	return v, ok
}

func (m *MapRWLocker[K, V]) Contain(key K) bool {
	m.locker.RLock()
	defer m.locker.RUnlock()
	_, ok := m.data[key]
	return ok
}

func (m *MapRWLocker[K, V]) Set(key K, value V) {
	m.locker.Lock()
	defer m.locker.Unlock()
	m.data[key] = value
}

func (m *MapRWLocker[K, V]) Foreach(handler func(key K, value V)) {
	m.locker.RLock()
	defer m.locker.RUnlock()
	for k, v := range m.data {
		handler(k, v)
	}
}
