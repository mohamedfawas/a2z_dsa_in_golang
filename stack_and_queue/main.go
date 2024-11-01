package main

import (
	"fmt"
	queueusingarray "stackandqueue/queue/queueUsingArray"
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

}
