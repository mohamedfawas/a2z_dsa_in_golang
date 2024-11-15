package main

import "fmt"

type MinHeap struct {
	elements []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		elements: make([]int, 0),
	}
}

func (h *MinHeap) Insert(value int) {
	h.elements = append(h.elements, value)

	// Bubble up
	h.BubbleUp(len(h.elements) - 1)
}

func (h *MinHeap) BubbleUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2

		if h.elements[parentIndex] <= h.elements[index] {
			break
		}

		h.elements[parentIndex], h.elements[index] = h.elements[index], h.elements[parentIndex]
		index = parentIndex
	}
}

func (h *MinHeap) ExtractMin() (int, error) {
	if len(h.elements) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	min := h.elements[0]

	lastIndex := len(h.elements) - 1
	h.elements[0] = h.elements[lastIndex]

	h.elements = h.elements[:lastIndex]

	if lastIndex > 0 {
		h.SinkDown(0)
	}

	return min, nil

}

func (h *MinHeap) SinkDown(index int) {
	size := len(h.elements)

	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		if leftChild < size && h.elements[leftChild] < h.elements[smallest] {
			smallest = leftChild
		}

		if rightChild < size && h.elements[rightChild] < h.elements[smallest] {
			smallest = rightChild
		}

		if smallest == index {
			break
		}

		h.elements[index], h.elements[smallest] = h.elements[smallest], h.elements[index]
		index = smallest
	}
}

func (h *MinHeap) Peek() (int, error) {
	if len(h.elements) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	return h.elements[0], nil
}

func (h *MinHeap) Size() int {
	return len(h.elements)
}

func main() {
	heap := NewMinHeap()

	numbers := []int{110, 7, 2, 9, 667, 35, 7, 2, 9, 4, 71}
	fmt.Println("Adding numbers : ", numbers)
	for _, num := range numbers {
		heap.Insert(num)
	}

	fmt.Println("Numbers are coming out : ")
	for heap.Size() > 0 {
		num, _ := heap.ExtractMin()
		fmt.Print(num, " ")
	}
	fmt.Println()
}
