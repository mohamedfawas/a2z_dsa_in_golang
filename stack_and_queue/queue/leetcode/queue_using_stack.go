package queueleetcode

// MyQueue struct that will hold two stacks
type MyQueue struct {
	stack1 []int // This stack is used for pushing elements
	stack2 []int // This stack is used for popping elements
}

// Constructor initializes a new MyQueue
func Constructor() MyQueue {
	return MyQueue{
		stack1: []int{},
		stack2: []int{},
	}
}

// Push pushes an element to the back of the queue
func (this *MyQueue) Push(x int) {
	// Simply push the element onto stack1
	this.stack1 = append(this.stack1, x)
}

// Pop removes the element from the front of the queue and returns it
func (this *MyQueue) Pop() int {
	// Ensure stack2 has the elements for popping
	this.moveToStack2()
	// Pop the top element from stack2 (the front of the queue)
	if len(this.stack2) == 0 {
		return -1 // return -1 if stack is empty (not expected in this problem)
	}
	top := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1] // Remove the top element
	return top
}

// Peek returns the element at the front of the queue without removing it
func (this *MyQueue) Peek() int {
	// Ensure stack2 has the elements for peeking
	this.moveToStack2()
	// Return the top element from stack2 (the front of the queue)
	if len(this.stack2) == 0 {
		return -1 // return -1 if stack is empty (not expected in this problem)
	}
	return this.stack2[len(this.stack2)-1]
}

// Empty returns true if the queue is empty, false otherwise
func (this *MyQueue) Empty() bool {
	// The queue is empty if both stacks are empty
	return len(this.stack1) == 0 && len(this.stack2) == 0
}

// moveToStack2 moves elements from stack1 to stack2 to maintain queue order
func (this *MyQueue) moveToStack2() {
	// Only move if stack2 is empty
	if len(this.stack2) == 0 {
		// Transfer all elements from stack1 to stack2
		for len(this.stack1) > 0 {
			// Pop from stack1 and push to stack2
			top := this.stack1[len(this.stack1)-1]
			this.stack1 = this.stack1[:len(this.stack1)-1] // Remove top from stack1
			this.stack2 = append(this.stack2, top)         // Push to stack2
		}
	}
}
