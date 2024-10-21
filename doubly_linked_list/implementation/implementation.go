package implementation

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

func InsertAtEnd(head *Node, data int) *Node {
	// Create a new node
	newNode := &Node{data: data}

	// If the list is empty, the new node becomes the head
	if head == nil {
		return newNode
	}

	// traverse to the last node of the list
	temp := head
	for temp.next != nil {
		temp = temp.next
	}

	// Update pointers to insert the new node at the end
	temp.next = newNode
	newNode.prev = temp

	return head
}

func InsertAtBeginning(data int, head *Node) *Node {
	newNode := &Node{data: data}

	// if list is empty
	if head == nil {
		return newNode
	}

	// update pointers of the head
	temp := head

	temp.prev = newNode
	newNode.next = temp

	return newNode
}

func InsertAfterNode(node *Node, data int) {
	newNode := &Node{data: data}

	// if the list is empty
	if node == nil {
		node = newNode
		return
	}

	// pointer
	temp := node

	// temp node might be connected to other nodes, so we need to make changes in that also
	temp.next = newNode.next
	temp.next = newNode
	newNode.prev = temp

	if newNode.next != nil {
		newNode.next.prev = newNode
	}
}

func InsertBeforeNode(node *Node, data int) {
	newNode := &Node{data: data}

	// if the list is empty
	if node == nil {
		node = newNode
		return
	}

	temp := node

	newNode.prev = temp.prev
	temp.prev = newNode
	newNode.next = temp

	if newNode.prev != nil {
		newNode.prev.next = newNode
	}
}

// DisplayList prints the elements of the list from head to tail
func DisplayList(head *Node) {
	temp := head
	for temp != nil {
		fmt.Printf("%d ", temp.data)
		temp = temp.next // Move to the next node
	}
	fmt.Println()
}

// ConvertArrayToDoublyLinkedList takes an array and converts it into a doubly linked list
func ConvertArrayToDoublyLinkedList(arr []int) *Node {
	// Check if the array is empty
	if len(arr) == 0 {
		return nil
	}

	// Create the head node from the first element in the array
	head := &Node{data: arr[0]}
	current := head

	// Iterate over the remaining elements in the array
	for i := 1; i < len(arr); i++ {
		// Create a new node for each element
		newNode := &Node{data: arr[i]}

		// Set the current node's next to the new node
		current.next = newNode

		// Set the new node's prev to the current node
		newNode.prev = current

		// Move the current pointer to the new node
		current = newNode
	}

	return head // Return the head of the doubly linked list
}
