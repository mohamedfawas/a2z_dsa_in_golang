package queueusingarray

import (
	"errors"
	"fmt"
)

type Queue struct {
	items []int
	size  int
}

func NewQueue(size int) *Queue {
	return &Queue{
		items: make([]int, 0, size),
		size:  size,
	}
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(item int) error {
	// check if the queue is full
	if len(q.items) >= q.size {
		return errors.New("queue is full") // error if trying to add to a full queue
	}

	// add the item to the end of the queue
	q.items = append(q.items, item)
	fmt.Println("Enqueued : ", item)
	return nil
}

// Dequeue removes an element from the front of the queue.
func Dequeue() {

}
