package memory

import "sync"

type Lookup[T any] struct {
	mu  sync.RWMutex
	Map map[string]T
}

func NewLookup[T any]() *Lookup[T] {
	return &Lookup[T]{
		Map: make(map[string]T),
	}
}

func (l *Lookup[T]) Get(key string) (T, bool) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	v, ok := l.Map[key]
	return v, ok
}

func (l *Lookup[T]) Put(key string, value T) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Map[key] = value
}

func (l *Lookup[T]) Delete(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.Map, key)
}
