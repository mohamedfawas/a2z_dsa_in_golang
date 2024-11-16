package main

import "fmt"

// TrieNode represents a node in the Trie
type TrieNode struct {
	isEndOfWord bool
	children    map[rune]*TrieNode
}

// NewTrieNode creates a new Trie node
func NewTrieNode() *TrieNode {
	return &TrieNode{
		isEndOfWord: false,
		children:    make(map[rune]*TrieNode),
	}
}

// Trie represents the Trie data structure
type Trie struct {
	root *TrieNode
}

// NewTrie creates a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

// Insert adds a word to the Trie
func (t *Trie) Insert(word string) {
	current := t.root

	// Traverse each character of the word
	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			current.children[char] = NewTrieNode()
		}
		current = current.children[char]
	}

	// Mark the end of the word
	current.isEndOfWord = true
}

// FindLongestPrefix returns the longest prefix shared by all words
func (t *Trie) FindLongestPrefix() string {
	prefix := ""
	current := t.root

	// Traverse the Trie until we reach a node with more than one child
	// or a node that is the end of a word
	for {
		// If the current node has more than one child or is the end of a word, stop
		if len(current.children) != 1 || current.isEndOfWord {
			break
		}

		// Get the single child (since len(current.children) == 1)
		for char, childNode := range current.children {
			prefix += string(char)
			current = childNode
			break
		}
	}

	return prefix
}

func main() {
	// Create a new Trie
	trie := NewTrie()

	// Insert some test words
	words := []string{"flower", "flow", "flight"}
	for _, word := range words {
		trie.Insert(word)
	}

	// Find the longest prefix
	longestPrefix := trie.FindLongestPrefix()

	fmt.Printf("Words in Trie: %v\n", words)
	fmt.Printf("Longest prefix: %s\n", longestPrefix)
}
