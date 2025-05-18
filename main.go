package main

import (
	"fmt"

	"github.com/sauravagtl/copyonwrite-trie-go/internal/trie"
)

func main() {
	trieStore := trie.NewConcurrentTrieStore()
	trieStore.Put("1", "saurav")
	if valueGuard, ok := trieStore.Get("1"); ok {
		fmt.Printf("%v", valueGuard.Value)
	}

}
