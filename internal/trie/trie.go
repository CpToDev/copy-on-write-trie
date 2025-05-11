package trie

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
	value    string
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
		value:    "",
	}
}

func Clone(n *TrieNode) *TrieNode {
	newChildren := make(map[rune]*TrieNode)
	for k, v := range n.children {
		newChildren[k] = v
	}
	return &TrieNode{
		children: newChildren,
		isEnd:    n.isEnd,
		value:    n.value,
	}
}

func Get(root *TrieNode, key string) (string, bool) {
	curr := root
	for _, ch := range key {
		_, ok := curr.children[ch]
		if !ok {
			return "", false
		} else {
			curr = curr.children[ch]
		}
	}
	return curr.value, curr.isEnd
}

func Put(root *TrieNode, key string, value string) *TrieNode {
	newRoot := Clone(root)
	curr := newRoot
	for i, ch := range key {
		_, ok := curr.children[ch]
		if !ok {
			newChild := NewTrieNode()
			if i == len(key)-1 {
				newChild.isEnd = true
				newChild.value = value
			}
			curr.children[ch] = newChild
			curr = curr.children[ch]
		} else {
			curr.children[ch] = Clone(curr.children[ch])
			if len(key)-1 == i {
				curr.children[ch].isEnd = true
				curr.children[ch].value = value
			}
			curr = curr.children[ch]
		}
	}
	return newRoot
}
func DeleteHelper(node *TrieNode, key string, i int) (*TrieNode, bool) {

	if i == len(key) {
		if node.isEnd {
			if len(node.children) > 0 {
				clone := Clone(node)
				clone.isEnd = false
				clone.value = ""
				return clone, true

			} else {
				return nil, true
			}
		} else {
			return node, false
		}

	}
	_, ok := node.children[rune(key[i])]
	if ok {
		child, exists := DeleteHelper(node.children[rune(key[i])], key, i+1)
		if exists {
			clone := Clone(node)
			if child == nil {
				delete(clone.children, rune(key[i]))
			} else {
				clone.children[rune(key[i])] = child
			}

			return clone, true
		} else {
			return node, false
		}
	} else {
		return nil, false
	}

}
func Delete(root *TrieNode, key string) *TrieNode {
	newNode, _ := DeleteHelper(root, key, 0)
	return newNode
}
