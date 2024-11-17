package findkthlargestbst

import "fmt"

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

// InOrderTraversal performs an in-order traversal of the BST,
// storing the values of the nodes in sorted order in the 'elements' slice.
func InOrderTraversal(node *Node, elements *[]int) {
	if node == nil {
		return // Base case: if the node is nil, stop recursion.
	}

	InOrderTraversal(node.Left, elements)     // Visit left subtree.
	*elements = append(*elements, node.Value) // Add current node value.
	InOrderTraversal(node.Right, elements)    // Visit right subtree.
}

// KthLargest finds the k-th largest element in the BST.
func (bst *BST) KthLargest(k int) int {
	if bst.Root == nil {
		fmt.Println("tree is empty")
		return -1 // Return -1 if the BST is empty.
	}

	elements := []int{}                   // Slice to store the in-order traversal result.
	InOrderTraversal(bst.Root, &elements) // Perform in-order traversal.

	n := len(elements) // Total number of elements in the tree.
	if k > n || k <= 0 {
		fmt.Println("invalid k")
		return -1 // Return -1 if k is invalid (out of range).
	}

	return elements[n-k] // Return the k-th largest element.
}
