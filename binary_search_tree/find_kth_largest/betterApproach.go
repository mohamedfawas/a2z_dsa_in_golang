package findkthlargestbst

// Function to find the kth largest element in the BST
func findKthLargest(node *Node, k *int, result *int) {
	if node == nil || *k <= 0 {
		// If node is nil or we've already found the kth largest, stop
		return
	}

	// Step 1: Go to the right subtree first (largest elements)
	findKthLargest(node.Right, k, result)

	// Step 2: Process the current node
	*k-- // Decrease k because we visited one node
	if *k == 0 {
		*result = node.Value // We've found the kth largest
		return
	}

	// Step 3: Go to the left subtree (smaller elements)
	findKthLargest(node.Left, k, result)
}
