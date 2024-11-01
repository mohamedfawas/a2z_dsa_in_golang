package main

import (
	"fmt"
	stackusingarray "stackandqueue/stack/stackUsingArray"
)

func main() {
	// create a stack with defined size
	stack := stackusingarray.NewStack(5)

	// example usage with print stack
	fmt.Println("pushing elements : 10,20,30,40")
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)

	// print initial stack
	stack.PrintStack()

	// Pop an element
	fmt.Println("\nPopping one element")
	if element, err := stack.Pop(); err == nil {
		fmt.Printf("Popped : %d \n", element)
	}

	// Print stack after popping
	stack.PrintStack()

	// Push another element
	fmt.Println("\nPushing element: 50")
	stack.Push(50)

	// Print final stack state
	stack.PrintStack()
}
