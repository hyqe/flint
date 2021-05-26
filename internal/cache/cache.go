package cache

import "sync"

type Cacher interface {
	Get(k string) (interface{}, bool)
	Put(k string, v interface{})
	Delete(k string) bool
}

func New() Cacher {
	return &cache{
		db: make(map[string]interface{}),
	}
}

type cache struct {
	sync.RWMutex
	db map[string]interface{}
}

func (c *cache) Get(k string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	v, ok := c.db[k]
	return v, ok
}

func (c *cache) Put(k string, v interface{}) {
	c.Lock()
	defer c.Unlock()
	c.db[k] = v
}

func (c *cache) Delete(k string) bool {
	c.Lock()
	defer c.Unlock()
	_, ok := c.db[k]
	if !ok {
		return false
	}
	delete(c.db, k)
	return true
}
