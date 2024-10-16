package singlylinkedlist

func (ll *LinkedList) Traverse() []int {
	var elements []int
	current := ll.head
	for current != nil {
		elements = append(elements, current.data)
		current = current.next
	}
	return elements
}
