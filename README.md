# Concurrent Copy-On-Write Trie Store (Go)

A high-performance, **concurrent key-value store** built on top of a **copy-on-write Trie** (prefix tree) in Golang. Designed to handle **multiple readers** and a **single writer** concurrently without data loss, making it ideal for in-memory KV caches in multi-threaded applications.

---

## 🚀 Features

- ✅ Immutable, functional **copy-on-write Trie** core
- ✅ **Thread-safe** support for:
  - `Get(key)` – concurrent-safe reads
  - `Put(key, value)` – updates Trie root atomically
  - `Delete(key)` – removes a key from the store
- ✅ Minimal locking – supports readers accessing old roots during writes
- ✅ `ValueGuard`: safely holds a reference to value even after concurrent trie mutations
- ✅ Unit tested with edge cases and concurrency scenarios

---

## 📂 Project Structure

trie/
├── trie.go # Basic immutable trie implementation
├── concurrent_store.go # Concurrent KV wrapper using atomic root swapping
├── value_guard.go # Holds reference to value + Trie snapshot
├── store_test.go # Unit tests for store correctness and concurrency

---

## 🧠 Core Concepts

### ✅ Copy-On-Write Trie

- Each `Put` or `Delete` returns a **new Trie root**.
- Old versions are untouched, supporting **safe concurrent reads**.

### ✅ Concurrency Model

- Internally uses **atomic.Pointer** to swap Trie root.
- No locking on reads.
- Single writer can update the root safely.

### ✅ ValueGuard

- Prevents dangling references when a value node is removed.
- Holds both value reference and root snapshot it came from.

---

## 🛠️ Usage

```go
store := NewConcurrentTrieStore()

store.Put("foo", "bar")

valueGuard, ok := store.Get("foo")
if ok {
    fmt.Println("Value:", valueGuard.Value())
}

store.Delete("foo")

```

## 📌 Future Work

Here are some planned enhancements to improve functionality and scalability:

- **Persistent Storage Integration**  
  Enable writing Trie snapshots to disk or databases for durability and recovery across sessions.

- **Eviction Policies for Memory Limits**  
  Implement LRU (Least Recently Used) or LFU (Least Frequently Used) eviction strategies to maintain bounded memory usage for large-scale deployments.

- **Prefix-Based Querying**  
  Allow retrieval of all keys or values matching a given prefix, leveraging the trie structure for efficient lookups.
