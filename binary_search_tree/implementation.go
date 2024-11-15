package main

import "fmt"

// Node represents a node in the Binary Search Tree
// Each node has a value and pointers to left and right children
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BST represents the Binary Search Tree
type BST struct {
	Root *Node
}

// Insert adds a new value to the BST
func (bst *BST) Insert(value int) {
	// If tree is empty, create root node
	if bst.Root == nil {
		bst.Root = &Node{Value: value}
		return
	}

	// Start at root and traverse down to find correct position
	current := bst.Root
	for {
		// If value is less than current node, go left
		if value < current.Value {
			// If left child is nil, insert here
			if current.Left == nil {
				current.Left = &Node{Value: value}
				return
			}
			// Otherwise, move to left child
			current = current.Left
		} else {
			// If value is greater or equal, go right
			if current.Right == nil {
				current.Right = &Node{Value: value}
				return
			}
			// Otherwise, move to right child
			current = current.Right
		}
	}
}

// Search looks for a value in the BST and returns true if found
func (bst *BST) Search(value int) bool {
	current := bst.Root

	// Traverse until we find the value or reach a nil node
	for current != nil {
		if value == current.Value {
			return true
		} else if value < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return false
}

// findMin finds the minimum value in the subtree rooted at node
func findMin(node *Node) *Node {
	current := node
	// Keep going left until we hit nil
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Delete removes a value from the BST
func (bst *BST) Delete(value int) {
	bst.Root = deleteNode(bst.Root, value)
}

// deleteNode is a helper function that recursively deletes a node
func deleteNode(node *Node, value int) *Node {
	// Base case: if node is nil, return nil
	if node == nil {
		return nil
	}

	// Recursively search for the node to delete
	if value < node.Value {
		node.Left = deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteNode(node.Right, value)
	} else {
		// Case 1: Node has no children (leaf node)
		if node.Left == nil && node.Right == nil {
			return nil
		}

		// Case 2: Node has only one child
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}

		// Case 3: Node has two children
		// Find the smallest value in right subtree (successor)
		minNode := findMin(node.Right)
		// Copy successor's value to current node
		node.Value = minNode.Value
		// Delete the successor
		node.Right = deleteNode(node.Right, minNode.Value)
	}
	return node
}

// InorderTraversal performs inorder traversal of the BST
func (bst *BST) InorderTraversal() {
	inorderHelper(bst.Root)
	fmt.Println() // New line after traversal
}

// inorderHelper is a recursive helper function for inorder traversal
func inorderHelper(node *Node) {
	if node != nil {
		// Traverse left subtree
		inorderHelper(node.Left)
		// Print current node
		fmt.Printf("%d ", node.Value)
		// Traverse right subtree
		inorderHelper(node.Right)
	}
}

func main() {
	// Create a new BST
	bst := &BST{}

	// Insert some values
	values := []int{5, 3, 7, 1, 4, 6, 8}
	for _, v := range values {
		bst.Insert(v)
	}

	// Demonstrate usage
	fmt.Println("Inorder traversal:")
	bst.InorderTraversal() // Should print: 1 3 4 5 6 7 8

	fmt.Println("Searching for 4:", bst.Search(4)) // Should print: true
	fmt.Println("Searching for 9:", bst.Search(9)) // Should print: false

	fmt.Println("Deleting 3...")
	bst.Delete(3)
	fmt.Println("Inorder traversal after deletion:")
	bst.InorderTraversal() // Should print: 1 4 5 6 7 8
}
