package stackusingdoublequeue

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

// Implement stack using double queue
type Stack struct {
	q1 *Queue
	q2 *Queue
}

func NewStack() *Stack {
	return &Stack{
		q1: &Queue{},
		q2: &Queue{},
	}
}

func (s *Stack) Push(x int) {
	s.q2.Enqueue(x)

	for !s.q1.IsEmpty() {
		front, _ := s.q1.Dequeue()
		s.q2.Enqueue(front)
	}

	// swap q1 and q2
	s.q1, s.q2 = s.q2, s.q1
}

func (s *Stack) Pop() (int, error) {
	if s.q1.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.q1.Dequeue()
}

func (s *Stack) Peek() (int, error) {
	if s.q1.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.q1.Peek()
}

func (s *Stack) Size() int {
	return s.q1.Size()
}
