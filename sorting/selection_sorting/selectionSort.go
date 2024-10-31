package selectionsorting

func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-2; i++ {
		minIndex := i
		for j := i; j < n-1; j++ {
			if arr[j] < arr[i] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
