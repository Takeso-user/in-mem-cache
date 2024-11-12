package cache

import "sync"

type MCache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
	Delete(key string) bool
}

type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
	return true
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, exists := c.data[key]
	return val, exists
}

func (c *Cache) Delete(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, existed := c.data[key]
	if existed {
		delete(c.data, key)
	}
	return existed
}
