package queueleetcode

// MyStack is the struct that represents our stack, which internally uses two queues.
type MyStack struct {
	queue1 []int // Primary queue used to store elements
	queue2 []int // Temporary queue used for push operation
}

// Constructor initializes and returns an empty stack.
func Constructor() MyStack {
	return MyStack{
		queue1: []int{},
		queue2: []int{},
	}
}

// Push adds an element x to the top of the stack.
func (this *MyStack) Push(x int) {
	// Step 1: Add the new element to queue2 (temporary queue)
	this.queue2 = append(this.queue2, x)

	// Step 2: Move all elements from queue1 to queue2, making x the "top" of the stack
	for len(this.queue1) > 0 {
		this.queue2 = append(this.queue2, this.queue1[0]) // Move front of queue1 to queue2
		this.queue1 = this.queue1[1:]                     // Remove the front element from queue1
	}

	// Step 3: Swap the names of queue1 and queue2
	// Now, queue1 has all elements in correct stack order and queue2 is empty
	this.queue1, this.queue2 = this.queue2, this.queue1
}

// Pop removes the top element from the stack and returns it.
func (this *MyStack) Pop() int {
	// Pop operation is simple because the top element is at the front of queue1
	topElement := this.queue1[0]  // Get the front element of queue1 (top of the stack)
	this.queue1 = this.queue1[1:] // Remove the front element from queue1
	return topElement
}

// Top returns the top element of the stack without removing it.
func (this *MyStack) Top() int {
	return this.queue1[0] // The front element of queue1 is the top of the stack
}

// Empty returns true if the stack is empty, otherwise false.
func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0 // Stack is empty if queue1 has no elements
}
