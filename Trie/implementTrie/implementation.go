package trieimplementation

import "fmt"

// TrieNode represents each node in our Trie
// Think of it as a box that can hold:
// - a letter (but we don't actually need to store it)
// - a flag telling us if a word ends here
// - links to other letter boxes (children)
type TrieNode struct {
	// isEndOfWord tells us if this node represents the end of a word
	// Like putting a special mark at the end of a word
	isEndOfWord bool

	// children stores links to next letters
	// Think of it as 26 doors, each leading to another letter
	// We use a map instead of an array to save space
	children map[rune]*TrieNode
}

// NewTrieNode creates a new empty box (node) for our letters
func NewTrieNode() *TrieNode {
	return &TrieNode{
		isEndOfWord: false,
		children:    make(map[rune]*TrieNode),
	}
}

// Trie is our main word storage structure
// It has a special starting box (root) from where all words begin
type Trie struct {
	root *TrieNode
}

// NewTrie creates a new empty word storage (Trie)
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

// Insert adds a new word to our Trie
// Imagine drawing a path of letters, one by one
func (t *Trie) Insert(word string) {
	// Start from the first box (root)
	current := t.root

	// For each letter in the word
	for _, char := range word {
		// If there's no path for this letter yet
		if _, exists := current.children[char]; !exists {
			// Create a new box for this letter
			current.children[char] = NewTrieNode()
		}
		// Move to the next letter's box
		current = current.children[char]
	}

	// Mark this box as the end of a word
	current.isEndOfWord = true
}

// Search looks for a complete word in our Trie
// Returns true if the word exists, false otherwise
func (t *Trie) Search(word string) bool {
	// Start from the first box
	current := t.root

	// Follow the path of letters
	for _, char := range word {
		// If we can't find the next letter
		if _, exists := current.children[char]; !exists {
			// The word doesn't exist in our Trie
			return false
		}
		// Move to the next letter's box
		current = current.children[char]
	}

	// We found all letters, but is it a complete word?
	// Check if this is marked as the end of a word
	return current.isEndOfWord
}

// StartsWith checks if there are any words that begin with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	// Start from the first box
	current := t.root

	// Follow the path of letters in the prefix
	for _, char := range prefix {
		// If we can't find the next letter
		if _, exists := current.children[char]; !exists {
			// No words start with this prefix
			return false
		}
		// Move to the next letter's box
		current = current.children[char]
	}

	// If we get here, we found all letters of the prefix
	return true
}

// Example usage
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
