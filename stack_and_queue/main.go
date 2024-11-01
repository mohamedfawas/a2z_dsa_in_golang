package main

import (
	"fmt"
	queueusingarray "stackandqueue/queue/queueUsingArray"
	stackusingarray "stackandqueue/stack/stackUsingArray"
	stackusinglinkedlist "stackandqueue/stack/stackUsingLinkedList"
)

func main() {

	stack := &stackusingarray.Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Size of stack:", stack.Size())

	if peek, err := stack.Peek(); err == nil {
		fmt.Println("the top element is : ", peek)
	}

	if pop, err := stack.Pop(); err == nil {
		fmt.Println("the popped element is : ", pop)
	}

	fmt.Println("Size of stack after pop:", stack.Size()) // Output: 2
	isEmpty := stack.IsEmpty()
	fmt.Println("Is stack empty?", isEmpty)

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
