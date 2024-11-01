package a2zdsaingolang

type MinStack struct {
	Stack    []int
	MinStack []int
}

func Constructor() MinStack {
	return MinStack{
		Stack:    make([]int, 0),
		MinStack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	// add the value to the main stack
	this.Stack = append(this.Stack, val)

	if len(this.MinStack) == 0 || val < this.MinStack[len(this.MinStack)-1] {
		this.MinStack = append(this.MinStack, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.Stack) == 0 {
		return
	}

	if this.Stack[len(this.Stack)-1] == this.MinStack[len(this.MinStack)-1] {
		this.MinStack = this.MinStack[:len(this.MinStack)-1]
	}

	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinStack[len(this.MinStack)-1]
}
