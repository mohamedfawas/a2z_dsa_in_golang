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

func (t *Trie) findNode(prefix string) *TrieNode {
	current := t.root

	for _, char := range prefix {
		if _, exists := current.children[char]; !exists {
			return nil
		}

		current = current.children[char]
	}

	return current
}

func findAllWords(node *TrieNode, prefix string, words *[]string) {
	if node.isEndOfWord {
		*words = append(*words, prefix)
	}

	for char, childNode := range node.children {
		findAllWords(childNode, prefix+string(char), words)
	}
}

func (t *Trie) AutoComplete(prefix string) []string {
	node := t.findNode(prefix)
	if node == nil {
		return []string{}
	}

	var words []string
	findAllWords(node, prefix, &words)
	return words
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
	fmt.Println(trie.AutoComplete("ca")) // Will print: [cat car card carpet]

	// Try finding words that start with "de"
	fmt.Println("\nWords starting with 'de':")
	fmt.Println(trie.AutoComplete("de")) // Will print: [deer deal]
}
