package cache

import (
	"context"
	"sync"
	"time"
)

type MCache interface {
	Set(key string, value interface{}, ttl time.Duration) bool
	Get(key string) (interface{}, bool)
	Delete(key string) bool
}

type Cache struct {
	data        map[string]interface{}
	timestamps  map[string]time.Time
	expirations map[string]time.Duration
	mu          sync.RWMutex
	cancelFunc  context.CancelFunc
}

func NewCache() *Cache {
	ctx, cancel := context.WithCancel(context.Background())
	cache := &Cache{
		data:        make(map[string]interface{}),
		timestamps:  make(map[string]time.Time),
		expirations: make(map[string]time.Duration),
		cancelFunc:  cancel,
	}
	go cache.startEvWorker(ctx)
	return cache
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
	c.timestamps[key] = time.Now()
	c.expirations[key] = ttl
	return true
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, exists := c.data[key]
	if !exists {
		return nil, false
	}
	if time.Since(c.timestamps[key]) > c.expirations[key] {
		go c.Delete(key)
		return nil, false
	}
	return val, exists
}

func (c *Cache) Delete(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, existed := c.data[key]
	if existed {
		delete(c.data, key)
		delete(c.timestamps, key)
		delete(c.expirations, key)
		return true
	}
	return false
}

func (c *Cache) startEvWorker(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.evictExpiredKeys()
		case <-ctx.Done():
			return
		}
	}
}

func (c *Cache) evictExpiredKeys() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, timestamp := range c.timestamps {
		if time.Since(timestamp) > c.expirations[key] {
			delete(c.timestamps, key)
			delete(c.data, key)
			delete(c.expirations, key)
		}
	}
}

func (c *Cache) Stop() {
	c.cancelFunc()
}
