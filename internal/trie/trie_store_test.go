package trie

import (
	"sync"
	"testing"
)

func TestBasicOperationsConcurrentTrieStore(t *testing.T) {
	store := NewConcurrentTrieStore()

	store.Put("apple", "fruit")
	store.Put("car", "vehicle")

	v, ok := store.Get("apple")
	if !ok || v.Value != "fruit" {
		t.Errorf("Expected apple -> fruit, got %v", v.Value)
	}

	v, ok = store.Get("car")
	if !ok || v.Value != "vehicle" {
		t.Errorf("Expected car -> vehicle, got %v", v.Value)
	}

	// Overwrite
	store.Put("apple", "company")
	v, ok = store.Get("apple")
	if !ok || v.Value != "company" {
		t.Errorf("Expected apple -> company, got %v", v.Value)
	}
}

func TestDeleteConcurrentTrieStore(t *testing.T) {
	store := NewConcurrentTrieStore()
	store.Put("banana", "yellow")

	v, ok := store.Get("banana")
	if !ok || v.Value != "yellow" {
		t.Errorf("Expected banana -> yellow, got %v", v.Value)
	}

	store.Delete("banana")
	if _, ok := store.Get("banana"); ok {
		t.Error("Expected banana to be deleted")
	}

	// Delete non-existent key
	store.Delete("nonexistent") // should not panic
}

func TestConcurrentAccess(t *testing.T) {
	store := NewConcurrentTrieStore()
	keys := []string{"a", "b", "c", "d", "e"}
	var wg sync.WaitGroup

	// Writers
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			for _, k := range keys {
				store.Put(k, k)
			}
		}
	}()

	// Readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				for _, k := range keys {
					_, _ = store.Get(k)
				}
			}
		}(i)
	}

	wg.Wait()
}

func TestValueGuardPersistence(t *testing.T) {
	store := NewConcurrentTrieStore()
	store.Put("earth", "planet")

	v1, ok := store.Get("earth")
	if !ok || v1.Value != "planet" {
		t.Errorf("Expected earth -> planet, got %v", v1.Value)
	}

	// Mutate the trie afterwards
	store.Put("earth", "blue planet")
	store.Delete("earth")

	// v1 should still hold original value
	if v1.Value != "planet" {
		t.Errorf("ValueGuard should preserve value, got %v", v1.Value)
	}
}
