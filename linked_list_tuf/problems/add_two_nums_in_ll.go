package problems

/*
numbers are stored in ll in reverse order
store the result in reverse order

solving approach :
- initialize a dummy node
- While loop when t1,t2 and carry have values
- add values from t1,t2,carry and add it to new node
- return what comes after dummy
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbersInLL(t1, t2 *ListNode) *ListNode {
	dummyNode := &ListNode{Val: -1}
	current := dummyNode
	carry := 0

	for t1 != nil || t2 != nil || carry > 0 {
		// Get values from the lists or 0 if list is exhausted
		val1 := 0
		if t1 != nil {
			val1 = t1.Val
			t1 = t1.Next
		}

		val2 := 0
		if t2 != nil {
			val2 = t2.Val
			t2 = t2.Next
		}

		// calculate sum and carry
		sum := val1 + val2
		carry = sum / 10
		digit := sum % 10

		// Create new node with the calculated digit
		current.Next = &ListNode{Val: digit}
		current = current.Next
	}

	return dummyNode.Next
}
