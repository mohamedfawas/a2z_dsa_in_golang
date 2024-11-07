package main

import (
	"fmt"
	queueusingstack "stackandqueue/queue/queueUsingStack"
)

func main() {
	// Implement stack using two queues
	q := queueusingstack.NewQeueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	front, _ := q.Dequeue()
	fmt.Println("popped element : ", front)
	front, _ = q.Peek()
	fmt.Println("front element : ", front)
	front, _ = q.Dequeue()
	fmt.Println("popped element : ", front)
}
