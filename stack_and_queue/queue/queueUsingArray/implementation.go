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
func (q *Queue) Dequeue() (int, error) {
	// check if the queue is empty
	if len(q.items) == 0 {
		return 0, errors.New("queue is empty")
	}

	// Retrieve and remove the first item
	item := q.items[0]
	q.items = q.items[1:]
	fmt.Println("dequeued :", item)
	return item, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) IsFull() bool {
	return len(q.items) == q.size
}

func (q *Queue) Display() {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	fmt.Println("queue : ", q.items)
}
