package findclosestvaluebst

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

// FindClosestValue finds the value in the BST that is closest to the target
func (bst *BST) FindClosestValue(target int) int {
	if bst.Root == nil {
		return 0 // or handle empty tree case as needed
	}
	return findClosestValueHelper(bst.Root, target, bst.Root.Value)
}

// This is our helper function that does all the hard work
func findClosestValueHelper(node *Node, target, closest int) int {
	// If we reach the end of our tree (like reaching the end of a treasure map),
	// we return what we found as the closest number
	if node == nil {
		return closest
	}

	// Let's check if our current closest number is really the closest!
	// We do this by comparing the differences between numbers
	currentDiff := abs(target - node.Value) // How far is current number from target?
	closestDiff := abs(target - closest)    // How far is closest number from target?

	// If we found a number that's closer to our target, let's remember it!
	if currentDiff < closestDiff {
		closest = node.Value
	}

	// Now comes the fun part! We need to decide which way to go:
	// Should we go left (to smaller numbers) or right (to bigger numbers)?

	// If our target is smaller than current number, go left
	if target < node.Value {
		return findClosestValueHelper(node.Left, target, closest)
		// If our target is bigger than current number, go right
	} else if target > node.Value {
		return findClosestValueHelper(node.Right, target, closest)
		// If we found our exact target, that's the closest possible number!
	} else {
		return closest
	}
}

// This helper function tells us how far apart two numbers are
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Create a sample tree:
	//       10
	//      /  \
	//     5    15
	//    / \   / \
	//   2   7 13  22

	bst := &BST{
		Root: &Node{
			Value: 10,
			Left: &Node{
				Value: 5,
				Left:  &Node{Value: 2},
				Right: &Node{Value: 7},
			},
			Right: &Node{
				Value: 15,
				Left:  &Node{Value: 13},
				Right: &Node{Value: 22},
			},
		},
	}

	result := bst.FindClosestValue(12)
	fmt.Printf("The closest value to 12 is: %d\n", result)
}
