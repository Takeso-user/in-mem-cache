package cache

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	c := NewCache()
	key := "userId"
	value := 42
	ttl := time.Second * 10

	c.Set(key, value, ttl)

	val, found := c.Get(key)
	assert.True(t, found, "expected key to be found")
	assert.Equal(t, value, val, "expected value to be equal")
}

func TestCache_Expiration(t *testing.T) {
	cache := NewCache()
	ttl := 1 * time.Second
	cache.Set("key1", "value1", ttl)
	time.Sleep(2 * time.Second)

	_, ok := cache.Get("key1")
	assert.False(t, ok, "expected key1 to be expired")
}

func TestGet(t *testing.T) {
	c := NewCache()
	key := "userId"
	value := 42
	ttl := time.Second * 10

	c.Set(key, value, ttl)

	val, found := c.Get(key)
	assert.True(t, found, "expected key to be found")
	assert.Equal(t, value, val, "expected value to be equal")

	val, found = c.Get("nonExistingKey")
	assert.False(t, found, "expected key to not be found")
	assert.Nil(t, val, "expected value to be nil")
}

func TestDelete(t *testing.T) {
	c := NewCache()
	key := "userId"
	value := 42
	ttl := time.Second * 10

	c.Set(key, value, ttl)

	deleted := c.Delete(key)
	assert.True(t, deleted, "expected true for existing key")

	val, found := c.Get(key)
	assert.False(t, found, "expected key to be deleted")
	assert.Nil(t, val, "expected value to be nil")

	deleted = c.Delete("nonExistingKey")
	assert.False(t, deleted, "expected false for non-existing key")
}

func TestSaveToFile(t *testing.T) {
	c := NewCache()
	key := "userId"
	var value float64 = 42
	ttl := time.Second * 10

	c.Set(key, value, ttl)

	err := c.SaveToFile("test_cache.json")
	assert.NoError(t, err, "expected no error when saving cache to file")

	defer func() {
		err := os.Remove("test_cache.json")
		if err != nil {
			t.Error("expected no error when removing test cache file")
		}
	}()
}

func TestLoadFromFile(t *testing.T) {
	c := NewCache()
	key := "userId"
	var value float64 = 42
	ttl := time.Second * 10

	c.Set(key, value, ttl)
	err := c.SaveToFile("test_cache.json")
	if err != nil {
		t.Error("expected no error when saving cache to file")
		return
	}

	defer func() {
		err := os.Remove("test_cache.json")
		if err != nil {
			t.Error("expected no error when removing test cache file")
		}
	}()

	newCache := NewCache()
	err = newCache.LoadFromFile("test_cache.json")
	assert.NoError(t, err, "expected no error when loading cache from file")

	val, found := newCache.Get(key)
	assert.True(t, found, "expected key to be found in loaded cache")
	assert.Equal(t, value, val, "expected value to be equal in loaded cache")
}
