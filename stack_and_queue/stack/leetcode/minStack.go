package stackleetcode

// MinStack structure contains two slices:
// - stack: stores all the values
// - minStack: keeps track of minimums at each step
type MinStack struct {
	stack    []int // Main stack to store all elements
	minStack []int // Auxiliary stack to keep track of minimums
}

// Constructor initializes a new MinStack
// Returns an empty MinStack instance
func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0), // Initialize empty main stack
		minStack: make([]int, 0), // Initialize empty minimum stack
	}
}

// Push adds a new value to the stack and updates minimum if necessary
// Time Complexity: O(1)
func (this *MinStack) Push(val int) {
	// always push the new value to the main stack
	this.stack = append(this.stack, val)

	// For minStack:
	// If it's empty OR new value is smaller than current minimum
	// Push the new value as the new minimum
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

// Pop removes the element on top of the stack
// Time Complexity: O(1)
func (this *MinStack) Pop() {
	// If stack is empty, return (though problem guarantees this won't happen)
	if len(this.stack) == 0 {
		return
	}

	// If the value being popped is the current minimum,
	// we need to remove it from minStack too
	if this.stack[len(this.stack)-1] == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}

	// remove the top element from the main stack
	this.stack = this.stack[:len(this.stack)-1]
}

// Top returns the top element of the stack
// Time Complexity: O(1)
func (this *MinStack) Top() int {
	// Return the last element in the stack
	// Problem guarantees stack won't be empty when called
	return this.stack[len(this.stack)-1]
}

// GetMin returns the minimum element in the stack
// Time Complexity: O(1)
func (this *MinStack) GetMin() int {
	// Return the top of minStack which maintains the current minimum
	// Problem guarantees stack won't be empty when called
	return this.minStack[len(this.minStack)-1]
}
