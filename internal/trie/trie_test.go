package trie

import (
	"testing"
)

func TestPutAndGet(t *testing.T) {
	root := NewTrieNode()
	v1 := Put(root, "cat", "animal")
	v2 := Put(v1, "car", "vehicle")
	v3 := Put(v2, "dog", "animal")

	tests := []struct {
		name     string
		trie     *TrieNode
		key      string
		expected string
		found    bool
	}{
		{"Get from v1", v1, "cat", "animal", true},
		{"Missing from v1", v1, "car", "", false},
		{"Get from v2", v2, "car", "vehicle", true},
		{"Get cat still in v2", v2, "cat", "animal", true},
		{"Get from v3", v3, "dog", "animal", true},
		{"Get all from v3", v3, "car", "vehicle", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, ok := Get(tt.trie, tt.key)
			if ok != tt.found || val != tt.expected {
				t.Errorf("Get(%q) = (%q, %v); want (%q, %v)", tt.key, val, ok, tt.expected, tt.found)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	root := NewTrieNode()
	v1 := Put(root, "cat", "animal")
	v2 := Put(v1, "car", "vehicle")
	v3 := Put(v2, "dog", "animal")
	v4 := Delete(v3, "car")

	// Validate deletion
	val, ok := Get(v4, "car")
	if ok || val != "" {
		t.Errorf("Expected 'car' to be deleted in v4, got (%q, %v)", val, ok)
	}

	// Validate immutability
	val, ok = Get(v3, "car")
	if !ok || val != "vehicle" {
		t.Errorf("Expected 'car' to remain in v3, got (%q, %v)", val, ok)
	}

	// Validate others still exist
	if val, ok = Get(v4, "cat"); !ok || val != "animal" {
		t.Errorf("Expected 'cat' to still exist in v4, got (%q, %v)", val, ok)
	}
	if val, ok = Get(v4, "dog"); !ok || val != "animal" {
		t.Errorf("Expected 'dog' to still exist in v4, got (%q, %v)", val, ok)
	}
}

func TestDeleteNonExistentKey(t *testing.T) {
	root := NewTrieNode()
	v1 := Put(root, "cat", "animal")
	v2 := Delete(v1, "car") // 'car' doesn't exist

	val, ok := Get(v2, "cat")
	if !ok || val != "animal" {
		t.Errorf("Expected 'cat' to remain unchanged, got (%q, %v)", val, ok)
	}
}

func TestBasicPutGet(t *testing.T) {
	root := NewTrieNode()
	v1 := Put(root, "a", "1")
	v2 := Put(v1, "ab", "2")
	v3 := Put(v2, "abc", "3")

	tests := []struct {
		version  *TrieNode
		key      string
		expected string
		found    bool
	}{
		{v1, "a", "1", true},
		{v2, "ab", "2", true},
		{v3, "abc", "3", true},
		{v3, "a", "1", true},
		{v3, "ab", "2", true},
		{v3, "abcd", "", false}, // non-existent
	}

	for _, tt := range tests {
		val, ok := Get(tt.version, tt.key)
		if ok != tt.found || val != tt.expected {
			t.Errorf("Get(%q) = (%q, %v); want (%q, %v)", tt.key, val, ok, tt.expected, tt.found)
		}
	}
}

func TestDeleteCascading(t *testing.T) {
	root := NewTrieNode()
	v1 := Put(root, "abc", "1")
	v2 := Put(v1, "abcd", "2")
	v3 := Put(v2, "abce", "3")

	v4 := Delete(v3, "abcd") // Deletes one child
	v5 := Delete(v4, "abc")  // Should not delete "abce"
	v6 := Delete(v5, "abce") // Entire path deleted

	if _, ok := Get(v4, "abcd"); ok {
		t.Errorf("Expected abcd to be deleted")
	}
	if val, ok := Get(v4, "abc"); !ok || val != "1" {
		t.Errorf("Expected abc to still exist, got (%q, %v)", val, ok)
	}
	if val, ok := Get(v4, "abce"); !ok || val != "3" {
		t.Errorf("Expected abce to still exist, got (%q, %v)", val, ok)
	}

	if _, ok := Get(v5, "abc"); ok {
		t.Errorf("Expected abc to be deleted")
	}
	if val, ok := Get(v5, "abce"); !ok || val != "3" {
		t.Errorf("Expected abce to still exist, got (%q, %v)", val, ok)
	}

	if _, ok := Get(v6, "abce"); ok {
		t.Errorf("Expected abce to be deleted")
	}
}

func TestPrefixShadowing(t *testing.T) {
	root := NewTrieNode()
	v1 := Put(root, "a", "A")
	v2 := Put(v1, "ab", "AB")
	v3 := Put(v2, "abc", "ABC")
	v4 := Put(v3, "abcd", "ABCD")

	// Delete middle node
	v5 := Delete(v4, "ab")

	if _, ok := Get(v5, "ab"); ok {
		t.Errorf("Expected ab to be deleted")
	}
	if val, ok := Get(v5, "abc"); !ok || val != "ABC" {
		t.Errorf("Expected abc to remain, got (%q, %v)", val, ok)
	}
	if val, ok := Get(v5, "abcd"); !ok || val != "ABCD" {
		t.Errorf("Expected abcd to remain, got (%q, %v)", val, ok)
	}
}

func TestImmutability(t *testing.T) {
	root := NewTrieNode()
	v1 := Put(root, "hello", "world")
	v2 := Delete(v1, "hello")

	val1, ok1 := Get(v1, "hello")
	val2, ok2 := Get(v2, "hello")

	if !ok1 || val1 != "world" {
		t.Errorf("v1 should have 'hello' => 'world', got (%q, %v)", val1, ok1)
	}
	if ok2 {
		t.Errorf("v2 should have deleted 'hello', got (%q, %v)", val2, ok2)
	}
}
