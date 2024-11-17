package main

import "fmt"

// Heapify a subtree rooted at index i, assuming left and right subtrees are already heaps.
func maxHeapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// If left child is larger than root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// If right child is larger than root
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// If largest is not root, swap and heapify the affected subtree
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, n, largest)
	}
}

// BuildMaxHeap builds a max-heap from the given array
func buildMaxHeap(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify(arr, n, i)
	}
}

// HeapSortMax performs heap sort using Max-Heap
func heapSortMax(arr []int) {
	n := len(arr)

	// Build the max-heap
	buildMaxHeap(arr)

	// One by one extract elements from the heap
	for i := n - 1; i > 0; i-- {
		// Swap the root (max element) with the last element
		arr[i], arr[0] = arr[0], arr[i]

		// Heapify the root element
		maxHeapify(arr, i, 0)
	}
}

// Heapify a subtree rooted at index i, assuming left and right subtrees are already heaps.
func minHeapify(arr []int, n, i int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	// If left child is smaller than root
	if left < n && arr[left] < arr[smallest] {
		smallest = left
	}

	// If right child is smaller than root
	if right < n && arr[right] < arr[smallest] {
		smallest = right
	}

	// If smallest is not root, swap and heapify the affected subtree
	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		minHeapify(arr, n, smallest)
	}
}

// BuildMinHeap builds a min-heap from the given array
func buildMinHeap(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		minHeapify(arr, n, i)
	}
}

// HeapSortMin performs heap sort using Min-Heap
func heapSortMin(arr []int) {
	n := len(arr)

	// Build the min-heap
	buildMinHeap(arr)

	// One by one extract elements from the heap
	for i := n - 1; i > 0; i-- {
		// Swap the root (min element) with the last element
		arr[i], arr[0] = arr[0], arr[i]

		// Heapify the root element
		minHeapify(arr, i, 0)
	}
}

func main() {
	// Max-Heap for Descending Order
	arrMax := []int{12, 11, 13, 5, 6, 7}
	heapSortMax(arrMax)
	fmt.Println("Sorted array (Max-Heap - Descending Order):", arrMax)

	// Min-Heap for Ascending Order
	arrMin := []int{12, 11, 13, 5, 6, 7}
	heapSortMin(arrMin)
	fmt.Println("Sorted array (Min-Heap - Ascending Order):", arrMin)
}
