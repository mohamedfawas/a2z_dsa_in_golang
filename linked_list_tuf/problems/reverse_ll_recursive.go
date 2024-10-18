package problems

func reverseListRec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListRec(head.Next)
	front := head.Next
	front.Next = head
	head.Next = nil

	return newHead
}
