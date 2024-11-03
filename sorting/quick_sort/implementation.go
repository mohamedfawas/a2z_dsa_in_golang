package quicksort

// quickSort function takes an array and sorts it in-place
// low is the starting index, high is the ending index
func QuickSort(arr []int, low, high int) {
	// Base case: if the array has one or zero elements, it is already sorted
	if low < high {
		// Partition the array and get the pivot index
		pi := Partition(arr, low, high)

		// Recursively sort elements before and after the partition
		QuickSort(arr, low, pi-1)  // Left side of pivot
		QuickSort(arr, pi+1, high) // Right side of pivot
	}
}

// partition function rearranges elements in the array such that
// all elements less than the pivot are on its left, and those greater are on its right
// It returns the index of the pivot after partitioning
func Partition(arr []int, low, high int) int {
	// Choose the last element as the pivot
	pivot := arr[high]

	// i will track the "border" for elements less than the pivot
	i := low - 1

	// Loop through the array to rearrange elements around the pivot
	for j := low; j < high; j++ {
		// If the current element is less than or equal to pivot,
		// we place it in the "less than" section by swapping it
		if arr[j] <= pivot {
			i++                             // Move the boundary for smaller elements
			arr[i], arr[j] = arr[j], arr[i] // Swap to place smaller element
		}
	}

	// Place the pivot in its correct position by swapping
	arr[i+1], arr[high] = arr[high], arr[i+1]

	// Return the index of the pivot after partitioning
	return i + 1
}
