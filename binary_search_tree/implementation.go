package binarysearchtree

// Node represents a single node in the Binary Search Tree
type Node struct {
	Value int   // The value stored in the node
	Left  *Node // Pointer to left child node
	Right *Node // Pointer to right child node
}

// BST represents the Binary Search Tree
type BST struct {
	Root *Node // Pointer to the root node of the tree
}

func (t *BST) Insert(value int) {
	// If tree is empty, create root node
	if t.Root == nil {
		t.Root = &Node{Value: value}
		return
	}

	// Start at root and traverse to find insertion point
	current := t.Root
	for {
		// If value is less than current node's value, go left
		if value < current.Value {
			// If left child is nil, insert new node here
			if current.Left == nil {
				current.Left = &Node{Value: value}
				return
			}
			// Otherwise, continue traversing left
			current = current.Left
		} else {
			// If value is greater or equal, go right
			if current.Right == nil {
				current.Right = &Node{Value: value}
				return
			}
			// Otherwise, continue traversing right
			current = current.Right
		}
	}
}

