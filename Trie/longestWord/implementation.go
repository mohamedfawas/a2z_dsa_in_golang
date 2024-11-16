package longestwordusingtrie

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

func (t *Trie) Prefix(word string) bool {

	current := t.root
	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			return false
		}
		current = current.children[char]
	}

	return true
}

// FindLongestWord finds the longest word in the trie
func (t *Trie) FindLongestWord() string {
	longestWord := ""

	// Helper function to perform DFS
	var dfs func(node *TrieNode, currentWord string)
	dfs = func(node *TrieNode, currentWord string) {
		// If this is a valid word and longer than our current longest word,
		// update the longest word
		if node.isEndOfWord && len(currentWord) > len(longestWord) {
			longestWord = currentWord
		}

		// Explore all possible children
		for char, childNode := range node.children {
			dfs(childNode, currentWord+string(char))
		}
	}

	// Start DFS from root with empty string
	dfs(t.root, "")
	return longestWord
}

func main() {
	// Create a new trie
	trie := NewTrie()

	// Insert some test words
	words := []string{
		"apple",
		"app",
		"banana",
		"cat",
		"doghouse",
		"dog",
	}

	for _, word := range words {
		trie.Insert(word)
	}

	// Find the longest word
	longest := trie.FindLongestWord()
	fmt.Printf("Words in trie: %v\n", words)
	fmt.Printf("Longest word: %s\n", longest)
}
