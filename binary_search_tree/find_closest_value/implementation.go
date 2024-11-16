package findclosestvaluebst

import "fmt"

// First, let's create a special box (struct) that will help us make our tree
type BinarySearchTree struct {
	Value int               // This is like putting a number in a box
	Left  *BinarySearchTree // This points to a smaller box on the left
	Right *BinarySearchTree // This points to a bigger box on the right
}

// This is our main function that will help us find the closest number
func (tree *BinarySearchTree) FindClosestValue(target int) int {
	// We'll start by saying the closest value is the very first number we see
	return findClosestValueHelper(tree, target, tree.Value)
}

// This is our helper function that does all the hard work
func findClosestValueHelper(tree *BinarySearchTree, target, closest int) int {
	// If we reach the end of our tree (like reaching the end of a treasure map),
	// we return what we found as the closest number
	if tree == nil {
		return closest
	}

	// Let's check if our current closest number is really the closest!
	// We do this by comparing the differences between numbers
	currentDiff := abs(target - tree.Value) // How far is current number from target?
	closestDiff := abs(target - closest)    // How far is closest number from target?

	// If we found a number that's closer to our target, let's remember it!
	if currentDiff < closestDiff {
		closest = tree.Value
	}

	// Now comes the fun part! We need to decide which way to go:
	// Should we go left (to smaller numbers) or right (to bigger numbers)?

	// If our target is smaller than current number, go left
	if target < tree.Value {
		return findClosestValueHelper(tree.Left, target, closest)
		// If our target is bigger than current number, go right
	} else if target > tree.Value {
		return findClosestValueHelper(tree.Right, target, closest)
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

// Here's an example of how to use our tree!
func main() {
	// Let's build a simple tree:
	//       10
	//      /  \
	//     5    15
	//    / \   / \
	//   2   7 13  22

	tree := &BinarySearchTree{
		Value: 10,
		Left: &BinarySearchTree{
			Value: 5,
			Left:  &BinarySearchTree{Value: 2},
			Right: &BinarySearchTree{Value: 7},
		},
		Right: &BinarySearchTree{
			Value: 15,
			Left:  &BinarySearchTree{Value: 13},
			Right: &BinarySearchTree{Value: 22},
		},
	}

	// Let's try to find the closest value to 12
	result := tree.FindClosestValue(12)
	fmt.Printf("The closest value to 12 is: %d\n", result)
}
