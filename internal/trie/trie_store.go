package trie

import "sync/atomic"

type ValueGuard struct {
	Value string
	Root  *TrieNode
}

type ConcurrentTrieStore struct {
	value atomic.Value
}

func NewConcurrentTrieStore() *ConcurrentTrieStore {
	store := &ConcurrentTrieStore{}
	store.value.Store(NewTrieNode())
	return store

}

func (c *ConcurrentTrieStore) Get(key string) (*ValueGuard, bool) {
	root := c.value.Load().(*TrieNode)
	val, exists := Get(root, key)
	if !exists {
		return &ValueGuard{}, false
	} else {
		return &ValueGuard{Root: root, Value: val}, true
	}
}

func (c *ConcurrentTrieStore) Put(key string, value string) {
	for {
		oldRoot := c.value.Load().(*TrieNode)
		newRoot := Put(oldRoot, key, value)
		if c.value.CompareAndSwap(oldRoot, newRoot) {
			return
		}
	}
}
func (c *ConcurrentTrieStore) Delete(key string) {
	for {
		oldRoot := c.value.Load().(*TrieNode)
		newRoot := Delete(oldRoot, key)
		if c.value.CompareAndSwap(oldRoot, newRoot) {
			return
		}
	}
}
