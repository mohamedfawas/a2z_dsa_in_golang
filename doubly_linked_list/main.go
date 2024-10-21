package main

import "doublell/implementation"

func main() {
	arr := []int{1, 2, 2, 34, 5, 6, 3, 2, 4, 6}
	head := implementation.ConvertArrayToDoublyLinkedList(arr)

	implementation.DisplayList(head)
}
