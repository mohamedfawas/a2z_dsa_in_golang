package singlylinkedlist

import "fmt"

// Delete from the Beginning
func (ll *LinkedList) DeleteFromBeginning() error {
	// Check if the linked list is empty (i.e., no nodes)
	if ll.head == nil {
		// If the list is empty, return an error indicating that there's nothing to delete
		return fmt.Errorf("list is empty")
	}

	// If the list is not empty, move the head pointer to the next node.
	// This effectively removes the first node from the list
	ll.head = ll.head.next
	ll.size--
	return nil
}

// DeleteFromEnd removes the last node from the linked list
func (ll *LinkedList) DeleteFromEnd() error {

	//  Check if the linked list is empty by checking if the head is nil
	if ll.head == nil {
		return fmt.Errorf("list is empty")
	}

	// Check if the linked list has only one node
	//  If the head node's next is nil, it means there's only one node in the list
	// To delete it, set the head to nil
	if ll.head.next == nil {
		ll.head = nil
	} else {

		// If the list has more than one node, traverse the list to find the second-to-last node
		// Create a variable `current` and start from the head
		current := ll.head

		// Loop through the list until `current.next.next` is nil
		// This means `current` will stop at the second-to-last node, as the next node is the last one
		for current.next.next != nil {
			current = current.next
		}

		// Set `current.next` to nil, effectively removing the last node from the list
		current.next = nil
	}
	ll.size--
	return nil
}

// Delete at a Specific Position
// head -> 1 -> 2 -> 3 -> 4 -> 5 -> nil
func (ll *LinkedList) DeleteAtPosition(position int) error {
	if position < 0 || position >= ll.size {
		return fmt.Errorf("position out of bounds")
	}

	if position == 0 {
		ll.head = ll.head.next
	} else {
		current := ll.head
		for i := 0; i < position-1; i++ {
			current = current.next
		}

		current.next = current.next.next
	}

	ll.size--
	return nil
}

func (ll *LinkedList) DeleteTheNodeWithTheValue(value int) error {
	/*
		check if list is empty
		- if the head node have the value
		- traverse the linkedlist until the first occurence of the given value
	*/
	if ll.head == nil {
		return fmt.Errorf("list is empty")
	}

	if ll.head.data == value {
		ll.head = ll.head.next
		ll.size--
		return nil
	}

	// Traverse the list to find the node with the given value
	current := ll.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}
	if current == nil {
		return fmt.Errorf("node with the given value not found in the linked list")
	}

	// Adjust the link to remove the node
	current.next = current.next.next
	ll.size--
	return nil
}
