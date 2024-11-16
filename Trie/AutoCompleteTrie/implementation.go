package trieautocomplete

import "fmt"

// Node represents each letter in our word tree
// Think of it like a box that can hold a letter and point to other letter boxes
type Node struct {
	// isEnd tells us if this letter is the end of a word
	// Like putting a special mark at the end of a word
	isEnd bool

	// children holds all possible next letters
	// Imagine it like having 26 different paths we can take from this letter
	children map[rune]*Node
}

// NewNode creates a new empty letter box
func NewNode() *Node {
	return &Node{
		isEnd:    false,
		children: make(map[rune]*Node),
	}
}

// Trie is like a big tree made up of letter boxes
// It helps us organize words in a way that makes them easy to find
type Trie struct {
	root *Node
}

// NewTrie creates a new empty word tree
func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

// Insert adds a new word to our tree
// Imagine walking through the tree and creating new paths for each letter
func (t *Trie) Insert(word string) {
	// Start at the beginning of our tree
	current := t.root

	// For each letter in the word
	for _, char := range word {
		// If we don't have a path for this letter yet
		if _, exists := current.children[char]; !exists {
			// Create a new path (new letter box)
			current.children[char] = NewNode()
		}
		// Move to the next letter box
		current = current.children[char]
	}
	// Mark this as the end of a word
	current.isEnd = true
}

// findNode helps us find where a prefix ends in our tree
func (t *Trie) findNode(prefix string) *Node {
	current := t.root

	// Walk through each letter of the prefix
	for _, char := range prefix {
		// If we can't find the next letter, return nil (nothing found)
		if _, exists := current.children[char]; !exists {
			return nil
		}
		// Move to the next letter
		current = current.children[char]
	}
	return current
}

// findAllWords helps us collect all words that start with our prefix
// It's like following all possible paths from where our prefix ends
func findAllWords(node *Node, prefix string, words *[]string) {
	// If we've reached the end of a word, add it to our list
	if node.isEnd {
		*words = append(*words, prefix)
	}

	// Look at all possible next letters
	for char, childNode := range node.children {
		// Follow each path and collect words
		findAllWords(childNode, prefix+string(char), words)
	}
}

// Autocomplete finds all words that start with a given prefix
func (t *Trie) Autocomplete(prefix string) []string {
	// Find where our prefix ends in the tree
	node := t.findNode(prefix)
	if node == nil {
		return []string{}
	}

	// Collect all possible words from this point
	var words []string
	findAllWords(node, prefix, &words)
	return words
}

func main() {
	// Create a new word tree
	trie := NewTrie()

	// Add some words to our tree
	words := []string{
		"cat",
		"car",
		"card",
		"carpet",
		"dog",
		"deer",
		"deal",
	}

	for _, word := range words {
		trie.Insert(word)
	}

	// Try finding words that start with "ca"
	fmt.Println("Words starting with 'ca':")
	fmt.Println(trie.Autocomplete("ca")) // Will print: [cat car card carpet]

	// Try finding words that start with "de"
	fmt.Println("\nWords starting with 'de':")
	fmt.Println(trie.Autocomplete("de")) // Will print: [deer deal]
}
