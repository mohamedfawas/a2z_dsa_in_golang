package main

import (
	"errors"
	"fmt"
)

type MinimumStack struct {
	stack    []int
	minStack []int
}

func (mstack *MinimumStack) Push(value int) {
	mstack.stack = append(mstack.stack, value)

	if len(mstack.minStack) == 0 || value < mstack.minStack[len(mstack.minStack)-1] {
		mstack.minStack = append(mstack.minStack, value)
	}
}

func (mstack *MinimumStack) IsEmpty() bool {
	return len(mstack.stack) == 0
}

func (mstack *MinimumStack) Pop() (int, error) {
	if mstack.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	top := mstack.stack[len(mstack.stack)-1]
	topMinStack := mstack.minStack[len(mstack.minStack)-1]

	// remove from main stack
	mstack.stack = mstack.stack[:len(mstack.stack)-1]
	if top == topMinStack {
		mstack.minStack = mstack.minStack[:len(mstack.minStack)-1]
	}

	return top, nil
}

func (mstack *MinimumStack) GetMin() int {
	return mstack.minStack[len(mstack.minStack)-1]
}

func main() {
	minstack := &MinimumStack{}
	fmt.Println("Initial stack")
	fmt.Println(minstack)
	fmt.Println("After adding elements")
	minstack.Push(10)
	minstack.Push(2)
	minstack.Push(3)
	minstack.Push(4)
	fmt.Println("main stack : ", minstack.stack)
	fmt.Println("min stack : ", minstack.minStack)
	fmt.Println("min value in the stack : ", minstack.GetMin())
}
