package probsreview

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// define two pointers
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow := slow.Next
		fast := fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
