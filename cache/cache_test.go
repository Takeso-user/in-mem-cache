package cache

import (
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
