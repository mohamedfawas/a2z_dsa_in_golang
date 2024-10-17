package main

import (
	"fmt"
	singlylinkedlist "linked_list/singly_linked_list"
)

func main() {

	// Convert array to a linked list
	fmt.Println("solution to convert array to linkedlist")
	ll := singlylinkedlist.LinkedList{}
	arr := []int{10, 20, 30, 40, 50}
	ll.ArrayToLinkedList(arr)
	fmt.Println("Linked list created by converting the array : ")
	fmt.Println(ll.Traverse())

	// Add a node at the end & beginning
	fmt.Println("After adding the node at the beginning : ")
	ll.InsertAtTheBeginning(33)
	fmt.Println(ll.Traverse())
	fmt.Println("After adding the node at the end : ")
	ll.InsertAtTheEnd(11)
	fmt.Println(ll.Traverse())

	// Delete the node with the value specified
	fmt.Println("Deleting the node with the given value")
	ll.DeleteTheNodeWithTheValue(20)
	fmt.Println(ll.Traverse())

	// reverse the linkedlist

}
