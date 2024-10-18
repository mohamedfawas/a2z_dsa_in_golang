package problems

func reverseList(head *ListNode) *ListNode {
	current := head
	var prev *ListNode

	for current != nil {
		front := current.Next
		current.Next = prev
		prev = current
		current = front
	}

	return prev
}
