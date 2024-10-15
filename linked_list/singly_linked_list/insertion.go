package singlylinkedlist

import "fmt"

type Node struct {
	data int   // stores the int value
	next *Node // denotes next element in the linked list
}

type LinkedList struct {
	head *Node // Points to the first node in the list
	size int   // Keeps track of the number of nodes in the list.
}

// Insert at the beginning
func (ll *LinkedList) InsertAtTheBeginning(data int) {
	newNode := &Node{data: data, next: ll.head}
	ll.head = newNode
	ll.size++
}

// Insert at the end
func (ll *LinkedList) InsertAtTheEnd(data int) {
	newNode := &Node{data: data}

	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head

		for current.next != nil {
			current = current.next
		}

		current.next = newNode
	}

	ll.size++
}

// Insert at a position
func (ll *LinkedList) InsertAtPosition(data int, position int) error {
	// Check if the position is valid. It must be between 0 and the size of the list.
	if position < 0 || position > ll.size {
		return fmt.Errorf("position out of bounds")
	}

	// Create a new node with the given data.
	// The 'next' field is set to nil initially
	newNode := &Node{data: data}

	// If the position is 0 (insertion at the beginning of the list)
	if position == 0 {
		// Point the new node's next to the current head of the list
		newNode.next = ll.head
		// Update the head of the list to the new node
		ll.head = newNode
	} else {
		// For insertion at any other position (not the beginning)
		// Start traversing the list from the head
		current := ll.head

		// Move to the node just before the target position by iterating 'position - 1' times
		for i := 0; i < position-1; i++ {
			current = current.next // Advance to the next node
		}

		// Set the new node's next to point to the node currently at the target position
		newNode.next = current.next
		// Update the current node's next to point to the new node, effectively inserting it
		current.next = newNode
	}
	ll.size++
	return nil
}
