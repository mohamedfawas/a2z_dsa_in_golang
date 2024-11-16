package longestprefixtrie

import "fmt"

// TrieNode represents a magical tree branch!
type TrieNode struct {
	isEndOfWord bool               // Is this the end of a word?
	children    map[rune]*TrieNode // What are the next letters (branches)?
}

// NewTrieNode creates a new magical branch
func NewTrieNode() *TrieNode {
	return &TrieNode{
		isEndOfWord: false,                    // This branch is not the end of a word yet
		children:    make(map[rune]*TrieNode), // No branches (letters) yet
	}
}

// Trie is the whole magical tree
type Trie struct {
	root *TrieNode // The start of our tree (empty root)
}

// NewTrie creates a new magical tree
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(), // Start with an empty tree
	}
}

// Insert adds a word to the tree
func (t *Trie) Insert(word string) {
	current := t.root // Start at the root

	// Walk through each letter in the word
	for _, char := range word {
		// If there's no branch for this letter, make one
		if _, exists := current.children[char]; !exists {
			current.children[char] = NewTrieNode()
		}

		// Move to the next branch (letter)
		current = current.children[char]
	}

	// Mark the last branch as the end of the word
	current.isEndOfWord = true
}

// FindLongestPrefix finds the longest part of words that are the same at the start
func (t *Trie) FindLongestPrefix() string {
	prefix := ""      // Start with an empty prefix
	current := t.root // Start climbing the tree from the root

	// Keep climbing while the rules allow it
	for {
		// Stop if:
		// 1. There are many branches (splits in the tree).
		// 2. We've reached the end of a word.
		if len(current.children) != 1 || current.isEndOfWord {
			break
		}

		// If there’s exactly one branch, follow it
		for char, childNode := range current.children {
			prefix += string(char) // Add the letter to our prefix
			current = childNode    // Move to the next branch
			break                  // Exit the loop since there’s only one child
		}
	}

	return prefix // The longest prefix we found!
}

func main() {
	// Make a new magical tree
	trie := NewTrie()

	// Add some words to the tree
	words := []string{"flower", "flow", "flight"}
	for _, word := range words {
		trie.Insert(word) // Insert each word into the tree
	}

	// Find the longest shared prefix
	longestPrefix := trie.FindLongestPrefix()

	// Show the words and the longest prefix
	fmt.Printf("Words in Trie: %v\n", words)
	fmt.Printf("Longest prefix: %s\n", longestPrefix)
}
