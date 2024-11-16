package main

import "fmt"

type TrieNode struct {
	isEndOfWord bool
	children    map[rune]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		isEndOfWord: false,
		children:    make(map[rune]*TrieNode),
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	current := t.root

	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			current.children[char] = NewTrieNode()
		}

		current = current.children[char]
	}

	current.isEndOfWord = true
}

func (t *Trie) Search(word string) bool {
	current := t.root

	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			return false
		}
		current = current.children[char]
	}

	return current.isEndOfWord
}

func (t *Trie) StartsWith(word string) bool {
	current := t.root

	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			return false
		}
		current = current.children[char]
	}

	return true
}

func main() {
	// Create a new empty Trie
	trie := NewTrie()

	// Add some words
	trie.Insert("hello")
	trie.Insert("world")
	trie.Insert("hi")

	// Try searching for words
	fmt.Println("Searching for 'hello':", trie.Search("hello")) // true
	fmt.Println("Searching for 'world':", trie.Search("world")) // true
	fmt.Println("Searching for 'hi':", trie.Search("hi"))       // true
	fmt.Println("Searching for 'hey':", trie.Search("hey"))     // false

	// Try prefix searches
	fmt.Println("Prefix 'he':", trie.StartsWith("he")) // true
	fmt.Println("Prefix 'wo':", trie.StartsWith("wo")) // true
	fmt.Println("Prefix 'ha':", trie.StartsWith("ha")) // false
}
