package cache

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

type MCache interface {
	Set(key string, value interface{}, ttl time.Duration) bool
	Get(key string) (interface{}, bool)
	Delete(key string) bool
	SaveToFile(filename string) error
	LoadFromFile(filename string) error
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

type cacheData struct {
	Data        map[string]interface{}
	Timestamps  map[string]time.Time
	Expirations map[string]time.Duration
}

func (c *Cache) SaveToFile(filename string) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	data := cacheData{
		Data:        c.data,
		Timestamps:  c.timestamps,
		Expirations: c.expirations,
	}

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}

func (c *Cache) LoadFromFile(filename string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var data cacheData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	c.data = data.Data
	c.timestamps = data.Timestamps
	c.expirations = data.Expirations
	return nil
}
