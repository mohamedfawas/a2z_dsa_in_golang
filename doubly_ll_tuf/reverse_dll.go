package doublylltuf

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

// ReverseDoublyLinkedList reverses the doubly linked list and returns the new head
func ReverseDoublyLinkedList(head *Node) *Node {
	// If the list is empty (head is nil), return nil as there's nothing to reverse
	if head == nil {
		return nil
	}

	var temp *Node  // Temporary variable to hold the previous node during swapping
	current := head // Start with the head of the list

	// Traverse the list and swap the prev and next pointers for each node
	for current != nil {
		temp = current.Prev         // Save the current node's previous pointer to temp
		current.Prev = current.Next // Set current node's prev pointer to its next pointer (swap)
		current.Next = temp         // Set current node's next pointer to what was originally prev (swap)
		current = current.Prev      // Move to the next node, which is now the previous node due to swapping
	}

	// After the loop, temp will point to the second-to-last node (previous head). Adjust the new head.
	if temp != nil {
		head = temp.Prev // temp.prev is the new head of the reversed list
	}

	return head // Return the new head of the reversed list
}
