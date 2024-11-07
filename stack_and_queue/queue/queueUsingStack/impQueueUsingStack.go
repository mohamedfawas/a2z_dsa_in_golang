package queueusingstack

import "errors"

type Queue struct {
	stackInput  []int
	stackOutput []int
}

func NewQeueue() *Queue {
	return &Queue{
		stackInput:  make([]int, 0),
		stackOutput: make([]int, 0),
	}
}

func (q *Queue) Enqueue(value int) {
	q.stackInput = append(q.stackInput, value)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.stackOutput) == 0 {
		for len(q.stackInput) > 0 {
			top := q.stackInput[len(q.stackInput)-1]
			q.stackInput = q.stackInput[:len(q.stackInput)-1]
			q.stackOutput = append(q.stackOutput, top)
		}
	}

	if len(q.stackOutput) == 0 {
		return 0, errors.New("queue is empty")
	}

	front := q.stackOutput[len(q.stackOutput)-1]
	q.stackOutput = q.stackOutput[:len(q.stackOutput)-1]

	return front, nil
}

func (q *Queue) Peek() (int, error) {
	if len(q.stackOutput) == 0 {
		for len(q.stackInput) > 0 {
			top := q.stackInput[len(q.stackInput)-1]
			q.stackInput = q.stackInput[:len(q.stackInput)-1]
			q.stackOutput = append(q.stackOutput, top)
		}
	}

	if len(q.stackOutput) == 0 {
		return 0, errors.New("queue is empty")
	}

	front := q.stackOutput[len(q.stackOutput)-1]
	return front, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.stackInput) == 0 && len(q.stackOutput) == 0
}
