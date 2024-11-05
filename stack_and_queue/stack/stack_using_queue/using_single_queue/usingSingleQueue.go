package stackusingsinglequeue

import "errors"

type Queue struct {
	items []int
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Enqueue(element int) {
	q.items = append(q.items, element)
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	front := q.items[0]
	q.items = q.items[1:]

	return front, nil
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	front := q.items[0]
	return front, nil

}

func (q *Queue) Size() int {
	return len(q.items)
}

// Implement stack using single queue
type Stack struct {
	queue *Queue
}

func NewStack() *Stack {
	return &Stack{queue: &Queue{}}
}

func (s *Stack) Push(x int) {
	s.queue.Enqueue(x)

	for i := 0; i < s.queue.Size()-1; i++ {
		front, _ := s.queue.Dequeue()
		s.queue.Enqueue(front)
	}
}

func (s *Stack) Pop() (int, error) {
	if s.queue.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.queue.Dequeue()
}

func (s *Stack) Top() (int, error) {
	if s.queue.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.queue.Peek()
}

func (s *Stack) IsEmpty() bool {
	return s.queue.IsEmpty()
}
