package singlylinkedlist

func (ll *LinkedList) ArrayToLinkedList(arr []int) {
	for _, value := range arr {
		ll.InsertAtTheEnd(value)
	}
}
