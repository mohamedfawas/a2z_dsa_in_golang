package stackusingarray

import (
	"errors"
)

// Stack represents a stack that holds a slice of integers.
type Stack struct {
	items []int
}

func (s *Stack) Push(element int) {
	s.items = append(s.items, element)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	n := len(s.items)
	top := s.items[n-1]
	s.items = s.items[:n-1]
	return top, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack) Size() int {
	return len(s.items)
}
