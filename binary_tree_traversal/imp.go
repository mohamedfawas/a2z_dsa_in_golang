package binarytreetraversal

import "fmt"

// Node represents each point in our tree
// It's like a box that holds a number and has two arrows pointing to other boxes
type Node struct {
	Value int   // The number stored in the box
	Left  *Node // Arrow pointing to left box
	Right *Node // Arrow pointing to right box
}

// NewNode creates a new box (node) with a number inside
func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

// InorderTraversal visits all nodes in-order: left child, parent, right child
// It's like reading a book from left to right
func InorderTraversal(root *Node) {
	// If we don't have any box to look at, we stop
	if root == nil {
		return
	}

	// Step 1: First, look at everything in the left side
	InorderTraversal(root.Left)

	// Step 2: Then, look at the current box
	fmt.Printf("%d ", root.Value)

	// Step 3: Finally, look at everything in the right side
	InorderTraversal(root.Right)
}

// PreorderTraversal visits nodes in pre-order: parent, left child, right child
// It's like saying "hi" to the parent first, then their children
func PreorderTraversal(root *Node) {
	// If we don't have any box to look at, we stop
	if root == nil {
		return
	}

	// Step 1: First, look at the current box
	fmt.Printf("%d ", root.Value)

	// Step 2: Then, look at everything in the left side
	PreorderTraversal(root.Left)

	// Step 3: Finally, look at everything in the right side
	PreorderTraversal(root.Right)
}

// PostorderTraversal visits nodes in post-order: left child, right child, parent
// It's like children going first, then the parent
func PostorderTraversal(root *Node) {
	// If we don't have any box to look at, we stop
	if root == nil {
		return
	}

	// Step 1: First, look at everything in the left side
	PostorderTraversal(root.Left)

	// Step 2: Then, look at everything in the right side
	PostorderTraversal(root.Right)

	// Step 3: Finally, look at the current box
	fmt.Printf("%d ", root.Value)
}

func main() {
	// Let's create a simple binary tree that looks like this:
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5

	// Create the boxes (nodes)
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)

	// Now let's try all three ways of walking through our tree!
	fmt.Println("Inorder traversal (Left -> Root -> Right):")
	InorderTraversal(root) // Should print: 4 2 5 1 3
	fmt.Println("\n")

	fmt.Println("Preorder traversal (Root -> Left -> Right):")
	PreorderTraversal(root) // Should print: 1 2 4 5 3
	fmt.Println("\n")

	fmt.Println("Postorder traversal (Left -> Right -> Root):")
	PostorderTraversal(root) // Should print: 4 5 2 3 1
	fmt.Println("\n")
}
