package implementtrieleetcode

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

func (this *Trie) Insert(word string) {
	current := this.root
	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			current.children[ch] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		current = current.children[ch]
	}
	current.isEnd = true
}

func (this *Trie) Search(word string) bool {
	current := this.root
	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			return false
		}
		current = current.children[ch]
	}
	return current.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this.root
	for _, ch := range prefix {
		if _, exists := current.children[ch]; !exists {
			return false
		}
		current = current.children[ch]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
