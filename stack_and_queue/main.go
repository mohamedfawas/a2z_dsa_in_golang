package main

import (
	"fmt"
	queueusingarray "stackandqueue/queue/queueUsingArray"
	stackusingarray "stackandqueue/stack/stackUsingArray"
	stackusinglinkedlist "stackandqueue/stack/stackUsingLinkedList"
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

	fmt.Println("========================queue implementaion is shown below =============")

	// Initialize a queue with a maximum size of 5
	queue := queueusingarray.NewQueue(5)

	// Enqueue items to the queue
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(40)
	queue.Enqueue(50)

	// Display queue contents
	queue.Display()

	// Dequeue items from the queue
	queue.Dequeue()
	queue.Dequeue()

	// Display queue contents again
	queue.Display()

	// Try to enqueue one more item to a full queue
	err := queue.Enqueue(60)
	if err != nil {
		fmt.Println("Error:", err)
	}

	queue.Display()

	fmt.Println("========================stack implementaion using linked list is shown below =============")
	// Create a new stack
	stackll := stackusinglinkedlist.NewStack()

	// Example usage of stack operations
	fmt.Println("Pushing elements: 10, 20, 30")
	stackll.Push(10)
	stackll.Push(20)
	stackll.Push(30)

	// Print the stack
	stackll.PrintStack()

	// Peek at top element
	fmt.Printf("Top element: %d\n", stackll.Peek())

	// Pop an element
	fmt.Printf("Popped element: %d\n", stackll.Pop())

	// Print the stack again
	stackll.PrintStack()
}
