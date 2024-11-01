package stackusingarray

import "fmt"

// Stack represents a stack data structure using a fixed-size array
type Stack struct {
	items []int // Array to store stack elements
	top   int   // Index of the top element in the stack
	size  int   // Maximum size of the stack
}

// NewStack creates and initializes a new stack with the given size
func NewStack(size int) *Stack {
	return &Stack{
		items: make([]int, size), // Create an array with the specified size
		top:   -1,                // Initialize top as -1 to indicate empty stack
		size:  size,              // Store the maximum size of the stack
	}
}

func (s *Stack) Push(element int) error {
	if s.IsFull() {
		return fmt.Errorf("stack overflow : cannot push element %d", element)
	}

	// Increment the top and add the element
	s.top++                  // Move top index up by 1
	s.items[s.top] = element // Place the new element at top position
	return nil
}

func (s *Stack) Pop() (int, error) {
	// Check if stack is empty
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack underflow: cannot pop from empty stack")
	}

	// Get the top element and decrement top
	element := s.items[s.top]
	s.top--
	return element, nil
}

func (s *Stack) Peek() (int, error) {
	// check if empty
	if s.IsEmpty() {
		return 0, fmt.Errorf("cannot peek: stack is empty")
	}

	return s.items[s.top], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

// IsFull checks if the stack is full
func (s *Stack) IsFull() bool {
	return s.top == s.size-1
}

func (s *Stack) Size() int {
	return s.top + 1
}

func (s *Stack) PrintStack() {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return
	}

	fmt.Println("\nStack contents (top to bottom):")
	fmt.Println("--------------------------------")

	for i := s.top; i >= 0; i-- {
		if i == s.top {
			fmt.Printf("%d <- TOP \n", s.items[i])
		} else {
			fmt.Printf("%d\n", s.items[i])
		}
	}
	fmt.Println("--------------------------------")
	fmt.Printf("Size: %d, Capacity: %d\n", s.Size(), s.size)
}
