package stackusinglinkedlist

import (
	"fmt"
)

// Node represents each element in the stack
// It contains the data and a pointer to the next node
type Node struct {
	Value int
	next  *Node
}

// Stack represents the stack structure
// It contains a pointer to the top node and the size of the stack
type Stack struct {
	top  *Node
	size int
}

func NewStack() *Stack {
	return &Stack{
		top:  nil,
		size: 0,
	}
}

// Push adds a new element to the top of the stack
// Time Complexity: O(1)
func (s *Stack) Push(value int) {
	newNode := &Node{
		Value: value,
		next:  nil,
	}

	// If stack is not empty, make the new node point to current top
	// Then make the new node the top of the stack
	if s.top != nil {
		newNode.next = s.top
	}

	s.top = newNode
	s.size++
}

// Pop removes and returns the top element from the stack
// Returns -1 if the stack is empty
// Time Complexity: O(1)
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return -1
	}

	// store the value to return
	value := s.top.Value

	// move top to the next node
	s.top = s.top.next
	s.size--
	return value
}

// Peek returns the top element without removing it
// Returns -1 if the stack is empty
// Time Complexity: O(1)
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return -1
	}

	return s.top.Value
}

func (s *Stack) PrintStack() {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return
	}

	current := s.top
	fmt.Println("stack (top to bottom) : ")
	for current != nil {
		fmt.Printf("%d", current.Value)
		current = current.next
	}
	fmt.Println()
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}
