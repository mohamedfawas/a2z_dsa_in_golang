package main

// Node represents a single node in the binary search tree (BST).
type Node struct {
	Value int   // Value stored in the node
	Left  *Node // Pointer to the left child
	Right *Node // Pointer to the right child
}

// BST represents the binary search tree with a root node.
type BST struct {
	Root *Node // Root node of the BST
}

// NewNode creates a new node with the given value.
func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func InOrderHelper(node *Node, result *[]int) {

}
