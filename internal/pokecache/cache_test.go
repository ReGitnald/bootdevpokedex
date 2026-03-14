package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(5 * time.Second)

	key := "testKey"
	value := []byte("testValue")

	cache.Add(key, value)

	retrievedValue, exists := cache.Get(key)
	if !exists {
		t.Errorf("Expected key '%s' to exist in cache", key)
	} else if string(retrievedValue) != string(value) {
		t.Errorf("Expected value '%s', got '%s'", value, retrievedValue)
	}
}

func TestReapLoop(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)

	key := "testKey"
	value := []byte("testValue")

	cache.Add(key, value)

	// Wait for the reap loop to run
	time.Sleep(50 * time.Millisecond)

	_, exists := cache.Get(key)
	if exists {
		t.Errorf("Expected key '%s' to be reaped from cache", key)
	}

}
