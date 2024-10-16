package singlylinkedlist

// Search for data in ll
func (ll *LinkedList) Search(target int) bool {
	current := ll.head
	for current != nil {
		if current.data == target {
			return true
		}
		current = current.next
	}
	return false
}
