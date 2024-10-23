package problems

/*
Given the head of a linked list, remove the nth node from the end of the list and return its head
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Start by assigning the 'fast' pointer to the head of the list.
	fast := head

	// Move the 'fast' pointer 'n' steps forward in the list.
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// If 'fast' is now nil, it means that 'n' is the size of the list.
	// So, we need to remove the head node, hence we return the next node.
	if fast == nil {
		return head.Next
	}

	// Assign the 'slow' pointer to the head of the list.
	slow := head

	// Move both 'slow' and 'fast' pointers forward until 'fast' reaches the last node.
	// The 'slow' pointer will then be at the node before the node to be removed.
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// Skip the next node (which is the nth node from the end) by adjusting the 'Next' pointer of 'slow'.
	slow.Next = slow.Next.Next

	// Return the modified list's head.
	return head
}
