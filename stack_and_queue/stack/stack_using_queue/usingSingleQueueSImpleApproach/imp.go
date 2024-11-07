package usingsinglequeuesimpleapproach

import "errors"

type Stack struct {
	queue []int
}

func NewStack() *Stack {
	return &Stack{
		queue: make([]int, 0),
	}
}

func (s *Stack) Push(val int) {
	s.queue = append(s.queue, val)

	for i := 0; i < len(s.queue)-1; i++ {
		front := s.queue[0]
		s.queue = s.queue[1:]
		s.queue = append(s.queue, front)
	}
}

func (s *Stack) Pop() (int, error) {
	if len(s.queue) == 0 {
		return 0, errors.New("stack is empty")
	}

	front := s.queue[0]
	s.queue = s.queue[1:]
	return front, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.queue) == 0 {
		return 0, errors.New("stack is empty")
	}

	front := s.queue[0]
	return front, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.queue) == 0
}

func (s *Stack) Size() int {
	return len(s.queue)
}
