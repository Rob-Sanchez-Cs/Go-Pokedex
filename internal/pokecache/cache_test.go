package pokecache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	c := NewCache(1 * time.Second)
	if c == nil {
		t.Fatal("expected NewCache to return a non-nil Cache")
	}
	if c.data == nil {
		t.Fatal("expected cache data map to be initialized")
	}
}

func TestAddAndGet(t *testing.T) {
	c := NewCache(1 * time.Minute)
	key := "test-key"
	value := []byte("test-value")

	c.Add(key, value)

	got, ok := c.Get(key)
	if !ok {
		t.Fatal("expected key to exist in cache")
	}
	if string(got) != string(value) {
		t.Errorf("expected value %s, got %s", value, got)
	}
}

func TestGetMissingKey(t *testing.T) {
	c := NewCache(1 * time.Minute)
	_, ok := c.Get("non-existent-key")
	if ok {
		t.Fatal("expected key not to be found in cache")
	}
}