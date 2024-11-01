package stackusinglinkedlist

// Node represents each element in the stack
// It contains the data and a pointer to the next node
type Node struct {
	Value int
	next  *Node
}

// Stack represents the stack structure
// It contains a pointer to the top node and the size of the stack
type Stack struct {
	top  *Node
	size int
}

func NewStack() *Stack {
	return &Stack{
		top:  nil,
		size: 0,
	}
}

// Push adds a new element to the top of the stack
// Time Complexity: O(1)
func (s *Stack) Push(value int) {
	newNode := &Node{
		Value: value,
		next:  nil,
	}

	// If stack is not empty, make the new node point to current top
	// Then make the new node the top of the stack
	if s.top != nil {
		newNode.next = s.top
	}

	s.top = newNode
	s.size++
}
