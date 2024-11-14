package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func InorderTraversal(root *Node) {
	if root == nil {
		return
	}

	InorderTraversal(root.Left)

	fmt.Printf(" %d ", root.Value)

	InorderTraversal(root.Right)
}

func PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.Value)
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
}

func PostOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.Left)
	PostOrderTraversal(root.Right)
	fmt.Printf("%d ", root.Value)
}

func main() {
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
	PreOrderTraversal(root) // Should print: 1 2 4 5 3
	fmt.Println("\n")

	fmt.Println("Postorder traversal (Left -> Right -> Root):")
	PostOrderTraversal(root) // Should print: 4 5 2 3 1
	fmt.Println("\n")
}
