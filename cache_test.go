package cache

import (
	"testing"
)

func TestSet(t *testing.T) {
	c := NewCache()
	key := "userId"
	value := 42

	c.Set(key, value)

	if val, found := c.Get(key); !found || val != value {
		t.Errorf("Set() failed: expected %v, got %v", value, val)
	}
}

func TestGet(t *testing.T) {
	c := NewCache()
	key := "userId"
	value := 42

	c.Set(key, value)

	if val, found := c.Get(key); !found || val != value {
		t.Errorf("Get() failed: expected %v, got %v", value, val)
	}

	if val, found := c.Get("nonExistingKey"); found || val != nil {
		t.Errorf("Get() for non-existing key failed: expected nil, got %v", val)
	}
}

func TestDelete(t *testing.T) {
	c := NewCache()
	key := "userId"
	value := 42

	c.Set(key, value)

	if deleted := c.Delete(key); !deleted {
		t.Errorf("Delete() failed: expected true for existing key")
	}

	if val, found := c.Get(key); found || val != nil {
		t.Errorf("Delete() failed: expected key to be deleted, but got %v", val)
	}

	if deleted := c.Delete("nonExistingKey"); deleted {
		t.Errorf("Delete() failed: expected false for non-existing key")
	}
}
