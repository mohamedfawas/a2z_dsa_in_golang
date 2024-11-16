package bstimplementation

import "fmt"

// Node represents each piece in our tree, like a building block
// It has a value and can point to two other blocks (left and right)
type Node struct {
	Value int   // The number stored in this block
	Left  *Node // Points to the left block (smaller numbers)
	Right *Node // Points to the right block (bigger numbers)
}

// BST is our Binary Search Tree
// It's like a container that holds all our blocks, starting from the root
type BST struct {
	Root *Node // The first block in our tree (the top block)
}

// Insert adds a new number to our tree
func (bst *BST) Insert(value int) {
	// If the tree is empty, make this our first block
	if bst.Root == nil {
		bst.Root = &Node{Value: value}
		return
	}

	// If we already have blocks, let's find where to put the new one
	insertHelper(bst.Root, value)
}

// insertHelper helps us find the right spot for our new number
func insertHelper(node *Node, value int) {
	// If the new number is smaller, go to the left
	if value < node.Value {
		// If there's no block on the left, put it there
		if node.Left == nil {
			node.Left = &Node{Value: value}
			return
		}
		// If there is a block, keep looking on the left side
		insertHelper(node.Left, value)
	} else {
		// If the new number is bigger or equal, go to the right
		if node.Right == nil {
			node.Right = &Node{Value: value}
			return
		}
		// If there is a block, keep looking on the right side
		insertHelper(node.Right, value)
	}
}

// Contains checks if a number exists in our tree
func (bst *BST) Contains(value int) bool {
	return containsHelper(bst.Root, value)
}

// containsHelper helps us search for a number in the tree
func containsHelper(node *Node, value int) bool {
	// If we hit a dead end, the number isn't in our tree
	if node == nil {
		return false
	}

	// If we found the number, hooray!
	if node.Value == value {
		return true
	}

	// If our number is smaller, look on the left side
	if value < node.Value {
		return containsHelper(node.Left, value)
	}
	// If our number is bigger, look on the right side
	return containsHelper(node.Right, value)
}

// Delete removes a number from our tree
func (bst *BST) Delete(value int) {
	bst.Root = deleteHelper(bst.Root, value)
}

// deleteHelper helps us remove a number from the tree
func deleteHelper(node *Node, value int) *Node {
	// If we hit a dead end, nothing to delete
	if node == nil {
		return nil
	}

	// If we found our number, let's remove it!
	if value == node.Value {
		// Case 1: If the block has no children, just remove it
		if node.Left == nil && node.Right == nil {
			return nil
		}
		// Case 2: If the block has only one child, connect its parent to its child
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		// Case 3: If the block has two children, find the smallest number
		// on the right side and put it in place of the removed number
		smallest := findSmallest(node.Right)
		node.Value = smallest
		node.Right = deleteHelper(node.Right, smallest)
		return node
	}

	// Keep searching for our number
	if value < node.Value {
		node.Left = deleteHelper(node.Left, value)
	} else {
		node.Right = deleteHelper(node.Right, value)
	}
	return node
}

// findSmallest helps us find the smallest number in a subtree
func findSmallest(node *Node) int {
	current := node
	// Keep going left until we can't anymore
	for current.Left != nil {
		current = current.Left
	}
	return current.Value
}

// PreOrder visits root, then left, then right
// Like reading a book: title first, then left page, then right page
func (bst *BST) PreOrder() {
	preOrderHelper(bst.Root)
	fmt.Println() // Add a new line at the end
}

func preOrderHelper(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Value) // Print current number
	preOrderHelper(node.Left)     // Visit left side
	preOrderHelper(node.Right)    // Visit right side
}

// InOrder visits left, then root, then right
// Like reading numbers in order from smallest to biggest
func (bst *BST) InOrder() {
	inOrderHelper(bst.Root)
	fmt.Println() // Add a new line at the end
}

func inOrderHelper(node *Node) {
	if node == nil {
		return
	}
	inOrderHelper(node.Left)      // Visit left side
	fmt.Printf("%d ", node.Value) // Print current number
	inOrderHelper(node.Right)     // Visit right side
}

// PostOrder visits left, then right, then root
// Like cleaning up: clean left room, clean right room, then clean parent's room
func (bst *BST) PostOrder() {
	postOrderHelper(bst.Root)
	fmt.Println() // Add a new line at the end
}

func postOrderHelper(node *Node) {
	if node == nil {
		return
	}
	postOrderHelper(node.Left)    // Visit left side
	postOrderHelper(node.Right)   // Visit right side
	fmt.Printf("%d ", node.Value) // Print current number
}

// Example usage
func main() {
	// Create a new tree
	tree := &BST{}

	// Add some numbers
	numbers := []int{5, 3, 7, 1, 4, 6, 8}
	for _, num := range numbers {
		tree.Insert(num)
	}

	fmt.Println("Tree traversals:")
	fmt.Print("PreOrder: ")
	tree.PreOrder() // Should print: 5 3 1 4 7 6 8
	fmt.Print("InOrder: ")
	tree.InOrder() // Should print: 1 3 4 5 6 7 8
	fmt.Print("PostOrder: ")
	tree.PostOrder() // Should print: 1 4 3 6 8 7 5

	// Check if numbers exist
	fmt.Printf("Contains 4? %v\n", tree.Contains(4)) // true
	fmt.Printf("Contains 9? %v\n", tree.Contains(9)) // false

	// Delete a number and show the tree again
	fmt.Println("\nAfter deleting 3:")
	tree.Delete(3)
	fmt.Print("InOrder: ")
	tree.InOrder() // Should print: 1 4 5 6 7 8
}
