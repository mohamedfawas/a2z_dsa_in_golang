package probsreview

func reorderList(head *ListNode) {
	// Step 0: Handle edge cases
	// If list is empty or has only 1 or 2 nodes, no need to reorder
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	// Step 1: Find the middle of the linked list using slow and fast pointers
	// Slow moves one step at a time, fast moves two steps
	// When fast reaches end, slow will be at middle
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Step 2: Reverse the second half of the linked list
	// First, separate the second half
	secondHalf := slow.Next
	slow.Next = nil // Break the link between first and second half

	// Reverse the second half
	var prev *ListNode = nil
	curr := secondHalf
	for curr != nil {
		// Store next node
		nextTemp := curr.Next
		// Reverse the link
		curr.Next = prev
		// Move prev and curr one step forward
		prev = curr
		curr = nextTemp
	}
	// prev is now the head of reversed second half

	// Step 3: Merge the first half with reversed second half
	// Take one node from first half, then one from second half
	firstHalf := head
	secondHalf = prev
	for secondHalf != nil {
		// Store next pointers
		firstNext := firstHalf.Next
		secondNext := secondHalf.Next

		// Connect first half node to second half node
		firstHalf.Next = secondHalf
		// Connect second half node to next first half node
		secondHalf.Next = firstNext

		// Move pointers forward
		firstHalf = firstNext
		secondHalf = secondNext
	}
}
