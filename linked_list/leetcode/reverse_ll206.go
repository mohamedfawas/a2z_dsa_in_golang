package linkedlleetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseListIterative(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
