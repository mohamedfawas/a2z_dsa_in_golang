package problems

/*
Problem statement
Given a linked list of 'N' nodes, where each node has an integer value that can be 0, 1, or 2. You need to sort the linked list in non-decreasing order and the return the head of the sorted list.

Example:
Given linked list is 1 -> 0 -> 2 -> 1 -> 2.
The sorted list for the given linked list will be 0 -> 1 -> 1 -> 2 -> 2.
*/

func Sort012(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	zero := &ListNode{Val: -1}
	one := &ListNode{Val: -1}
	two := &ListNode{Val: -1}
	zeroHead := zero
	oneHead := one
	twoHead := two

	temp := head
	for temp != nil {
		if temp.Val == 0 {
			zero.Next = temp
			zero = zero.Next
		}
		if temp.Val == 1 {
			one.Next = temp
			one = one.Next
		}
		if temp.Val == 2 {
			two.Next = temp
			two = two.Next
		}
		temp = temp.Next
	}

	// Ensure the end of each list is properly terminated
	two.Next = nil

	// Connect zero ll pointer to one head if it have elements, then connect one head to twoHead
	if oneHead.Next != nil {
		zero.Next = oneHead.Next
		one.Next = twoHead.Next // earlier we made sure end of "two" is properly terminated
	} else if twoHead.Next != nil {
		zero.Next = twoHead.Next
	}

	return zeroHead.Next
}
