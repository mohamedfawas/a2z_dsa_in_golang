package slprobset

type Node struct {
	data int
	next *Node
}

// takes head of linked list as input and returns the new head after reversal
func ReverseLinkedList(head *Node) *Node {
	var prev *Node // holds the pointer to the previous node
	current := head

	for current != nil {
		next := current.next // temporarily stores the next node
		current.next = prev
		prev = current
		current = next
	}

	return prev
}
