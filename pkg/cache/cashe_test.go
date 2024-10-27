package cache

import (
	"testing"
)

func TestCache(t *testing.T) {
	cache := NewCache()
	cache.Set("key", "value", 600)

	val, err := cache.Get("key")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if val != "value" {
		t.Errorf("Expected 'value', got %s", val)
	}

	_, err = cache.Get("nonexistent")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
