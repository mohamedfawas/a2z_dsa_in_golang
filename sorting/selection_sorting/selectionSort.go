package selectionsorting

func SelectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		// Find min element in unsorted portion of the array
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// Swap the found minimum element with the first element
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
