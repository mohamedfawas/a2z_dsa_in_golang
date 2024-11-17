package maxheap

// MaxHeap represents our special tree where bigger numbers are parents
type MaxHeap struct {
	// heap is like a tree stored in a list, where for any number at position i:
	// - its left child is at position 2*i + 1
	// - its right child is at position 2*i + 2
	// - its parent is at position (i-1)/2
	heap []int
}

// NewMaxHeap creates a new empty heap
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		heap: []int{},
	}
}

// BuildHeap creates a max heap from a list of numbers
func (h *MaxHeap) BuildHeap(arr []int) {
	// First, copy all numbers to our heap
	h.heap = arr

	// Start from the last parent node (parent of last element)
	// and bubble down each parent to its correct position
	for i := len(h.heap)/2 - 1; i >= 0; i-- {
		h.bubbleDown(i)
	}
}

// Insert adds a new number to our heap
func (h *MaxHeap) Insert(val int) {
	// Add the new number at the end of our list
	h.heap = append(h.heap, val)

	// Let the new number bubble up to its correct position
	h.bubbleUp(len(h.heap) - 1)
}

// Delete removes and returns the biggest number (root) from our heap
func (h *MaxHeap) Delete() (int, bool) {
	// If heap is empty, we can't delete anything
	if len(h.heap) == 0 {
		return 0, false
	}

	// Save the biggest number (it's at the root/first position)
	maxVal := h.heap[0]

	// Move the last number to the root position
	lastIndex := len(h.heap) - 1
	h.heap[0] = h.heap[lastIndex]

	// Remove the last position
	h.heap = h.heap[:lastIndex]

	// If heap isn't empty, bubble down the new root to its correct position
	if len(h.heap) > 0 {
		h.bubbleDown(0)
	}

	return maxVal, true
}

// bubbleUp moves a number up the tree until it's in the right spot
func (h *MaxHeap) bubbleUp(index int) {
	// While we're not at the root and our parent is smaller than us
	for index > 0 {
		parentIndex := (index - 1) / 2

		// If parent is bigger or equal, we're in the right spot
		if h.heap[parentIndex] >= h.heap[index] {
			break
		}

		// Swap with parent and move up
		h.heap[index], h.heap[parentIndex] = h.heap[parentIndex], h.heap[index]
		index = parentIndex
	}
}

// bubbleDown moves a number down the tree until it's in the right spot
func (h *MaxHeap) bubbleDown(index int) {
	for {
		largest := index
		leftChild := 2*index + 1
		rightChild := 2*index + 2

		// Check if left child exists and is larger than current largest
		if leftChild < len(h.heap) && h.heap[leftChild] > h.heap[largest] {
			largest = leftChild
		}

		// Check if right child exists and is larger than current largest
		if rightChild < len(h.heap) && h.heap[rightChild] > h.heap[largest] {
			largest = rightChild
		}

		// If we're already the largest, we're done
		if largest == index {
			break
		}

		// Swap with the largest child and continue down that path
		h.heap[index], h.heap[largest] = h.heap[largest], h.heap[index]
		index = largest
	}
}
