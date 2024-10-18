package practice

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumsInLL(l1, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{Val: 0}
	current := dummyNode
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		// calculate sum and carry, add digit
		sum := val1 + val2 + carry
		carry = sum / 10
		digit := sum % 10

		// create new node with the digit
		newNode := &ListNode{Val: digit}
		current.Next = newNode
		current = current.Next
	}

	return dummyNode.Next
}
