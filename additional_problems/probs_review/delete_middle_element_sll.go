package probsreview

type Node struct {
	data int
	next *Node
}

func deleteMiddle(head *Node) *Node {
	// If list is empty or has only one node
	if head == nil || head.next == nil {
		return nil
	}

	// If list has only two nodes, delete the second node
	if head.next.next == nil {
		head.next = nil
		return head
	}

	// Initialize two pointers and a previous pointer
	slow := head
	fast := head
	var prevSlow *Node

	// Move pointers until fast reaches the end
	// When fast reaches end, slow will be at middle
	for fast != nil && fast.next != nil {
		fast = fast.next.next // Move fast two steps
		prevSlow = slow       // Store previous position of slow
		slow = slow.next      // Move slow one step
	}

	// Delete middle node by updating the next pointer
	// of previous node to skip the middle node
	prevSlow.next = slow.next

	return head
}
