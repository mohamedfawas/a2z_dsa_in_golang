package main

type MaxHeap struct {
	elements []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		elements: make([]int, 0),
	}
}

func (h *MaxHeap) Insert(val int) {
	h.elements = append(h.elements, val)
	h.BubbleUp(len(h.elements) - 1)
}

func (h *MaxHeap) Delete(val int) (int, bool) {
	if len(h.elements) == 0 {
		return 0, false
	}

	maxVal := h.elements[0]

	lastIndex := len(h.elements) - 1
	h.elements[0] = h.elements[lastIndex]

	h.elements = h.elements[:lastIndex]

	if len(h.elements) > 0 {
		h.BubbleDown(0)
	}

	return maxVal, true
}

func (h *MaxHeap) BubbleUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2

		if h.elements[parentIndex] >= h.elements[index] {
			break
		}

		h.elements[parentIndex], h.elements[index] = h.elements[index], h.elements[parentIndex]
	}
}

func (h *MaxHeap) BubbleDown(index int) {
	for {
		largest := index
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2

		if leftChildIndex < len(h.elements) && h.elements[leftChildIndex] > h.elements[largest] {
			largest = leftChildIndex
		}

		if rightChildIndex < len(h.elements) && h.elements[rightChildIndex] > h.elements[largest] {
			largest = rightChildIndex
		}

		if largest == index {
			break
		}

		h.elements[index], h.elements[largest] = h.elements[largest], h.elements[index]
		index = largest
	}
}
