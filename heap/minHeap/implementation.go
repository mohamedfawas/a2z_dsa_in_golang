package minheap

import "fmt"

// MinHeap represents our special family tree (heap)
// It's just a slice (like a growing list) of numbers
type MinHeap struct {
	elements []int
}

// NewMinHeap creates a new, empty Min Heap
// Think of it like getting a new empty box to store your numbers
func NewMinHeap() *MinHeap {
	return &MinHeap{
		elements: make([]int, 0),
	}
}

// Insert adds a new number to our heap
// Imagine dropping a new marble (number) into our box and making sure it rolls to the right place!
func (h *MinHeap) Insert(value int) {
	// First, we put the new number at the end of our list
	h.elements = append(h.elements, value)

	// Then, we help it "bubble up" to its correct position
	h.bubbleUp(len(h.elements) - 1)
}

// bubbleUp helps a number float up to its right position
// Imagine a balloon floating up until it finds where it belongs!
func (h *MinHeap) bubbleUp(index int) {
	// We keep going until we reach the top (index 0)
	// or until we find the right spot
	for index > 0 {
		// To find the parent's position, we use this special math:
		// parentIndex = (childIndex - 1) / 2
		parentIndex := (index - 1) / 2

		// If the parent is already smaller, we're done!
		// The balloon has found its spot
		if h.elements[parentIndex] <= h.elements[index] {
			break
		}

		// If not, we swap the parent and child
		// Like swapping two balloons' positions
		h.elements[index], h.elements[parentIndex] = h.elements[parentIndex], h.elements[index]

		// Move up to the parent's position and check again
		index = parentIndex
	}
}

// ExtractMin removes and returns the smallest number from our heap
// Like picking up the lightest marble from the top of our pile
func (h *MinHeap) ExtractMin() (int, error) {
	// If the heap is empty, we can't take anything out!
	if len(h.elements) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	// Remember the smallest number (it's at the top!)
	min := h.elements[0]

	// Take the last number and put it at the top
	lastIndex := len(h.elements) - 1
	h.elements[0] = h.elements[lastIndex]

	// Remove the last element (we just moved it to the top)
	h.elements = h.elements[:lastIndex]

	// If we still have elements, we need to help the new top number
	// find its right place by sinking down
	if lastIndex > 0 {
		h.sinkDown(0)
	}

	return min, nil
}

// sinkDown helps a number sink to its correct position
// Imagine a heavy marble sinking through water until it finds its spot!
func (h *MinHeap) sinkDown(index int) {
	size := len(h.elements)

	for {
		// Find the positions of the left and right children
		// leftChild = 2 * parentIndex + 1
		// rightChild = 2 * parentIndex + 2
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		// Check if the left child is smaller than the current smallest
		if leftChild < size && h.elements[leftChild] < h.elements[smallest] {
			smallest = leftChild
		}

		// Check if the right child is smaller than the current smallest
		if rightChild < size && h.elements[rightChild] < h.elements[smallest] {
			smallest = rightChild
		}

		// If we haven't found a smaller child, we're done!
		// The marble has found its resting place
		if smallest == index {
			break
		}

		// Swap with the smallest child and continue sinking
		h.elements[index], h.elements[smallest] = h.elements[smallest], h.elements[index]
		index = smallest
	}
}

// Peek lets us look at the smallest number without removing it
// Like looking at the top marble without picking it up
func (h *MinHeap) Peek() (int, error) {
	if len(h.elements) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	return h.elements[0], nil
}

// Size tells us how many numbers are in our heap
// Like counting how many marbles are in our box
func (h *MinHeap) Size() int {
	return len(h.elements)
}

// Here's an example of how to use our Min Heap!
func main() {
	// Create a new Min Heap
	heap := NewMinHeap()

	// Add some numbers
	numbers := []int{5, 3, 7, 1, 4, 6, 2}
	fmt.Println("Adding numbers:", numbers)
	for _, num := range numbers {
		heap.Insert(num)
	}

	// Let's remove all numbers one by one
	// They will come out in sorted order (smallest to largest)!
	fmt.Print("Numbers coming out: ")
	for heap.Size() > 0 {
		num, _ := heap.ExtractMin()
		fmt.Print(num, " ")
	}
	fmt.Println()
}
